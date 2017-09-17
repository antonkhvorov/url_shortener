package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/cli"
	"fmt"
	"golang.org/x/net/context"
	proto "gomicroexample/proto"
	"errors"
	"os"
)

type UrlShortener struct {}

// Генерация короткого URL на основе полного
func (g *UrlShortener) GenerateShortUrl(ctx context.Context, url *proto.Url, rsp *proto.Response) error {
	//TODO: Хэш
	//TODO: Запись в БД
	if url.Url != "pop" {
		rsp.OperationResponse = "Done"
	} else {
		rsp.OperationResponse = "Already existing"
	}
	fmt.Println(rsp.OperationResponse)
	return nil
}

// Получение короткого URL по полному
func (g *UrlShortener) AcquireShortUrl(ctx context.Context, url *proto.Url, shortUrl *proto.ShortUrl) error {
	//TODO: Добавить условие что не существует в БД этот URL
	if url.Url != "" {
		//TODO: Взять из БД сокращеный URL
		shortUrl.ShortUrl = "Abc123"
	} else {
		return errors.New("URL was not found")
	}
	fmt.Println(shortUrl.ShortUrl)
	return nil
}

// Замена всех URL в произвольном тексте на сокращенные URL
func (g *UrlShortener) ReplaceAllUrlsByShortUrl(ctx context.Context, text *proto.TextWithUrls, textWithShort *proto.TextWithShortUrls) error {
	//TODO: Взять из БД все сокращенные URL
	//TODO: Пройтись по ним, если находится совпадение в тексте — заменить
	if text.Text != "" {

	} else {
		return errors.New("empty text was passed")
	}

	return nil
}

func runClient(service micro.Service) {
	// Create new greeter client
	urlShortener := proto.NewUrlShortenerClient("url_shortener", service.Client())

	rsp, err := urlShortener.AcquireShortUrl(context.TODO(), &proto.Url{"testUrl"})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print response
	fmt.Println(rsp.ShortUrl)
}

func main() {
	service := micro.NewService(
		micro.Name("url_shortener"),
		micro.Version("alpha"),
		micro.Flags(
			cli.StringFlag{
				Name: "environment",
				Usage: "Just for test",
			},
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

	proto.RegisterUrlShortenerHandler(service.Server(), new(UrlShortener))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}