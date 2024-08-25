package main

import (
	"go-mediator/domain/orders"
	"go-mediator/domain/payments"
	"go-mediator/mediator"
)

func main() {
	// Initialize mediator
	m := mediator.NewMediator()

	// Initialize aggregates (each manages its own event subscriptions)
	orderAggregate := orders.NewOrderAggregate(m)
	paymentAggregate := payments.NewPaymentAggregate(m)

	// Execute domain logic
	orderAggregate.PlaceOrder(123)
	orderAggregate.ShipOrder(123)
	paymentAggregate.ProcessPayment(123)
}
