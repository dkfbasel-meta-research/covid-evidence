package pipeline

import (
	"encoding/csv"
	"io"
	"os"

	"dkfbasel.ch/covid-evidence/logger"
	"dkfbasel.ch/covid-evidence/sources"
	"dkfbasel.ch/covid-evidence/stores"
	"dkfbasel.ch/covid-evidence/stores/mysql"
)

const exportPath = "./exports"

// Pipeline ...
type Pipeline struct {
	Sources []sources.Source
	Store   stores.IStore
	Tmp     string
}

// NewPipeline ...
func NewPipeline(db *mysql.Store, sources []sources.Source, tmpPath string) *Pipeline {
	p := Pipeline{}
	p.Store = db
	p.Sources = sources
	p.Tmp = tmpPath

	return &p
}

// Start ...
func (p *Pipeline) Start(screening, cove bool) {
	if screening {
		id := p.StartSourceStep()
		p.StartScreeningStep(id)
	}
	if cove {
		p.StartCoveBasicStep()
		p.UpdateTopics()
	}
	logger.Info("Pipeline done!")
}

func writeCSV(path string, data []map[string]string) error {

	// collect all available keys in data
	keys := make(map[string]bool)
	for i := range data {
		for k := range data[i] {
			keys[k] = true
		}
	}

	// add all keys ass title row
	csvData := make([][]string, len(data)+1)
	title := []string{}
	for k := range keys {
		title = append(title, k)
	}
	csvData[0] = title

	for i, record := range data {
		row := make([]string, len(title))

		for j, k := range title {
			if v, ok := record[k]; ok {
				row[j] = v
			} else {
				row[j] = ""
			}
		}

		csvData[i+1] = row
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.WriteAll(csvData)
	if err != nil {
		return err
	}

	return nil
}

func readCSV(path string) ([]map[string]string, error) {

	// Open the file
	csvfile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer csvfile.Close()

	// Parse the file
	r := csv.NewReader(csvfile)

	data := []map[string]string{}

	var title []string

	// Iterate through the records
	isTitle := true
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if isTitle {
			title = record
			isTitle = false
			continue
		}

		row := make(map[string]string)

		for i := range title {
			row[title[i]] = record[i]
		}

		data = append(data, row)
	}

	return data, nil
}
