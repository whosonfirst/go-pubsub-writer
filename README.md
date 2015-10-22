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

     _, err = w.WriteString(msg)

     if err != nil {
     	panic(err)
     }     
}
```

## Caveats

Currently this only connects to a Redis server on localhost:6379
