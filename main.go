package main

import (
	"github.com/brandonhilkert/pipeline-activity/pubsub"
	"github.com/brandonhilkert/pipeline-activity/socket"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	redisUrl := os.Getenv("REDIS_URL")

	l := pubsub.NewListener(redisUrl)
	go l.Listen()

	s := socket.NewServer("/socket")
	go s.Listen()

	go func() {
		for {
			m := <-l.MsgCh
			sm := &socket.Message{m.Model, m.Action}
			s.SendAll(sm)
		}
	}()

	http.Handle("/", http.FileServer(http.Dir(".")))

	port := "9999"
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
