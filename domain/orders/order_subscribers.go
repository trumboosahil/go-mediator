package orders

import "fmt"

type SendEmailHandler struct{}

func (h *SendEmailHandler) Handle(event interface{}) {
    if e, ok := event.(OrderPlacedEvent); ok {
        fmt.Printf("Sending email for order: %d\n", e.OrderID)
    }
}

type LogOrderHandler struct{}

func (h *LogOrderHandler) Handle(event interface{}) {
    if e, ok := event.(OrderPlacedEvent); ok {
        fmt.Printf("Logging order: %d\n", e.OrderID)
    }
}
