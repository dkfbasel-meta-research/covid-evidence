package clinicaltrials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"dkfbasel.ch/covid-evidence/logger"
)

type fetchResponse struct {
	FullStudiesResponse struct {
		APIVrs           string
		DataVrs          string
		Expression       string
		NStudiesAvail    int
		NStudiesFound    int
		MinRank          int
		MaxRank          int
		NStudiesReturned int
		FullStudies      []json.RawMessage
	}
}

// Fetch will fetch data from clinicaltrials.gov with the given query paramenter
// and return the data as string matrix (i.e. for csv tables)
func (ct *ClinicalTrials) Fetch() (string, error) {

	// prepare the api url
	apiUrl, err := url.Parse(requestURL)
	if err != nil {
		return "", fmt.Errorf("could not parse url: %w", err)
	}

	// add additional query params
	params := url.Values{}
	params.Add("expr", query)
	params.Add("fmt", "JSON")

	// save all studies exported from clinicaltrials.gov
	studies := []json.RawMessage{}

	// do not make more then 100 fetch attempts (just to ensure that we do not,
	// overload the clinicaltrials.gov api if something goes wrong on our side)
	currentFetchAttempt := 0
	maxFetchAttempts := 100

	// initialize variables to select the min and max ranking of the results
	// (required for pagination of the clinicaltrials.gov results)
	var currentRank, maxRank int

	for {

		// fetch the next batch of studies (maximum 100 per batch)
		currentRank = currentFetchAttempt*100 + 1
		maxRank = currentRank + 99
		params.Set("min_rnk", strconv.Itoa(currentRank))
		params.Set("max_rnk", strconv.Itoa(maxRank))

		// encode the url parameters
		apiUrl.RawQuery = params.Encode()

		// perform the request and parse the response
		response, err := performRequest(apiUrl.String())
		if err != nil {
			return "", fmt.Errorf("could not fetch data: %s, %w", apiUrl.String(), err)
		}

		// append the results to the previously fetched studies
		for _, study := range response.FullStudiesResponse.FullStudies {
			studies = append(studies, study)
		}

		// log some information
		logger.Debug("MaxRank", logger.Any("rank", response.FullStudiesResponse.MaxRank),
			logger.Any("studies", response.FullStudiesResponse.NStudiesFound))

		// stop fetching if the rank is higher
		if response.FullStudiesResponse.MaxRank >= response.FullStudiesResponse.NStudiesFound {
			break
		}

		currentFetchAttempt++
		if currentFetchAttempt > maxFetchAttempts {
			break
		}

	}

	// convert the information to json
	asJSON, err := json.Marshal(studies)
	if err != nil {
		return "", logger.NewError("could not marshal studies", err)
	}

	return string(asJSON), nil
}

// performRequest will perform the request for clinicaltrials.gov and return
// the parsed results
func performRequest(url string) (*fetchResponse, error) {

	request, err := http.Get(url)
	if err != nil {
		return nil, logger.NewError("could not fetch data", err)
	}

	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return nil, logger.NewError("could not read response", err)
	}
	request.Body.Close()

	var response fetchResponse
	err = json.Unmarshal(content, &response)
	if err != nil {
		return nil, logger.NewError("could not parse response", err)
	}

	return &response, nil

}
