package mysql

import (
	"dkfbasel.ch/covid-evidence/logger"
	"dkfbasel.ch/covid-evidence/stores"
)

// AddBasic ...
func (s *Store) AddBasic(records []*stores.Record) error {

	stmt := `
	INSERT INTO cove_basic
		(
			created_on,
			is_covid, is_trial, is_observational, is_duplicate, is_rct,
			review_status, reviewer_name, source, source_id,
			entry_type, entry_type_certainty, url, url_certainty, title,
			title_certainty, abstract, abstract_certainty, authors, authors_certainty,
			corresp_author_lastname, corresp_author_lastname_certainty, corresp_author_email,
			corresp_author_email_certainty, journal, journal_certainty, doi, doi_certainty,
			status, status_certainty, status_date, status_date_certainty, country,
			country_certainty, countries, countries_certainty, continents, continents_certainty,
			international, international_certainty, registration_date, registration_date_certainty,
			covid_status, covid_status_certainty, randomized, randomized_certainty,
			blinding, blinding_certainty, longitudinal_structure, longitudinal_structure_certainty,
			n_arms, n_arms_certainty, n_enrollment, n_enrollment_certainty, population_condition,
			population_condition_certainty, population_gender, population_gender_certainty,
			population_age, population_age_certainty, intervention_substance, intervention_substance_certainty,
			intervention_type, intervention_type_certainty, intervention_name, intervention_name_certainty,
			control_type, control_type_certainty, control, control_certainty,
			out_primary_measure, out_primary_desc, out_primary_timeframe, out_primary_measure_certainty,
			out_primary_desc_certainty, out_primary_timeframe_certainty, start_date, start_date_certainty,
			end_date, end_date_certainty, results_available, results_available_certainty, results_expected_date,
			results_expected_date_certainty, ipd_sharing, ipd_sharing_certainty, linkage_info, linkage_info_certainty,
			out_secondary_measure, out_secondary_desc, out_secondary_timeframe, out_secondary_measure_certainty,
			out_secondary_desc_certainty, out_secondary_timeframe_certainty, inclusion_criteria,
			inclusion_criteria_certainty, exclusion_criteria, exclusion_criteria_certainty, funding, funding_certainty,
			funding_type, funding_type_certainty, protocol_available, protocol_available_certainty,
			protocol_link, protocol_link_certainty, extraction_comment
		)
	VALUES
		(
			now(),
			:is_covid, :is_trial, :is_observational, :is_duplicate, :is_rct,
			:review_status, :reviewer_name, :source, :source_id,
			:entry_type, :entry_type_certainty, :url, :url_certainty, :title,
			:title_certainty, :abstract, :abstract_certainty, :authors, :authors_certainty,
			:corresp_author_lastname, :corresp_author_lastname_certainty, :corresp_author_email,
			:corresp_author_email_certainty, :journal, :journal_certainty, :doi, :doi_certainty,
			:status, :status_certainty, :status_date, :status_date_certainty, :country,
			:country_certainty, :countries, :countries_certainty, :continents, :continents_certainty,
			:international, :international_certainty, :registration_date, :registration_date_certainty,
			:covid_status, :covid_status_certainty, :randomized, :randomized_certainty,
			:blinding, :blinding_certainty, :longitudinal_structure, :longitudinal_structure_certainty,
			NULLIF(:n_arms, ''), :n_arms_certainty, NULLIF(:n_enrollment, ''), :n_enrollment_certainty, :population_condition,
			:population_condition_certainty, :population_gender, :population_gender_certainty,
			:population_age, :population_age_certainty, NULLIF(:intervention_substance, ''), :intervention_substance_certainty,
			NULLIF(:intervention_type, ''), :intervention_type_certainty, :intervention_name, :intervention_name_certainty,
			NULLIF(:control_type, ''), :control_type_certainty, :control, :control_certainty,
			:out_primary_measure, :out_primary_desc, :out_primary_timeframe, :out_primary_measure_certainty,
			:out_primary_desc_certainty, :out_primary_timeframe_certainty, :start_date, :start_date_certainty,
			:end_date, :end_date_certainty, NULLIF(:results_available, ''), :results_available_certainty, :results_expected_date,
			:results_expected_date_certainty, :ipd_sharing, :ipd_sharing_certainty, :linkage_info, :linkage_info_certainty,
			:out_secondary_measure, :out_secondary_desc, :out_secondary_timeframe, :out_secondary_measure_certainty,
			:out_secondary_desc_certainty, :out_secondary_timeframe_certainty, :inclusion_criteria,
			:inclusion_criteria_certainty, :exclusion_criteria, :exclusion_criteria_certainty, :funding, :funding_certainty,
			:funding_type, :funding_type_certainty, :protocol_available, :protocol_available_certainty,
			:protocol_link, :protocol_link_certainty, :extraction_comment
		);`

	for i := range records {

		for k := range coveDict {
			if _, ok := records[i].Fields[k]; !ok {
				records[i].Fields[k] = ""
			}
		}

		if i%10 == 0 {
			logger.Info("add to basic", logger.Any("#", i))
		}

		if !s.DryRun {
			_, err := s.DB.NamedExec(stmt, records[i].Fields)
			if err != nil {
				logger.NewError("could not add record", err,
					logger.String("id", records[i].Field("id")),
					logger.String("source", records[i].Field("source")),
					logger.String("source_id", records[i].Field("source_id")))
			}
		}
	}

	return nil
}
