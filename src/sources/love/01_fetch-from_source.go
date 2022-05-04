package love

import (
	"io/ioutil"
	"os"
	"path"

	"dkfbasel.ch/covid-evidence/logger"
)

// Fetch will have nothing to do, because, it is done manually
func (love *Love) Fetch() (string, error) {

	filepath := path.Join(love.Path, "love.csv")

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
