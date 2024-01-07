package usecase

import (
	"github.com/gustavohenriquess/go-intensive23/internal/entity"
)

type OrderInput struct {
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutput struct {
	ID         int64
	Price      float64
	Tax        float64
	FinalPrice float64
}

// SOLID - "D" - Dependency Inversion Principle
type CalculateFinalPrice struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPrice(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPrice {
	return &CalculateFinalPrice{
		OrderRepository: orderRepository,
	}
}

func (c *CalculateFinalPrice) Execute(input OrderInput) (*OrderOutput, error) {
	order, err := entity.NewOrder(input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	id, err := c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutput{
		ID:         id,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
