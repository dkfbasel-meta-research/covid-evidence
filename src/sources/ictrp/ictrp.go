package ictrp

import "dkfbasel.ch/covid-evidence/stores"

// ICTRP ...
type ICTRP struct {
	ID    string
	Path  string
	Store stores.IStore
}

// NewSource initialize a new ictrp source
func NewSource(store stores.IStore, path string) *ICTRP {
	source := ICTRP{}
	source.Path = path
	source.ID = "ictrp"
	source.Store = store
	return &source
}

// GetID will return the store identifier
func (who *ICTRP) GetID() string {
	return who.ID
}
