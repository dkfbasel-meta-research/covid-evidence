package sources

import (
	"fmt"
	"sort"
	"strings"

	"dkfbasel.ch/covid-evidence/logger"
)

func CountryFn(m string) (interface{}, bool) {
	splitByComma := strings.Contains(m, ", ")
	splitByComma = !splitByComma && strings.Contains(m, ",")
	splittedValues := []string{}
	if splitByComma {
		splittedValues = strings.Split(m, ",")
	} else {
		splittedValues = strings.Split(m, ";")
	}
	countryList := make(map[string]bool)
	countriesMap := make(map[string]bool)
	for _, country := range splittedValues {
		country = strings.TrimSpace(country)
		if country == "" {
			continue
		}

		if _, ok := countriesMap[country]; ok {
			continue
		}
		for _, termMap := range countryMap {
			if termMap.Term == country {
				if termMap.Country {
					countryList[termMap.Mapping] = true
					countriesMap[country] = true
				}
			}
		}

		if _, ok := countriesMap[country]; !ok {
			logger.Info("Not in list", logger.String("country", country))
		}
	}
	if len(countryList) == 0 {
		return "", true
	}

	countriesArray := []string{}
	for country := range countryList {
		countriesArray = append(countriesArray, country)
	}
	sort.Strings(countriesArray)
	countries := strings.Join(countriesArray, ";")

	return fmt.Sprintf(";%s;", countries), true
}

func InternationalFn(m string) (interface{}, bool) {
	splitByComma := strings.Contains(m, ", ")
	splitByComma = !splitByComma && strings.Contains(m, ",")
	splittedValues := []string{}
	if splitByComma {
		splittedValues = strings.Split(m, ",")
	} else {
		splittedValues = strings.Split(m, ";")
	}
	continentList := make(map[string]bool)
	countryList := make(map[string]bool)
	continentMap := make(map[string]bool)
	for _, country := range splittedValues {
		country = strings.TrimSpace(country)
		if country == "" {
			continue
		}

		if _, ok := continentMap[country]; ok {
			continue
		}

		for _, termMap := range countryMap {
			if termMap.Term == country {
				if termMap.Continent != "" {
					continentList[termMap.Continent] = true
				}
				if termMap.Country {
					countryList[termMap.Mapping] = true
				}

				continentMap[country] = true
			}
		}
	}
	if len(countryList) == 0 && len(continentList) == 0 {
		return "", true
	}

	if len(continentList) == 1 && len(countryList) == 1 {
		return "No", true
	}

	return "Yes", true
}

func ContinentFn(m string) (interface{}, bool) {
	splitByComma := strings.Contains(m, ", ")
	splitByComma = !splitByComma && strings.Contains(m, ",")
	splittedValues := []string{}
	if splitByComma {
		splittedValues = strings.Split(m, ",")
	} else {
		splittedValues = strings.Split(m, ";")
	}
	continentList := make(map[string]bool)
	continentMap := make(map[string]bool)
	for _, country := range splittedValues {
		country = strings.TrimSpace(country)
		if country == "" {
			continue
		}

		if _, ok := continentMap[country]; ok {
			continue
		}

		for _, termMap := range countryMap {
			if termMap.Term == country {
				if termMap.Continent != "" {
					continentList[termMap.Continent] = true
					continentMap[country] = true
				}
			}
		}
	}
	if len(continentList) == 0 {
		return "", true
	}

	continentsArray := []string{}
	for continent := range continentList {
		continentsArray = append(continentsArray, continent)
	}
	sort.Strings(continentsArray)
	continents := strings.Join(continentsArray, ";")

	if continents == m {
		return continents, false
	}

	return fmt.Sprintf(";%s;", continents), true
}
