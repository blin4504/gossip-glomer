package main

import (
	"encoding/json"
	"fmt"
	"log"

	maelstorm "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstorm.NewNode()

	n.Handle("echo", func(msg maelstorm.Message) error {
		fmt.Println("Pong")
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "echo_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
