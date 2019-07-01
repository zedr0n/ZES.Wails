package main

import (
	"github.com/leaanthony/mewn"
	"github.com/shurcooL/graphql"
	"github.com/wailsapp/wails"
)

func (client *MyClient) NumberOfRoots() int {

	sQuery := &StatsQuery{}
	client.RunQuery(sQuery)

	return sQuery.Stats.NumberOfRoots
}

func (client *MyClient) LastLog() string {
	aQuery := &LogQuery{}
	client.RunQuery(aQuery)

	return aQuery.Log
}

func (client *MyClient) SubscribeLogs(eventName string) {
	aQuery := &LogQuery{}
	client.Subscribe(eventName, aQuery)
}

func (client *MyClient) WailsInit(runtime *wails.Runtime) error {
	client.Client = graphql.NewClient("http://localhost:5000", nil)
	client.runtime = runtime

	return nil
}

func main() {

	client := &MyClient{}

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:     1024,
		Height:    768,
		Title:     "Wails test",
		JS:        js,
		CSS:       css,
		Colour:    "#131313",
		Resizable: true,
	})
	app.Bind(client)
	app.Run()
}
