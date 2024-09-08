package main

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	maelstorm "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstorm.NewNode()

	n.Handle("generate", func(msg maelstorm.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "generate_ok"
		body["id"] = uuid.New().String()

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})

	if err := n.Run(); err != nil {
		log.Fatal(err)
	}

}
