package cache

import (
	"sync"
	"time"
)

type cache struct {
	key        string
	v          interface{}
	expiration time.Duration
	Interval   time.Duration
	mu         *sync.Mutex
}

type Cache struct {
	*cache
}

const defaultExpiretionTime = 15 * time.Minute
const defaultIntervalTime = 1 * time.Minute

func New() *cache {
	return &cache{
		expiration: defaultExpiretionTime,
		Interval:   defaultIntervalTime,
	}
}

func (c *cache) set() {

}

func (c *cache) get() {

}

func (c *cache) del() {

}

func isExpire() bool {
	return true
}

func (c *cache) run() {
	timer := time.NewTicker(c.Interval)
	for {
		select {
		case <-timer.C:

		}
	}
}
