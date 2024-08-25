package orders

type OrderPlacedEvent struct {
    OrderID int
}

type OrderShippedEvent struct {
    OrderID int
}
