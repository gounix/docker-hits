package data

import (
	"fmt"
	"sync"
	"time"
)

var (
  mu sync.Mutex
  last []byte
  timestamp time.Time
)

func Put(value []byte) {
    mu.Lock()
    defer mu.Unlock()

    fmt.Println("Put ", string(value))
    last = value
    timestamp = time.Now()
}

func Get() []byte {
    mu.Lock()
    defer mu.Unlock()

    fmt.Println("Get ", string(last))
    return last
}

func Alive(interval int) bool {
	now := time.Now()
	diff := now.Sub(timestamp)
	fmt.Printf("Alive: data is %.1f seconds old OK=%t\n", diff.Seconds(), diff.Seconds() < float64(2 * interval))
	// consider the producer dead after it missed 2 intervals
	return diff.Seconds() < float64(2 * interval)
}

