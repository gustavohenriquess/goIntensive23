package entity

import "errors"

type Order struct {
	Price      float64
	Tax        float64
	FinalPrice float64
}

// Method to calculate the final price of an order
func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.Validade()

	if err != nil {
		return err
	}

	return nil
}

func (o *Order) Validade() error {
	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}
	if o.Tax <= 0 {
		return errors.New("invalid tax")
	}
	return nil
}

func NewOrder(price float64, tax float64) (*Order, error) {
	order := &Order{
		Price: price,
		Tax:   tax,
	}

	err := order.Validade()

	if err != nil {
		return nil, err
	}

	return order, nil
}
