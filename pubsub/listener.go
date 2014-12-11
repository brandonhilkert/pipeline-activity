package pubsub

import (
	"github.com/garyburd/redigo/redis"
	"log"
)

type Listener struct {
	RedisUrl   string
	PubSubConn redis.PubSubConn
	MsgCh      chan *Message
}

func NewListener(redisUrl string) *Listener {
	msgCh := make(chan *Message)

	return &Listener{
		RedisUrl: redisUrl,
		MsgCh:    msgCh,
	}
}

func (l *Listener) Connect() {
	c, err := redis.Dial("tcp", l.RedisUrl+":6379")

	if err != nil {
		log.Fatal("Couldn't connect to Redis")
	}

	l.PubSubConn = redis.PubSubConn{Conn: c}
	err = l.PubSubConn.PSubscribe("*")

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Listening to PubSub...")
}

func (l *Listener) Process() {
	for {
		switch v := l.PubSubConn.Receive().(type) {
		case redis.PMessage:
			m := NewMessage(v.Channel)
			log.Printf("Message found: %v\n", m)
			l.MsgCh <- m
		case error:
			log.Printf("error: %v\n", v)
		}
	}
}

func (l *Listener) Listen() {
	l.Connect()
	l.Process()
	l.PubSubConn.Conn.Close()
}
