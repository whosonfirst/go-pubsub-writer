package pubsub

import (
	"github.com/go-redis/redis"
	"strings"
)

type Writer struct {
	Channel string
	Client  *redis.Client
	Pubsub  *redis.PubSub
}

func NewWriter(channel string) (*Writer, error) {
	return newwriter("localhost:6379", "", 0, channel)
}

/*
func NewWriter(addr string, pswd string, db int, channel string) (*Writer, error) {
     return newwriter(addr, pswd, db, channel)
}
*/

func newwriter(addr string, pswd string, db int64, channel string) (*Writer, error) {

	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pswd,
		DB:       db,
	})

	_, err := client.Ping().Result()

	if err != nil {
		return nil, err
	}

	pubsub, err := client.Subscribe(channel)

	if err != nil {
		return nil, err
	}

	defer pubsub.Close()

	w := Writer{
		Channel: channel,
		Client:  client,
		Pubsub:  pubsub,
	}

	return &w, nil
}

func (w Writer) WriteString(s string) (n int64, err error) {
	r := strings.NewReader(s)
	return r.WriteTo(w)
}

func (w Writer) Write(p []byte) (n int, err error) {

	var msg string
	msg = string(p[:])

	err = w.Client.Publish(w.Channel, msg).Err()

	if err != nil {
		return 0, err
	}

	count := len(msg)
	return count, nil
}

func (w Writer) Close() error {

	w.Pubsub.Close()
	return nil
}
