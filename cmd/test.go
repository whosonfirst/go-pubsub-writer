package main

import (
	"flag"
	"github.com/whosonfirst/go-pubsub-writer"
	"strings"
)

func main() {

	var channel = flag.String("channel", "", "The PubSub channel to connect to")

	flag.Parse()
	args := flag.Args()

	w, err := pubsub.NewWriter(*channel)

	if err != nil {
		panic(err)
	}

	msg := strings.Join(args, " ")

	_, err = w.WriteString(msg)

	if err != nil {
		panic(err)
	}
}
