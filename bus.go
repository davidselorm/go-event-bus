package bus

import (
	"sync"
)

type EventBus struct {
	mu sync.RWMutex
	subscribers map[string][]chan interface{}
}

func New() *EventBus {
	return &EventBus{subscribers: make(map[string][]chan interface{})}
}
