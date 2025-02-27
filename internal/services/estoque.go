package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
)

type Estoque struct {
	items map[string]models.Item
	logs  []models.Log
}

func NewEstoque() *Estoque {
	return &Estoque{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

func (e *Estoque) AddItem(item models.Item) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero! [Item ID: %d]", item.ID)
	}

	existingItem, exists := e.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}

	e.items[strconv.Itoa(item.ID)] = item
	return nil
}

func (e *Estoque) ListItems() []models.Item {
	var itemList []models.Item
	for _, item := range e.items {
		itemList = append(itemList, item)
	}

	return itemList
}
