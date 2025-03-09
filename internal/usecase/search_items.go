// responsible for searching items usecase

package usecase

import (
	"errors"
	"fmt"
	domain "mercari-shopper-go/internal/domain"
	api "mercari-shopper-go/internal/interface-adaptors/api"
)

// serachitemusecase as an interface
type SearchItemUsecase interface {
	SearchItems(query string) (domain.MercariItem, error)
}

// --- Real implementation of searchitemusecase ---
type RealSearchItemsUseCase struct {
	// the api client that will be used to search for items
	// for now we are using the dummy api client for this real use case
	apiClient api.DummyMercariAPIClient
}

// NewRealSearchItemsUseCase creates a new RealSearchItemsUseCase with the given API client.
func NewRealSearchItemsUseCase(client api.DummyMercariAPIClient) RealSearchItemsUseCase {
	return RealSearchItemsUseCase{apiClient: client}
}

func (uc RealSearchItemsUseCase) SearchItems(query string) ([]domain.MercariItem, error) {
	// search for items using the api client
	if query == "" {
		return nil, errors.New("query cannot be empty")
	}

	items, err := uc.apiClient.SearchItems(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching items from Mercari API: %w", err) // Wrap API errors
	}
	return items, nil
}


// --- Dummy implementation of searchitemusecase ---
type DummySearchItemUsecase struct {}

func (uc DummySearchItemUsecase) SearchItems(query string) ([]domain.MercariItem, error) {
	// create some dummy data
	items := []domain.MercariItem{
		{
			Name:        "dummy hardcoded item 1",
			PriceInYen:  12000,
			Condition:   "Used - Very Good",
			Description: "Classic dummy item 1.",
			ItemID:      "dummy-item1",
		},
		{
			Name:        "dummy hardcoded item 2",
			PriceInYen:  8000,
			Condition:   "Used - Good",
			Description: "Classic dummy item 2.",
			ItemID:      "dummy-item2",
		},
	}
	return items, nil
}

// --- End of Dummy implementation of searchitemusecase ---