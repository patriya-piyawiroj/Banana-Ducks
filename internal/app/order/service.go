package order

type Shop interface {
	GetAllOrders() (Order, error)
	CreateOrder() error
}

type ShopSvc struct {
}
