package sources

// Source defines the standard functions that a source must fulfil to work with
// the library
type Source interface {
	// Step 1
	GetID() string
	Fetch() (string, error)
	Parse(string) ([]map[string]string, error)
	Clean([]map[string]string) ([]map[string]string, int, int, error)

	// Step 2
	ImportScreening([]map[string]string) error

	// Step 3
	UpdateCoveBasic() error
	UpdateScreening() error
}
