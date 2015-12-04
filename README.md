# go-pubsub-writer

A Go package for sending Redis PubSub messages using a standard io.Writer interface

## Usage

```
package main

import (
       "flag"
       "github.com/whosonfirst/go-pubsub-writer"
       "strings"
       )

func main() {

     flag.Parse()
     args := flag.Args()

     var channel = flag.String("channel", "", "The PubSub channel to connect to")

     w, err := pubsub.NewWriter(*channel)

     if err != nil {
     	panic(err)
     }

     msg := strings.Join(args, " ")

     b := []byte("Write " + msg)

     _, err = w.Write(b)

     if err != nil {
   	panic(err)
     }
	
     // WriteString is provided as a convenience if you don't
     // feel like []byte("-ing") all the things per the default
     // io.Writer interface spec

     _, err = w.WriteString("WriteString " + msg)

     if err != nil {
     	panic(err)
     }     
}
```

## Caveats

Currently this only connects to a Redis server on localhost:6379
