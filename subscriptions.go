package main

import "time"

func (client *MyClient) Subscribe(eventName string, query Query, varsParams ...map[string]interface{}) error {
	go func() {
		for {
			client.RunQuery(query, varsParams...)
			client.Emit(eventName, query)

			time.Sleep(5 * time.Second)
		}
	}()

	return nil
}
