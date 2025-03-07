package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
	"time"
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

func (e *Estoque) AddItem(item models.Item, user string) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("quantity must be greater than zero! [Item ID: %d]", item.ID)
	}

	existingItem, exists := e.items[strconv.Itoa(item.ID)]
	if exists {
		item.Quantity += existingItem.Quantity
	}

	e.items[strconv.Itoa(item.ID)] = item

	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Entrada de estoque",
		User:      user,
		ItemID:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Adicionando novos item no estoque",
	})

	return nil
}

func (e *Estoque) RemoveItem(itemID int, quantity int, user string) error {
	existingItem, exits := e.items[strconv.Itoa(itemID)]
	if !exits {
		return fmt.Errorf("erro ao remover item: [ID: %d] não existe no estoque", itemID)
	}

	if quantity <= 0 {
		return fmt.Errorf(
			"erro ao remover item: quantidade inválida (zero ou negativa) para [ID: %d]",
			itemID,
		)
	}

	if existingItem.Quantity < quantity {
		return fmt.Errorf(
			"erro ao remover item: estoque insuficiente para [ID:%d]. Disponível: %d, Solicitado: %d",
			itemID, existingItem.Quantity, quantity,
		)
	}

	existingItem.Quantity -= quantity
	if existingItem.Quantity == 0 {
		delete(e.items, strconv.Itoa(itemID))
	} else {
		e.items[strconv.Itoa(itemID)] = existingItem
	}

	e.logs = append(e.logs, models.Log{
		Timestamp: time.Now(),
		Action:    "Saída de estoque",
		User:      user,
		ItemID:    itemID,
		Quantity:  quantity,
		Reason:    "Removendo itens do estoque",
	})

	return nil
}

func (e *Estoque) ListItems() []models.Item {
	var itemList []models.Item
	for _, item := range e.items {
		itemList = append(itemList, item)
	}

	return itemList
}

func (e *Estoque) ViewAuditLog() []models.Log {
	return e.logs
}

func (e *Estoque) CalculateTotalCost() float64 {
	var totalCost float64

	for _, item := range e.items {
		totalCost += float64(item.Quantity) * item.Price
	}

	return totalCost
}

// Generics
func FindBy[T any](data []T, comparator func(T) bool) ([]T, error) {
	var result []T
	for _, v := range data {
		if comparator(v) {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("nenhum item foi encontrado")
	}

	return result, nil
}
