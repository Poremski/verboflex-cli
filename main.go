package main

import (
	"github.com/Poremski/verboflex-cli/verboflex"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

func main() {

	js := mewn.String("./frontend/build/static/js/main.js")
	css := mewn.String("./frontend/build/static/css/main.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "Verboflex",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(verboflex.NewController())
	if err := app.Run(); err != nil {
		panic(err)
	}
}
