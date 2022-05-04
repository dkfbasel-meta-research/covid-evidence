package stores

import "fmt"

type Record struct {
	BasicID       int32                             `json:"-"`
	ID            int                               `json:"id,omitempty"`
	Sequence      int                               `json:"sequence,omitempty"`
	CreatedAt     string                            `json:"createdAt,omitempty"`
	CreatedBy     string                            `json:"createdBy,omitempty"`
	ModifiedAt    string                            `json:"modifiedAt,omitempty"`
	ModifiedBy    string                            `json:"modifiedBy,omitempty"`
	Fields        map[string]interface{}            `json:"fields"`
	IsUpdated     bool                              `json:"-"`
}

func (r *Record) Field(name string) string {
	value, ok := r.Fields[name]
	if !ok {
		return ""
	}
	return asString(value)
}

func (r *Record) Key() string {
	source := r.Field("source")
	sourceID := r.Field("source_id")
	if source == "" && sourceID == "" {
		return ""
	}
	if sourceID == "" && source != "" {
		return source
	}
	if source == "" && sourceID != "" {
		return sourceID
	}

	key := fmt.Sprintf("%s::%s", source, sourceID)
	return key
}
