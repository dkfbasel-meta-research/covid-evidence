package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"dkfbasel.ch/covid-evidence/logger"

	"dkfbasel.ch/covid-evidence/pipeline"
	"dkfbasel.ch/covid-evidence/sources"

	"dkfbasel.ch/covid-evidence/sources/clinicaltrials"
	"dkfbasel.ch/covid-evidence/sources/ictrp"

	"dkfbasel.ch/covid-evidence/sources/johnshopkins"
	"dkfbasel.ch/covid-evidence/sources/love"
	"dkfbasel.ch/covid-evidence/stores/mysql"
)

func main() {

	var mysqlConnectionString string
	flag.StringVar(&mysqlConnectionString, "mysql", "", "mysql connection string [mandatory]")

	var tmpSource string
	flag.StringVar(&tmpSource, "tmp", "./tmp", "path to a directory to save temporary files")

	var ictrpSource string
	flag.StringVar(&ictrpSource, "ictrp", "", "path to ICTRP CSV file to import data")

	var loveSource string
	flag.StringVar(&loveSource, "love", "", "path to love CSV file to import data")

	var johnshopkinsSource bool
	flag.BoolVar(&johnshopkinsSource, "jh", false, "download COVID-19 cases from the john hopkins university")

	var clinicaltrialsSource bool
	flag.BoolVar(&clinicaltrialsSource, "ct", false, "download and import registry entries from clinicaltrials")

	var steps string
	flag.StringVar(&steps, "steps", "all", "One of:\n- \"screening\": import from source to screening table\n- \"cove\": import entries from screening table to cove basic table\n-\"all\": do both steps decribed above (default)")

	flag.Parse()

	if mysqlConnectionString == "" {
		flag.PrintDefaults()
		log.Fatal("MySQL connection string is mandatory!")
	}

	tmpPath, err := filepath.Abs(tmpSource)
	if err != nil {
		log.Fatalf("could not find tmp directory: %s", tmpSource)
	} else {
		log.Printf("save temporary files in the directory: %s", tmpSource)
	}

	screeningStep := steps == "screening" || steps == "all" || strings.Contains(steps, "screening")
	coveBasicStep := steps == "cove" || steps == "all" || strings.Contains(steps, "cove")

	// setup database
	db := &mysql.Store{}
	db.Setup(mysqlConnectionString, false)
	defer db.Close()

	// Initialize sources
	sources := []sources.Source{}
	msg := []string{}
	if clinicaltrialsSource == true {
		msg = append(msg, "clinicaltrials.gov")
		sources = append(sources, clinicaltrials.NewSource(db))
	}
	if ictrpSource != "" {
		msg = append(msg, "ICTRP")
		ictrpPath, err := filepath.Abs(ictrpSource)
		if err != nil {
			log.Fatal("could not find ICTRP CSV file")
		}
		sources = append(sources, ictrp.NewSource(db, ictrpPath))
	}
	if loveSource != "" {
		msg = append(msg, "LOVE Database")
		lovePath, err := filepath.Abs(loveSource)
		if err != nil {
			log.Fatal("could not find LOVE CSV file")
		}
		sources = append(sources, love.NewSource(db, lovePath))
	}
	if johnshopkinsSource {
		msg = append(msg, "Johns Hopkins")
		sources = append(sources, johnshopkins.NewSource(db))
	}

	if len(msg) > 0 {
		logger.Info(fmt.Sprintf("Import data from: %s", strings.Join(msg, ", ")))
	}

	pipeline := pipeline.NewPipeline(db, sources, tmpPath)
	pipeline.Start(screeningStep, coveBasicStep)
}
