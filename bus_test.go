package eventbus

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestEventBusWildcardMatching(t *testing.T) {
	bus := NewEventBus()
	var received atomic.Int32

	bus.Subscribe("telemetry.*", func(topic string, payload any) {
		received.Add(1)
	})

	bus.Publish("telemetry.cpu", 85)
	bus.Publish("telemetry.memory", 72)
	bus.Publish("orders.created", 101)

	time.Sleep(50 * time.Millisecond)
	if received.Load() != 2 {
		t.Fatalf("Expected 2 matching telemetry events, got %d", received.Load())
	}
}
