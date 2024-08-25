package mediator

import (
    "sync"
    "reflect"
)

type EventHandler interface {
    Handle(event interface{})
}

type Mediator struct {
    subscribers map[string][]EventHandler
    mu          sync.RWMutex
}

func NewMediator() *Mediator {
    return &Mediator{
        subscribers: make(map[string][]EventHandler),
    }
}

func (m *Mediator) Subscribe(eventType string, handler EventHandler) {
    m.mu.Lock()
    defer m.mu.Unlock()

    m.subscribers[eventType] = append(m.subscribers[eventType], handler)
}

func (m *Mediator) Publish(event interface{}, mode ExecutionMode) {
    m.mu.RLock()
    defer m.mu.RUnlock()

    eventType := reflect.TypeOf(event).Name()
    if handlers, found := m.subscribers[eventType]; found {
        if mode == Async {
            var wg sync.WaitGroup
            wg.Add(len(handlers))

            for _, handler := range handlers {
                go func(h EventHandler) {
                    defer wg.Done()
                    h.Handle(event)
                }(handler)
            }

            wg.Wait()
        } else {
            for _, handler := range handlers {
                handler.Handle(event)
            }
        }
    }
}

type ExecutionMode int

const (
    Sync ExecutionMode = iota
    Async
)
