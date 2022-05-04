package mysql

import (
	"dkfbasel.ch/covid-evidence/logger"
	"dkfbasel.ch/covid-evidence/stores"
)

// Update ...
func (s *Store) Update(source string, records []*stores.Record) error {
	var stmt string

	if source == "johns-hopkins" {
		stmt = `INSERT INTO johnshopkins_copy
					(country, cases, year, week)
				VALUES
					(:country, :cases, :year, :week)
				ON DUPLICATE KEY UPDATE week = :week;`

	} else if source == "clinicaltrials.gov" {
		logger.Error("clinicaltrials", nil)
		for i := range records {
			for key := range ctDict {
				if _, ok := records[i].Fields[key]; !ok {
					records[i].Fields[key] = nil
				}
			}
		}
		stmt = `
		INSERT INTO screening_clinicaltrials
			(
				id, cove_import_date, cove_update_date,
				nct_id, date_study_first_submitted, date_study_first_posted, date_last_update_posted,
				date_results_first_posted,
				date_started, date_started_type, date_completed, date_completed_type, status,
				brief_title, official_title, brief_summary, detailed_description, study_type, phase,
				allocation, intervention_model, intervention_model_description, primary_purpose,
				masking, ` + "`condition`" + `, intervention_type, intervention_name, intervention_desc,
				eligibility_criteria, gender, minimum_age, maximum_age, healthy_volunteers, enrollment,
				enrollment_type, primary_outcome_measure, primary_outcome_frame, primary_outcome_description,
				secondary_outcome_measure, secondary_outcome_time_frame, secondary_outcome_description,
				arm_group_arm_group_label, arm_group_arm_group_type, arm_group_description, location_name,
				location_city, location_country, patient_data_all, patient_data_sharing_ipd, sponsors_agency,
				sponsors_agency_class, publications_reference, publications_pmid, source_duplicates
			)
		VALUES
			(
				CAST(NULLIF(:id, '') AS UNSIGNED), now(), now(),
				:nct_id, :date_study_first_submitted, :date_study_first_posted, :date_last_update_posted,
				:date_results_first_posted,
				:date_started, :date_started_type, :date_completed, :date_completed_type, :status,
				:brief_title, :official_title, :brief_summary, :detailed_description, :study_type, :phase,
				:allocation, :intervention_model, :intervention_model_description, :primary_purpose,
				:masking, :condition, :intervention_type, :intervention_name, :intervention_desc,
				:eligibility_criteria, :gender, :minimum_age, :maximum_age, :healthy_volunteers, :enrollment,
				:enrollment_type, :primary_outcome_measure, :primary_outcome_frame, :primary_outcome_description,
				:secondary_outcome_measure, :secondary_outcome_time_frame, :secondary_outcome_description,
				:arm_group_arm_group_label, :arm_group_arm_group_type, :arm_group_description, :location_name,
				:location_city, :location_country, :patient_data_all, :patient_data_sharing_ipd, :sponsors_agency,
				:sponsors_agency_class, :publications_reference, :publications_pmid, CAST(:duplicates AS UNSIGNED)
			)
		ON DUPLICATE KEY UPDATE
			cove_update_date = now(), nct_id = :nct_id,
			date_study_first_submitted = :date_study_first_submitted,
			date_study_first_posted = :date_study_first_posted,
			date_results_first_posted = :date_results_first_posted,
			date_last_update_posted = :date_last_update_posted,
			date_started = :date_started, date_started_type = :date_started_type,
			date_completed = :date_completed, date_completed_type = :date_completed_type,
			status = :status, brief_title = :brief_title, official_title = :official_title,
			brief_summary = :brief_summary, detailed_description = :detailed_description,
			study_type = :study_type, phase = :phase, allocation = :allocation,
			intervention_model = :intervention_model,
			intervention_model_description = :intervention_model_description,
			primary_purpose = :primary_purpose, masking = :masking, ` + "`condition` =" + ` :condition,
			intervention_type = :intervention_type, intervention_name = :intervention_name,
			intervention_desc = :intervention_desc, eligibility_criteria = :eligibility_criteria,
			gender = :gender, minimum_age = :minimum_age, maximum_age = :maximum_age,
			healthy_volunteers = :healthy_volunteers, enrollment = :enrollment,
			enrollment_type = :enrollment_type, primary_outcome_measure = :primary_outcome_measure,
			primary_outcome_frame = :primary_outcome_frame, primary_outcome_description = :primary_outcome_description,
			secondary_outcome_measure = :secondary_outcome_measure,
			secondary_outcome_time_frame = :secondary_outcome_time_frame,
			secondary_outcome_description = :secondary_outcome_description,
			arm_group_arm_group_label = :arm_group_arm_group_label, arm_group_arm_group_type = :arm_group_arm_group_type,
			arm_group_description = :arm_group_description, location_name = :location_name,
			location_city = :location_city, location_country = :location_country,
			patient_data_all = :patient_data_all, patient_data_sharing_ipd = :patient_data_sharing_ipd,
			sponsors_agency = :sponsors_agency, sponsors_agency_class = :sponsors_agency_class,
			publications_reference = :publications_reference, publications_pmid = :publications_pmid, 
			source_duplicates = CAST(:duplicates AS UNSIGNED);`
	}
	if source == "ictrp" {
		for i := range records {
			for key, value := range ictrpDict {
				if v, ok := records[i].Fields[value]; ok {
					records[i].Fields[key] = v
				}
			}
			for key := range ictrpDict {
				if _, ok := records[i].Fields[key]; !ok {
					records[i].Fields[key] = nil
				}
			}
		}

		stmt = `
		INSERT INTO screening_ictrp
			(
				id, cove_import_date, cove_update_date,
				trial_id, last_refreshed_on, public_title, scientific_title, acronym, primary_sponsor,
				date_registration, date_registration_3, export_date, source_register, web_address, recruitment_status,
				other_records, inclusion_agemin, inclusion_agemax, inclusion_gender, date_enrollement, target_size,
				study_type, study_design, phase, countries, contact_firstname, contact_lastname, contact_address,
				contact_email, contact_tel, contact_affiliation, inclusion_criteria, exclusion_criteria, ` + "`condition`" + `,
				intervention, primary_outcome, results_date_posted, results_date_completed, results_url_link,
				retrospective_flag, bridging_flag_truefalse, bridged_type, results_yes_no, source_duplicates
			)
		VALUES
			(
				CAST(NULLIF(:id, '') AS UNSIGNED), now(), now(),
				:trial_id, :last_refreshed_on, :public_title, :scientific_title, :acronym, :primary_sponsor,
				:date_registration, :date_registration_3, :export_date, :source_register, :web_address, :recruitment_status,
				:other_records, :inclusion_agemin, :inclusion_agemax, :inclusion_gender, :date_enrollement, :target_size,
				:study_type, :study_design, :phase, :countries, :contact_firstname, :contact_lastname, :contact_address,
				:contact_email, :contact_tel, :contact_affiliation, :inclusion_criteria, :exclusion_criteria, :condition,
				:intervention, :primary_outcome, :results_date_posted, :results_date_completed, :results_url_link,
				:retrospective_flag, :bridging_flag_truefalse, :bridged_type, :results_yes_no, CAST(:duplicates AS UNSIGNED)
			)
		ON DUPLICATE KEY UPDATE
			cove_update_date = now(), trial_id = :trial_id, last_refreshed_on = :last_refreshed_on,
			public_title = :public_title, scientific_title = :scientific_title, acronym = :acronym,
			primary_sponsor = :primary_sponsor, date_registration = :date_registration,
			date_registration_3 = :date_registration_3, export_date = :export_date, source_register = :source_register,
			web_address = :web_address, recruitment_status = :recruitment_status, other_records = :other_records,
			inclusion_agemin = :inclusion_agemin, inclusion_agemax = :inclusion_agemax,
			inclusion_gender = :inclusion_gender, date_enrollement = :date_enrollement, target_size = :target_size,
			study_type = :study_type, study_design = :study_design, phase = :phase, countries = :countries,
			contact_firstname = :contact_firstname, contact_lastname = :contact_lastname, contact_address = :contact_address,
			contact_email = :contact_email, contact_tel = :contact_tel, contact_affiliation = :contact_affiliation,
			inclusion_criteria = :inclusion_criteria, exclusion_criteria = :exclusion_criteria,
			` + "`condition` =" + ` :condition, intervention = :intervention, primary_outcome = :primary_outcome,
			results_date_posted = :results_date_posted, results_date_completed = :results_date_completed,
			results_url_link = :results_url_link, retrospective_flag = :retrospective_flag,
			bridging_flag_truefalse = :bridging_flag_truefalse, bridged_type = :bridged_type,
			results_yes_no = :results_yes_no, source_duplicates = CAST(:duplicates AS UNSIGNED);`
	}
	if source == "love" {
		for i := range records {
			for key := range LoveDict {
				if v, ok := records[i].Fields[key]; ok {
					records[i].Fields[key] = v
					if v == "" {
						records[i].Fields[key] = nil
					}
				} else {
					records[i].Fields[key] = nil
				}
			}
		}

		stmt = `
		INSERT INTO screening_love
			(
				id, created_on, cove_import_date, cove_update_date, cove_screening,
				plasma, registry, ` + "`" + `key` + "`" + `, item_type, publication_year, author,
				title, publication_title, isbn, issn, doi, url,
				abstract_note, date, date_added, date_modified, access_date, pages,
				num_pages, issue, volume, number_of_volumes, journal_abbreviation, short_title, series,
				series_number, series_text, series_title, publisher, place, language,
				rights, type, archive, archive_location, library_catalog,
				call_number, extra, notes, file_attachements, link_attachments, manual_tags,
				automatic_tags, editor, series_editor, translator, contributor, attorney_agent,
				book_author, cast_member, commenter, composer, counsel, interviewer, producer,
				recipient, reviewed_author, scriptwriter, words_by, guest, number, edition,
				running_time, scale, medium, artwork_size, filing_date, application_number,
				assignee, issuing_authority, country, meeting_name, conference_name, court,
				` + "`" + `references` + "`" + `, reporter, legal_status, priority_numbers, programming_language, version,
				system, code, code_number, section, session, committee, history, legislative_body
			)
		VALUES
			(
				CAST(NULLIF(:id, '') AS UNSIGNED), now(), now(), now(), cove_screening,
				:plasma, :registry, :key, :item_type, :publication_year, :author,
				:title, :publication_title, :isbn, :issn, :doi, :url,
				:abstract_note, :date, :date_added, :date_modified, :access_date, :pages,
				:num_pages, :issue, :volume, :number_of_volumes, :journal_abbreviation, :short_title, :series,
				:series_number, :series_text, :series_title, :publisher, :place, :language,
				:rights, :type, :archive, :archive_location, :library_catalog,
				:call_number, :extra, :notes, :file_attachements, :link_attachments, :manual_tags,
				:automatic_tags, :editor, :series_editor, :translator, :contributor, :attorney_agent,
				:book_author, :cast_member, :commenter, :composer, :counsel, :interviewer, :producer,
				:recipient, :reviewed_author, :scriptwriter, :words_by, :guest, :number, :edition,
				:running_time, :scale, :medium, :artwork_size, :filing_date, :application_number,
				:assignee, :issuing_authority, :country, :meeting_name, :conference_name, :court,
				:references, :reporter, :legal_status, :priority_numbers, :programming_language, :version,
				:system, :code, :code_number, :section, :session, :committee, :history, :legislative_body
			)
		ON DUPLICATE KEY UPDATE
			cove_update_date = now(), plasma = :plasma, registry = :registry,
			cove_screening = :cove_screening,
			` + "`" + `key` + "`" + ` = :key, item_type = :item_type, publication_year = :publication_year,
			author = :author, title = :title,
			publication_title = :publication_title, isbn = :isbn, issn = :issn,
			doi = :doi, url = :url, abstract_note = :abstract_note,
			date = :date, date_added = :date_added,
			date_modified = :date_modified, access_date = :access_date, pages = :pages,
			num_pages = :num_pages, issue = :issue, volume = :volume, number_of_volumes = :number_of_volumes,
			journal_abbreviation = :journal_abbreviation, short_title = :short_title, series = :series,
			series_number = :series_number, series_text = :series_text, series_title = :series_title,
			publisher = :publisher, place = :place,
			language = :language, rights = :rights, type = :type,
			archive = :archive, archive_location = :archive_location,
			library_catalog = :library_catalog, call_number = :call_number,
			extra = :extra, notes = :notes,
			file_attachements = :file_attachements, link_attachments = :link_attachments,
			manual_tags = :manual_tags, automatic_tags = :automatic_tags, editor = :editor,
			series_editor = :series_editor, translator = :translator, contributor = :contributor,
			attorney_agent = :attorney_agent, book_author = :book_author, cast_member = :cast_member,
			commenter = :commenter, composer = :composer, counsel = :counsel, interviewer = :interviewer,
			producer = :producer, recipient = :recipient, reviewed_author = :reviewed_author,
			scriptwriter = :scriptwriter, words_by = :words_by, guest = :guest, number = :number,
			edition = :edition, running_time = :running_time, scale = :scale, medium = :medium,
			artwork_size = :artwork_size, filing_date = :filing_date, application_number = :application_number,
			assignee = :assignee, issuing_authority = :issuing_authority, country = :country,
			meeting_name = :meeting_name, conference_name = :conference_name, court = :court,
			` + "`" + `references` + "`" + ` = :references, reporter = :reporter, legal_status = :legal_status,
			priority_numbers = :priority_numbers, programming_language = :programming_language,
			version = :version, system = :system, code = :code, code_number = :code_number,
			section = :section, session = :session, committee = :committee, ` + "`" + `history` + "`" + ` = :history,
			legislative_body = :legislative_body;`
	}
	alreadyIn := 0
	notIn := 0
	for _, record := range records {
		if record.Field("id") == "" {
			notIn++
		} else {
			alreadyIn++
		}

		if !s.DryRun {
			_, err := s.DB.NamedExec(stmt, record.Fields)
			if err != nil {
				logger.Error("update screening", err)
				return err
			}
		}
	}

	return nil
}
