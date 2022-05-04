package sources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type coveBasic struct {
	CoveID              int    `json:"ninox_nr"`
	Source              string `json:"source"`
	ReviewStatus        string `json:"review_status"`
	ResultsAvailable    string `json:"results_available"`
	IpdSharing          string `json:"ipd_sharing"`
	InterventionType    string `json:"intervention_type"`
	InterventionName    string `json:"intervention_name"`
	NumberEnrollment    int    `json:"n_enrollment"`
	Country             string `json:"country"`
	Status              string `json:"status"`
	Randomized          string `json:"randomized"`
	NumberArms          int    `json:"n_arms"`
	Blinding            string `json:"blinding"`
	PopulationCondition string `json:"population_condition"`
	Control             string `json:"control"`
	OutPrimaryMeasure   string `json:"out_primary_measure"`
	StartDate           string `json:"start_date"`
	EndDate             string `json:"end_date"`
	SourceID            string `json:"source_id"`
	Title               string `json:"title"`
	URL                 string `json:"url"`
	IsCovid             string `json:"is_covid"`
	IsTrial             string `json:"is_trial"`
	IsObservational     string `json:"is_observational"`
	IsDuplicate         bool   `json:"is_duplicate,-"`
}

const coveURL = "https://test.test"

// CoveBasicToJSONP will write a jsonp file to automatically update the online database
func CoveBasicToJSONP() error {

	// Get the data
	resp, err := http.Get(coveURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var dta []coveBasic
	err = json.Unmarshal(content, &dta)
	if err != nil {
		log.Fatalf("could not parse data: %+v", err)
	}

	var filtered []*coveBasic

	covidFiltered := 0
	trialFiltered := 0
	duplicateFiltered := 0
	sourceFiltered := 0
	idFiltered := 0

	// filter out all data that is not a trial
	for i, item := range dta {

		// skip all non covid items
		if item.IsCovid == "no" {
			covidFiltered++
			continue
		}

		// skip all non trial items
		if item.IsTrial == "no" {
			trialFiltered++
			continue
		}

		// skip all duplicates
		if item.IsDuplicate {
			duplicateFiltered++
			continue
		}

		if item.Source == "" {
			sourceFiltered++
			continue
		}

		// skip all items without covid id
		if item.CoveID == 0 {
			idFiltered++
			continue
		}

		filtered = append(filtered, &dta[i])
	}

	output, err := json.Marshal(&filtered)
	if err != nil {
		return err
	}
	outputString := strings.ReplaceAll(string(output), "\"ninox_nr\"", "\"cove_id\"")
	outputString = fmt.Sprintf("window.getCoveBasicData = function() { return %s; }", outputString)

	err = ioutil.WriteFile("latest.js", []byte(outputString), 0644)
	if err != nil {
		return err
	}

	return nil
}
