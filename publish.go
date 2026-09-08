package eventbus

import "sync"

type SyncEventBus struct {
	bus *EventBus
}

func NewSyncEventBus() *SyncEventBus {
	return &SyncEventBus{
		bus: NewEventBus(),
	}
}

func (s *SyncEventBus) PublishSync(topic string, payload any) {
	s.bus.mu.RLock()
	defer s.bus.mu.RUnlock()

	var wg sync.WaitGroup
	for subTopic, subs := range s.bus.subscribers {
		if matchTopic(subTopic, topic) {
			for _, sub := range subs {
				wg.Add(1)
				go func(sb *Subscription) {
					defer wg.Done()
					sb.handler(topic, payload)
				}(sub)
			}
		}
	}
	wg.Wait()
}
