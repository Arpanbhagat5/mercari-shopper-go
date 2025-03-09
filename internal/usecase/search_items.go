// responsible for searching items usecase

package usecase

import (
	domain "mercari-shopper-go/internal/domain"
)

// serachitemusecase as an interface
type SearchItemUsecase interface {
	SearchItems(query string) (domain.MercariItem, error)
}

// dummy implementation of searchitemusecase
type DummySearchItemUsecase struct {}

// SearchItems is a dummy implementation of the SearchItems method
func (uc DummySearchItemUsecase) SearchItems(query string) ([]domain.MercariItem, error) {
	// create some dummy data
	items := []domain.MercariItem{
		{
			Name:        "dummy object 1",
			PriceInYen:  12000,
			Condition:   "Used - Very Good",
			Description: "Classic dummy object 1.",
			ItemID:      "m99887766554",
		},
		{
			Name:        "dummy object 2",
			PriceInYen:  8000,
			Condition:   "Used - Good",
			Description: "Classic dummy object 2.",
			ItemID:      "m99887766555",
		},
	}
	return items, nil
}
