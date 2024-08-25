package tests

import (
	"go-mediator/domain/orders"
	"go-mediator/mediator"
	"sync"
	"testing"
)

func TestMediatorPublishSync(t *testing.T) {
	m := mediator.NewMediator()

	// Create event handlers
	emailHandler := &orders.SendEmailHandler{}
	logHandler := &orders.LogOrderHandler{}

	// Subscribe handlers
	m.Subscribe("OrderPlacedEvent", emailHandler)
	m.Subscribe("OrderPlacedEvent", logHandler)

	// Publish an event synchronously
	event := orders.OrderPlacedEvent{OrderID: 123}
	m.Publish(event, mediator.Sync)
}

func TestMediatorPublishAsync(t *testing.T) {
	m := mediator.NewMediator()

	var wg sync.WaitGroup
	wg.Add(2)

	// Create event handlers
	emailHandler := &orders.SendEmailHandler{}
	logHandler := &orders.LogOrderHandler{}

	// Subscribe handlers to the event
	m.Subscribe("OrderPlacedEvent", emailHandler)
	m.Subscribe("OrderPlacedEvent", logHandler)

	// Publish an event asynchronously
	event := orders.OrderPlacedEvent{OrderID: 456}
	m.Publish(event, mediator.Async)

	// Wait for all handlers to finish processing
	wg.Wait()

}

func TestMediatorNoSubscribers(t *testing.T) {
	m := mediator.NewMediator()

	// Publish an event with no subscribers
	event := orders.OrderPlacedEvent{OrderID: 789}
	m.Publish(event, mediator.Sync)

	// Since there are no subscribers, nothing should happen
	// We can consider this test successful if no errors or panics occur
}
