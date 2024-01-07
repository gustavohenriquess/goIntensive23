package entity

type OrderRepositoryInterface interface {
	Save(order *Order) (int64, error)
	GetTotalTransactions() (int, error)
}
