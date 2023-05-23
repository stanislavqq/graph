package money

import (
	"errors"
	"fmt"
)

type MoneyCollection struct {
	Items []MoneyRow
}

func NewCollection() *MoneyCollection {
	return &MoneyCollection{Items: []MoneyRow{}}
}

func (c *MoneyCollection) Add(value []string) {
	c.Items = append(c.Items, NewFromRow(value))
}

func (c *MoneyCollection) Get(id int) (MoneyRow, error) {
	for _, item := range c.Items {
		if item.GetID() == id {
			return item, nil
		}
	}

	return MoneyRow{}, errors.New(fmt.Sprintf("Row not found with id(%d)", id))
}

func (c *MoneyCollection) GetLength() int {
	return len(c.Items)
}

func (c *MoneyCollection) Each(callable func(key int, row MoneyRow)) {
	for key, row := range c.Items {
		callable(key, row)
	}
}
