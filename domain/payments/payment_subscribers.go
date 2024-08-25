package payments

import (
	"fmt"
	"go-mediator/domain/orders"
)

type ProcessPaymentHandler struct{}

func (h *ProcessPaymentHandler) Handle(event interface{}) {
	switch e := event.(type) {
	case orders.OrderShippedEvent:
		fmt.Printf("Order shipped: %d. Starting payment processing...\n", e.OrderID)
		// Here you would initiate payment processing logic
	case PaymentProcessedEvent:
		fmt.Printf("Processing payment: %d\n", e.PaymentID)
	}
}
