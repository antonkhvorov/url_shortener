package main

import (
	"encoding/json"
	"log"
	"strings"

	shortener "../proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/errors"
	api "github.com/micro/micro/api/proto"

	"golang.org/x/net/context"
)

type Shortener struct {
	Client shortener.ShortenerClient
}

func (s *Shortener) AddShort(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Add Short API request")

	url, ok := req.Get["url"]
	if !ok || len(url.Values) == 0 {
		return errors.BadRequest("go.micro.shortener.api", "URL cannot be blank")
	}

	response, err := s.Client.AddShort(ctx, &shortener.UrlRequest{
		Url: strings.Join(url.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"response": response.OperationResponse,
	})
	rsp.Body = string(b)

	return nil
}

func (s *Shortener) GetShort(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Get Short API request")

	url, ok := req.Get["url"]
	if !ok || len(url.Values) == 0 {
		return errors.BadRequest("shortener.api", "URL cannot be blank")
	}

	response, err := s.Client.GetShort(ctx, &shortener.UrlRequest{
		Url: strings.Join(url.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"shortUrl": response.ShortUrl,
	})
	rsp.Body = string(b)

	return nil
}

func (s *Shortener) ReplaceAll(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Print("Received Replace All API request")

	url, ok := req.Get["text"]
	if !ok || len(url.Values) == 0 {
		return errors.BadRequest("shortener.api", "Text cannot be blank")
	}

	response, err := s.Client.ReplaceAll(ctx, &shortener.TextRequest{
		Text: strings.Join(url.Values, " "),
	})
	if err != nil {
		return err
	}

	rsp.StatusCode = 200
	b, _ := json.Marshal(map[string]string{
		"textWithShort": response.TextWithShort,
	})
	rsp.Body = string(b)

	return nil
}



func main() {
	service := micro.NewService(
		micro.Name("go.micro.api.shortener"),
	)

	// parse command line flags
	service.Init()

	service.Server().Handle(
		service.Server().NewHandler(
			&Shortener{Client: shortener.NewShortenerClient("go.micro.shortener", service.Client())},
		),
	)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
