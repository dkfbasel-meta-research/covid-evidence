package ictrp

import (
	"io/ioutil"
	"os"
	"path"

	"dkfbasel.ch/covid-evidence/logger"
)

const requestURL = "https://worldhealthorg-my.sharepoint.com/personal/karamg_who_int/_layouts/15/download.aspx?UniqueId=20b68817-2dd5-472e-b923-4e541db55510"

// Fetch will fetch data from https://www.who.int/ictrp/en/
func (who *ICTRP) Fetch() (string, error) {

	filepath := path.Join(who.Path)

	// check if import file exists
	_, err := os.Stat(filepath)
	if err != nil && os.IsNotExist(err) {
		logger.NewError("file does not exist", err)
		return "", nil
	}
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
