package stores

// IStore ...
type IStore interface {
	Setup(string, bool) // setup function is called when the application starts
	Close()             // close the connection to the store
	Fetch(string) ([]Record, error)
	Update(string, []*Record) error
	AddBasic([]*Record) error
	UpdateBasic([]*Record) error
	FetchBasic(sources ...string) ([]Record, Index, error)
	FetchScreening(string, []string, string) ([]Record, error)
	SetScreening(string) error
	SaveIntervention(int32, []map[string]string) error
	FindExistingInterventions(int32, string) error
	CheckForResults(string) (bool, error)
	SetTopics([]Record) error
	CleanTrials([]string) ([]string, error)
	AddTrial(string, string, []string) error
}
