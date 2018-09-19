package main

import (
	"github.com/gramework/gramework"
)

func main() {
	app := gramework.New()
	app.Use(app.CORSMiddleware())
	app.GET("/", func(ctx *gramework.Context) error {
		ctx.HTML()
		_, err := ctx.WriteString(`<html><body style="background-color: green"></body></html>`)
		return err
	})
	app.ListenAndServe(":8080")
}
