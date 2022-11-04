package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/ably/ably-go/ably"
)

func main() {
	fmt.Println("Type your ClientID")
	reader := bufio.NewReader(os.Stdin)
	clientID, _ := reader.ReadString('\n')
	clientID = strings.Replace(clientID, "\n", "", -1)

	// Connect to Ably using the API key and ClientID specified above
	client, err := ably.NewRealtime(ably.WithClientID(clientID))

	if err != nil {
		panic(err)
	}
}

func publishing(channel *ably.RealtimeChannel) {
	reader := bufio.NewREader(os.Stdin)

	for {
		text, _ := reader.ReadString('\n')
		text = strings.ReplaceAll(text, "\n", "")

		// Publish the message typed int he Ably Channel
		err := channel.Publish(context.Background(), "message", text)

		// await confirmation that message was received by Ably
		if err != nil {
			err := fmt.Errorf("publishing to channel: %w", err)
			fmt.Println(err)
		}
	}
}

func subscribing(channel *ably.RealtimeChannel) {
	// Subscribe to messages sent on the channel

	_, err := channel.SubscribeAll(context.Background(), func(msg *ably.Message) {
		fmt.Printf("Message received: %s\n", string(msg.ClientID), msg.Data)
	})

	if err != nil {
		err := fmt.Errorf("subscribing to channel: %w", err)
		fmt.Println(err)
	}
}
