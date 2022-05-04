package love

import (
	"crypto/md5"
	"fmt"
	"strings"

	"dkfbasel.ch/covid-evidence/logger"
)

type dbTrial struct {
	SourceID       string
	TrialID        string
	RegistrySearch []string
}

// UpdateScreening will update the screening variable in the screening table
func (love *Love) UpdateScreening() error {

	screeningRecords, err := love.Store.FetchScreening("love", nil, "cove_screening LIKE '%incl%' OR cove_screening LIKE '%preprint%'")
	if err != nil {
		return logger.NewError("could not fetch screening", err)
	}

	trials := []dbTrial{}
	for _, screeningRecord := range screeningRecords {

		registryString := strings.TrimSpace(screeningRecord.Field("registry"))
		if registryString == "" || registryString == "na" {
			continue
		}

		registryTrials := strings.Split(registryString, ";")
		for _, trialList := range registryTrials {

			currentTrial := dbTrial{}
			currentTrial.SourceID = screeningRecord.Field("url")
			trialString := currentTrial.SourceID + trialList
			currentTrial.TrialID = fmt.Sprintf("%x", md5.Sum([]byte(trialString)))
			currentTrial.RegistrySearch = []string{}

			registryIDs := strings.Split(trialList, ",")
			for _, registryID := range registryIDs {
				registrationID := strings.TrimSpace(registryID)
				if registrationID == "" {
					continue
				}

				currentTrial.RegistrySearch = append(currentTrial.RegistrySearch, registrationID)
			}

			trials = append(trials, currentTrial)
		}
	}

	// first clean trials, remove all trials that is not in the list
	loveID := []string{}
	for _, trial := range trials {
		loveID = append(loveID, trial.TrialID)
	}
	existingTrials, err := love.Store.CleanTrials(loveID)
	if err != nil {
		return logger.NewError("could not clean trial", err)
	}

	// filter trials
	trialsToInsert := []dbTrial{}
	for _, trial := range trials {
		trialAlreadyInDB := false
		for _, trialID := range existingTrials {
			if trial.TrialID == trialID {
				trialAlreadyInDB = true
				break
			}
		}

		if !trialAlreadyInDB {
			trialsToInsert = append(trialsToInsert, trial)
		}
	}

	// insert all new trials
	for _, trial := range trialsToInsert {
		err := love.Store.AddTrial(trial.TrialID, trial.SourceID, trial.RegistrySearch)
		if err != nil {
			return logger.NewError("could not insert trial", err)
		}
	}

	logger.Info("all trials inserted")

	return nil
}
