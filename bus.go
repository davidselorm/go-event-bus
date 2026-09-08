package eventbus

import (
	"strings"
	"sync"
)

type Handler func(topic string, payload any)

type Subscription struct {
	topic   string
	handler Handler
	ch      chan any
}

type EventBus struct {
	mu          sync.RWMutex
	subscribers map[string][]*Subscription
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]*Subscription),
	}
}

func (b *EventBus) Subscribe(topic string, h Handler) *Subscription {
	b.mu.Lock()
	defer b.mu.Unlock()

	sub := &Subscription{
		topic:   topic,
		handler: h,
		ch:      make(chan any, 64),
	}
	b.subscribers[topic] = append(b.subscribers[topic], sub)
	return sub
}

func (b *EventBus) Publish(topic string, payload any) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	for subTopic, subs := range b.subscribers {
		if matchTopic(subTopic, topic) {
			for _, sub := range subs {
				go sub.handler(topic, payload)
			}
		}
	}
}

func matchTopic(pattern, topic string) bool {
	if pattern == "#" || pattern == topic {
		return true
	}
	pParts := strings.Split(pattern, ".")
	tParts := strings.Split(topic, ".")

	for i := 0; i < len(pParts); i++ {
		if pParts[i] == "#" {
			return true
		}
		if i >= len(tParts) {
			return false
		}
		if pParts[i] != "*" && pParts[i] != tParts[i] {
			return false
		}
	}
	return len(pParts) == len(tParts)
}
