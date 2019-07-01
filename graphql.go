package main

import (
	"github.com/shurcooL/graphql"
	"github.com/wailsapp/wails"
)

type MyClient struct {
	*graphql.Client
	runtime *wails.Runtime
}

// Create new connection to localhost graphql server
func (client *MyClient) Connect() {
	client.Client = graphql.NewClient("http://localhost:5000", nil)
}

func (client *MyClient) Emit(eventName string, optionalData ...interface{}) {
	client.runtime.Events.Emit(eventName, optionalData)
}
