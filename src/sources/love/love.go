package love

import "dkfbasel.ch/covid-evidence/stores"

// Love ...
type Love struct {
	ID    string
	Store stores.IStore
	Path  string
}

// NewSource initialize a new love source
func NewSource(store stores.IStore, path string) *Love {
	source := Love{}
	source.Path = path
	source.ID = "love"
	source.Store = store
	return &source
}

// GetID will return the store identifier
func (love *Love) GetID() string {
	return love.ID
}
