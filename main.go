package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jaden-hoenes/go-server-template/constants/env"
	"github.com/jaden-hoenes/go-server-template/constants/file"
	"github.com/jaden-hoenes/go-server-template/constants/url"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/savsgio/atreugo/v11"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal().Msgf("Error loading .env file: %+v", err)
	}

	config := atreugo.Config{
		Addr: "0.0.0.0:" + os.Getenv(env.Port),
		PanicView: func(rc *atreugo.RequestCtx, i interface{}) {
			rc.TextResponse(fmt.Sprintf("%+v", i), http.StatusInternalServerError)
		},
	}
	server := atreugo.New(config)

	server.UseAfter(func(rc *atreugo.RequestCtx) error {
		log.Trace().Msgf("Endpoint Call: %s %s, Status Code: %d", rc.Method(), rc.Path(), rc.Response.Header.StatusCode())
		return nil
	})

	server.GET(url.Index, func(ctx *atreugo.RequestCtx) error {
		return ctx.HTTPResponse("<html><head><link rel=\"icon\" type=\"image/x-icon\" href=\"/assets/images/favicon.ico\"/></head><body>Hello World</body></html>", http.StatusOK)
	})

	server.Static(url.Favicon, file.Favicon)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
