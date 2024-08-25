package orders

import "go-mediator/mediator"

type OrderAggregate struct {
    mediator *mediator.Mediator
}

func NewOrderAggregate(m *mediator.Mediator) *OrderAggregate {
    aggregate := &OrderAggregate{mediator: m}
    aggregate.configureSubscriptions()
    return aggregate
}

func (a *OrderAggregate) configureSubscriptions() {
    a.mediator.Subscribe("OrderPlacedEvent", &SendEmailHandler{})
    a.mediator.Subscribe("OrderPlacedEvent", &LogOrderHandler{})
}

func (a *OrderAggregate) PlaceOrder(orderID int) {
    event := OrderPlacedEvent{OrderID: orderID}
    a.mediator.Publish(event, mediator.Async)
}

func (a *OrderAggregate) ShipOrder(orderID int) {
    event := OrderShippedEvent{OrderID: orderID}
    a.mediator.Publish(event, mediator.Sync)
}
