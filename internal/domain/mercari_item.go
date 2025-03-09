package domain

import "fmt"

// under stand why this domain imoport is needed

// MercariItem is a struct that represents an item for sale on Mercari
type MercariItem struct {
	Name        string
	PriceInYen  int
	Condition   string
	Description string
	ItemID      string
}

func (item MercariItem) Summarize() string {
	return fmt.Sprintf("%s - Condition: %s, Price: ¥%d", item.Name, item.Condition, item.PriceInYen)
}