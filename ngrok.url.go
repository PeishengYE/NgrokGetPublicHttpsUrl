package main

import (
	"context"
	"fmt"
	"os"

	"github.com/ngrok/ngrok-api-go/v4"
	"github.com/ngrok/ngrok-api-go/v4/tunnels"
)

func main() {
	//fmt.Println(example(context.Background()))
	getHttpsPublicURL(context.Background())
}

func getHttpsPublicURL(ctx context.Context) error {
	// construct the api client
	clientConfig := ngrok.NewClientConfig(os.Getenv("NGROK_API_KEY"))

	// list all online tunnels
	tunnels := tunnels.NewClient(clientConfig)
	iter := tunnels.List(nil)
	for iter.Next(ctx) {
		//fmt.Println(iter.Item().ID)
		//fmt.Println(iter.Item().PublicURL)
		if(iter.Item().Proto == "https"){
		  fmt.Println("HTTPs", iter.Item().PublicURL)
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}
	return nil
}
