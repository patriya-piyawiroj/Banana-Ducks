package order

type OrderRepo interface {
	FindAll() (Order, error)
	Insert() error
}

type OrderMongoRepo struct {
}
