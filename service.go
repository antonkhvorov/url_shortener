package main

import (
	"fmt"
	"golang.org/x/net/context"
	"os"
	fnv "hash/fnv"
	"strings"

	//project
	proto "./proto"

	//go-micro
	"github.com/micro/go-micro"
	"github.com/micro/cli"

	//MongoDB
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	//TODO: Удалить после альфа-версии
	"io/ioutil"
)

type Shortener struct{}

type Url struct {
	URL 			string 	`json:"url"`
	ShortURL		uint32	`json:"shortUrl"`
	TimesClicked	int		`json:"timesClicked"`
}

// Генерация короткого URL на основе полного
func (g *Shortener) AddShort(ctx context.Context, url *proto.UrlRequest, rsp *proto.UrlResponse) error {
	//TODO: Добавить проверок на валидность данных
	var u Url
	u.URL = url.Url
	u.ShortURL = hash(url.Url)
	u.TimesClicked = 0

	//TODO: Реюзать MongoDB код
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	c := session.DB("url_shortener").C("urls")
	err = c.Insert(u)

	rsp.OperationResponse = "Successfully added to db url: " + url.Url
	return nil
}

// Получение короткого URL по полному
func (g *Shortener) GetShort(ctx context.Context, url *proto.UrlRequest, shortUrl *proto.ShortUrlResponse) error {
	//TODO: Добавить проверок на валидность данных (в БД нет этого URL)

	//TODO: Реюзать MongoDB код
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	c := session.DB("url_shortener").C("urls")

	var u Url
	err = c.Find(bson.M{"url": url.Url}).One(&u)
	if err != nil {
		panic(err)
	}

	shortUrl.ShortUrl = fmt.Sprint(u.ShortURL)

	return nil
}

// Замена всех URL в произвольном тексте на сокращенные URL
func (g *Shortener) ReplaceAll(ctx context.Context, text *proto.TextRequest, textWithShort *proto.TextResponse) error {
	//TODO: Добавить проверок на валидность данных

	//TODO: Реюзать MongoDB код
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	c := session.DB("url_shortener").C("urls")

	var u []Url
	err = c.Find(bson.M{}).All(&u)

	fmt.Printf("%+v\n", u)

	textWithShort.TextWithShort = text.Text
	for _, sUrl := range u {
		textWithShort.TextWithShort = strings.Replace(textWithShort.TextWithShort, sUrl.URL, fmt.Sprint(sUrl.ShortURL), -1)
	}

	return nil
}

func runClient(service micro.Service) {
	// Имитация клиента :)
	Shortener := proto.NewShortenerClient("ushortenerclient", service.Client())

	// Добавим 3 URL
	rsp, err := Shortener.AddShort(context.TODO(), &proto.UrlRequest{"https://google.com"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.String())

	rsp, err = Shortener.AddShort(context.TODO(), &proto.UrlRequest{"https://yandex.ru"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.String())

	rsp, err = Shortener.AddShort(context.TODO(), &proto.UrlRequest{"https://vk.com"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp.String())

	// Запросим короткий URL у URL'a https://google.com
	rsp2, err2 := Shortener.GetShort(context.TODO(), &proto.UrlRequest{"https://google.com"})
	if err2 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp2.ShortUrl)


	// Заменим все URL из файла на короткие
	dat, err := ioutil.ReadFile("textWithURLs")
	fmt.Println(string(dat))
	rsp3, err3 := Shortener.ReplaceAll(context.TODO(), &proto.TextRequest{string(dat)})
	if err3 != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rsp3.TextWithShort)
}

func main() {
	// Настройка сервиса
	service := micro.NewService(
		micro.Name("go.micro.shortener"),
		micro.Version("alpha_v2"),
		micro.Flags(
			cli.BoolFlag{
				Name:  "run_client",
				Usage: "Launch the client",
			},
		),
	)

	service.Init(
		micro.Action(func(c *cli.Context) {
			if c.Bool("run_client") {
				runClient(service)
				os.Exit(0)
			}
		}),
	)

	proto.RegisterShortenerHandler(service.Server(), new(Shortener))

	// Запуск сервиса
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("url_shortener").C("urls")

	index := mgo.Index{
		Key:        []string{"url"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

// Функция хеша, результатом которой является последовательность из 9 цифр
func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}


//api.WithEndpoint(&api.Endpoint{
//// The RPC method
//Name: "UrlShortener.GenerateShortUrl",
//// The HTTP paths. This can be a POSIX regex
//Path: []string{"/generateshort"},
//// The HTTP Methods for this endpoint
//Method: []string{"POST"},
//// The API handler to use
//Handler: api.Rpc,
//
//}), api.WithEndpoint(&api.Endpoint{
//// The RPC method
//Name: "UrlShortener.AcquireShortUrl",
//// The HTTP paths. This can be a POSIX regex
//Path: []string{"/acquireshort"},
//// The HTTP Methods for this endpoint
//Method: []string{"GET"},
//// The API handler to use
//Handler: api.Rpc,
//
//}), api.WithEndpoint(&api.Endpoint{
//// The RPC method
//Name: "UrlShortener.ReplaceAllUrlsByShortUrl",
//// The HTTP paths. This can be a POSIX regex
//Path: []string{"/replaceallurls"},
//// The HTTP Methods for this endpoint
//Method: []string{"POST"},
//// The API handler to use
//Handler: api.Rpc,
//})