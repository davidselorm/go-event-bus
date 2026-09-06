package bus

func (b *EventBus) Publish(topic string, data interface{}) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, ch := range b.subscribers[topic] {
		select {
		case ch <- data:
		default:
		}
	}
}
