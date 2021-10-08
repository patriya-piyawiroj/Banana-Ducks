package order

const (
	Sell TxnType = iota
	Buy  
)

type (
	TxnType int

	Order {
		ID string
		TxnType TxnType
	}
)	