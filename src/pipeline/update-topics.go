package pipeline

import (
	"dkfbasel.ch/covid-evidence/logger"
)

// UpdateTopics set topics
func (p *Pipeline) UpdateTopics() error {

	// fetch topics from the database
	topics, err := p.Store.Fetch(`
		SELECT id, name, keyword_search_columns, keywords
		FROM topic
		WHERE (
			keyword_search_columns IS NOT NULL AND
			keyword_search_columns <> '' AND
			keywords IS NOT NULL AND
			keywords <> ''
		);
	`)
	if err != nil {
		return logger.NewError("could not fetch topics", nil)
	}

	return p.Store.SetTopics(topics)
}
