package api

import "mercari-shopper-go/internal/domain"

// MercariAPIClient interface defines the contract for interacting with the Mercari API.
type MercariAPIAdaptor interface {
	// this is the method that will be used to search for items on Mercari inthe usecase
	SearchItems(query string) ([]domain.MercariItem, error)
}

// --- Dummy implementation of MercariAPIClient ---
type DummyMercariAPIClient struct{}

func (client DummyMercariAPIClient) SearchItems(query string) ([]domain.MercariItem, error) {
	// Simulate API call and return dummy Mercari items
	items := []domain.MercariItem{
		{Name: "API Dummy Item 1", PriceInYen: 1500, Condition: "Like New", Description: "Dummy item from API adapter.", ItemID: "api-dummy1"},
		{Name: "API Dummy Item 2", PriceInYen: 3000, Condition: "Good", Description: "Another API dummy item.", ItemID: "api-dummy2"},
	}
	return items, nil
}

// --- End of Dummy implementation of MercariAPIClient ---