package mysql

import (
	"dkfbasel.ch/covid-evidence/logger"
	"dkfbasel.ch/covid-evidence/stores"
)

// UpdateBasic ...
func (s *Store) UpdateBasic(records []*stores.Record) error {

	stmt := `
	UPDATE cove_basic
	SET
		is_covid = :is_covid, is_trial = :is_trial, is_observational = :is_observational, is_rct = :is_rct,
		is_duplicate = :is_duplicate, review_status = :review_status, reviewer_name = :reviewer_name,
		source = :source, source_id = :source_id, entry_type = :entry_type,
		entry_type_certainty = :entry_type_certainty, url = :url, url_certainty = :url_certainty,
		title = :title, title_certainty = :title_certainty, abstract = :abstract,
		abstract_certainty = :abstract_certainty, authors = :authors, authors_certainty = :authors_certainty,
		corresp_author_lastname = :corresp_author_lastname, corresp_author_lastname_certainty = :corresp_author_lastname_certainty,
		corresp_author_email = :corresp_author_email, corresp_author_email_certainty = :corresp_author_email_certainty,
		journal = :journal, journal_certainty = :journal_certainty, doi = :doi, doi_certainty = :doi_certainty,
		status = :status, status_certainty = :status_certainty, status_date = :status_date,
		status_date_certainty = :status_date_certainty, country = :country,
		country_certainty = :country_certainty, countries = :countries, countries_certainty = :countries_certainty,
		continents = :continents, continents_certainty = :continents_certainty, international = :international,
		international_certainty = :international_certainty,
		registration_date = :registration_date, registration_date_certainty = :registration_date_certainty,
		covid_status = :covid_status, covid_status_certainty = :covid_status_certainty,
		randomized = :randomized, randomized_certainty = :randomized_certainty,  blinding = :blinding,
		blinding_certainty = :blinding_certainty, longitudinal_structure = :longitudinal_structure,
		longitudinal_structure_certainty = :longitudinal_structure_certainty,
		n_arms = NULLIF(:n_arms, ''), n_arms_certainty = :n_arms_certainty, n_enrollment = NULLIF(:n_enrollment, ''),
		n_enrollment_certainty = :n_enrollment_certainty, population_condition = :population_condition,
		population_condition_certainty = :population_condition_certainty, population_gender = :population_gender,
		population_gender_certainty = :population_gender_certainty, population_age = :population_age,
		population_age_certainty = :population_age_certainty,
		intervention_substance = NULLIF(:intervention_substance, ''), intervention_type = NULLIF(:intervention_type, ''),
		intervention_type_certainty = :intervention_type_certainty,
		intervention_name = :intervention_name, intervention_name_certainty = :intervention_name_certainty,
		control_type = NULLIF(:control_type, ''), control_type_certainty = :control_type_certainty,
		control = :control, control_certainty = :control_certainty,
		out_primary_measure = :out_primary_measure, out_primary_desc = :out_primary_desc,
		out_primary_timeframe = :out_primary_timeframe, out_primary_measure_certainty = :out_primary_measure_certainty,
		out_primary_desc_certainty = :out_primary_desc_certainty, out_primary_timeframe_certainty = :out_primary_timeframe_certainty,
		start_date = :start_date, start_date_certainty = :start_date_certainty,
		end_date = :end_date, end_date_certainty = :end_date_certainty, results_available = NULLIF(:results_available, ''),
		results_available_certainty = :results_available_certainty, results_expected_date = :results_expected_date,
		results_expected_date_certainty = :results_expected_date_certainty, ipd_sharing = :ipd_sharing, ipd_sharing_certainty = :ipd_sharing_certainty,
		linkage_info = :linkage_info, linkage_info_certainty = :linkage_info_certainty,
		out_secondary_measure = :out_secondary_measure, out_secondary_desc = :out_secondary_desc, out_secondary_timeframe = :out_secondary_timeframe,
		out_secondary_measure_certainty = :out_secondary_measure_certainty,
		out_secondary_desc_certainty = :out_secondary_desc_certainty, out_secondary_timeframe_certainty = :out_secondary_timeframe_certainty,
		inclusion_criteria = :inclusion_criteria, inclusion_criteria_certainty = :inclusion_criteria_certainty,
		exclusion_criteria = :exclusion_criteria, exclusion_criteria_certainty = :exclusion_criteria_certainty,
		funding = :funding, funding_certainty = :funding_certainty,
		funding_type = :funding_type, funding_type_certainty = :funding_type_certainty, protocol_available = :protocol_available,
		protocol_available_certainty = :protocol_available_certainty, protocol_link = :protocol_link, protocol_link_certainty = :protocol_link_certainty,
		extraction_comment = :extraction_comment
	WHERE source = :source AND source_id = :source_id;`

	for i := range records {

		for k := range coveDict {
			if _, ok := records[i].Fields[k]; !ok {
				records[i].Fields[k] = ""
			}
		}

		if i%10 == 0 {
			logger.Info("update basic record", logger.Any("#", i))
		}

		if !s.DryRun {
			_, err := s.DB.NamedExec(stmt, records[i].Fields)
			if err != nil {
				logger.NewError("could not update record", err,
					logger.String("source", records[i].Field("source")),
					logger.String("source_id", records[i].Field("source_id")))
			}
		}
	}

	return nil
}
