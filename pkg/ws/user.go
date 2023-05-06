package ws

import (
	"io"
	"sync"

	"github.com/eadydb/spider/pkg/gopool"
)

type User struct {
	io   sync.Mutex         // io mutex
	conn io.ReadWriteCloser // connection

	id   uint     // user id
	name string   // user name
	tags []string // user tags
	chat *Chat    // chat
}

type Chat struct {
	mu sync.RWMutex       // read-write mutex
	us []*User            // all users
	ts map[string][]*User // tags users
	ns map[string]*User   // signal user

	pool *gopool.Pool // goroutine pool
	out  chan []byte  // broadcast message
}
