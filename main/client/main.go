package main

import (
	"flag"
	"fmt"

	"github.com/joaosoft/socket"
)

func main() {
	var user string
	flag.StringVar(&user, "listen", "", "The user identifier")
	flag.Parse()

	if user == "" {
		panic("invalid user identifier")
	}

	client, err := socket.NewClient()
	if err != nil {
		panic(err)
	}

	if err := client.Start(); err != nil {
		panic(err)
	}

	if err := client.Subscribe(user, "messenger::in-box"); err != nil {
		panic(err)
	}

	client.Listen(user, "messenger-in-box", func(message []byte) error {
		fmt.Printf("\nreceived on listener the message %s", string(message))
		return nil
	})

	// Unsubscribe
	//if err := client.Unsubscribe("topic_1", "channel_1"); err != nil {
	//	panic(err)
	//}

	client.Wait()
}
