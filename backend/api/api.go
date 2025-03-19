package api

import (
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func Run() {
	r := router.New()
	r.GET("/api", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Content-Type", "application/json") 
		ctx.SetStatusCode(fasthttp.StatusOK)                   
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*") // CORS
		ctx.SetBodyString(`From Golang`) // Message
	})

	if err := fasthttp.ListenAndServe(":1234", r.Handler); err != nil {
        panic(err) // Обработка ошибки при запуске сервера
    }
}