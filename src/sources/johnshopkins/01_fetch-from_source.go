package johnshopkins

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"dkfbasel.ch/covid-evidence/logger"
)

// Fetch will fetch data from github created by the
// Johns Hopkins University
func (jh *JohnsHopkins) Fetch() (string, error) {

	// prepare the api url
	apiUrl, err := url.Parse(requestURL)
	if err != nil {
		return "", fmt.Errorf("could not parse url: %w", err)
	}

	request, err := http.Get(apiUrl.String())
	if err != nil {
		return "", logger.NewError("could not fetch data", err)
	}

	content, err := ioutil.ReadAll(request.Body)
	if err != nil {
		return "", logger.NewError("could not read response", err)
	}
	request.Body.Close()

	return string(content), nil
}
