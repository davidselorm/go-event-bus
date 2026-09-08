# go-event-bus

A high-performance in-memory pub/sub topic broker with wildcard topic matching in Go.

## Features
- **Wildcard Subscriptions**: Single-level (`*`) and multi-level (`#`) hierarchical topic dispatch.
- **Asynchronous & Synchronous Delivery**: Thread-safe concurrent fanout with goroutine worker pools.
