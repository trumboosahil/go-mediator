package payments

import "go-mediator/mediator"

type PaymentAggregate struct {
	mediator *mediator.Mediator
}

func NewPaymentAggregate(m *mediator.Mediator) *PaymentAggregate {
	aggregate := &PaymentAggregate{mediator: m}
	aggregate.configureSubscriptions()
	return aggregate
}

func (a *PaymentAggregate) configureSubscriptions() {
	a.mediator.Subscribe("PaymentProcessedEvent", &ProcessPaymentHandler{})
	a.mediator.Subscribe("OrderShippedEvent", &ProcessPaymentHandler{})
}

func (a *PaymentAggregate) ProcessPayment(paymentID int) {
	event := PaymentProcessedEvent{PaymentID: paymentID}
	a.mediator.Publish(event, mediator.Async)
}
