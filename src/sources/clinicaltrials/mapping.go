package clinicaltrials

var fieldMap = []struct {
	Ninox  string
	Name   string
	Search string
}{
	{"nct_id", "NCTId", "ProtocolSection.IdentificationModule.NCTId"},
	{"date_study_first_submitted", "StudyFirstSubmitDate", "ProtocolSection.StatusModule.StudyFirstSubmitDate"},
	{"date_study_first_posted", "StudyFirstPostDate", "ProtocolSection.StatusModule.StudyFirstPostDateStruct.StudyFirstPostDate"},
	{"date_results_first_posted", "ResultsFirstPostDate", "ProtocolSection.StatusModule.ResultsFirstPostDateStruct.ResultsFirstPostDate"},
	{"", "StudyFirstPostDateType", "ProtocolSection.StatusModule.StudyFirstPostDateStruct.StudyFirstPostDateType"},
	{"date_last_update_posted", "LastUpdatePostDate", "ProtocolSection.StatusModule.LastUpdatePostDateStruct.LastUpdatePostDate"},
	{"", "LastUpdatePostDateType", "ProtocolSection.StatusModule.LastUpdatePostDateStruct.LastUpdatePostDateType"},
	{"date_started", "StartDate", "ProtocolSection.StatusModule.StartDateStruct.StartDate"},
	{"date_started_type", "StartDateType", "ProtocolSection.StatusModule.StartDateStruct.StartDateType"},
	{"date_completed", "CompletionDate", "ProtocolSection.StatusModule.CompletionDateStruct.CompletionDate"},
	{"date_completed_type", "CompletionDateType", "ProtocolSection.StatusModule.CompletionDateStruct.CompletionDateType"},
	{"status", "OverallStatus", "ProtocolSection.StatusModule.OverallStatus"},
	{"brief_title", "BriefTitle", "ProtocolSection.IdentificationModule.BriefTitle"},
	{"official_title", "OfficialTitle", "ProtocolSection.IdentificationModule.OfficialTitle"},
	{"brief_summary", "BriefSummary", "ProtocolSection.DescriptionModule.BriefSummary"},
	{"detailed_description", "DetailedDescription", "ProtocolSection.DescriptionModule.DetailedDescription"},
	{"study_type", "StudyType", "ProtocolSection.DesignModule.StudyType"},
	{"phase", "Phase", "ProtocolSection.DesignModule.PhaseList.Phase"},
	{"allocation", "DesignAllocation", "ProtocolSection.DesignModule.DesignInfo.DesignAllocation"},
	{"intervention_model", "DesignInterventionModel", "ProtocolSection.DesignModule.DesignInfo.DesignInterventionModel"},
	{"intervention_model_description", "DesignInterventionModelDescription", "ProtocolSection.DesignModule.DesignInfo.DesignInterventionModelDescription"},
	{"primary_purpose", "DesignPrimaryPurpose", "ProtocolSection.DesignModule.DesignInfo.DesignPrimaryPurpose"},
	{"masking", "DesignMasking", "ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignMasking"},
	{"", "DesignMaskingDescription", "ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignMaskingDescription"},
	{"", "DesignWhoMasked", "ProtocolSection.DesignModule.DesignInfo.DesignMaskingInfo.DesignWhoMaskedList.DesignWhoMasked"},
	{"condition", "Condition", "ProtocolSection.ConditionsModule.ConditionList.Condition"},
	{"", "Keyword", "ProtocolSection.ConditionsModule.KeywordList.Keyword"},
	{"intervention_type", "InterventionType", "ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention.#.InterventionType"},
	{"intervention_name", "InterventionName", "ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention.#.InterventionName"},
	{"intervention_desc", "InterventionDescription", "ProtocolSection.ArmsInterventionsModule.InterventionList.Intervention.#.InterventionDescription"},
	{"eligibility_criteria", "EligibilityCriteria", "ProtocolSection.EligibilityModule.EligibilityCriteria"},
	{"gender", "Gender", "ProtocolSection.EligibilityModule.Gender"},
	{"minimum_age", "MinimumAge", "ProtocolSection.EligibilityModule.MinimumAge"},
	{"maximum_age", "MaximumAge", "ProtocolSection.EligibilityModule.MaximumAge"},
	{"", "StdAge", "ProtocolSection.EligibilityModule.StdAgeList"},
	{"healthy_volunteers", "HealthyVolunteers", "ProtocolSection.EligibilityModule.HealthyVolunteers"},
	{"enrollment", "EnrollmentCount", "ProtocolSection.DesignModule.EnrollmentInfo.EnrollmentCount"},
	{"enrollment_type", "EnrollmentType", "ProtocolSection.DesignModule.EnrollmentInfo.EnrollmentType"},
	{"primary_outcome_measure", "PrimaryOutcomeMeasure", "ProtocolSection.OutcomesModule.PrimaryOutcomeList.PrimaryOutcome.#.PrimaryOutcomeMeasure"},
	{"primary_outcome_description", "PrimaryOutcomeDescription", "ProtocolSection.OutcomesModule.PrimaryOutcomeList.PrimaryOutcome.#.PrimaryOutcomeDescription"},
	{"primary_outcome_time_frame", "PrimaryOutcomeTimeFrame", "ProtocolSection.OutcomesModule.PrimaryOutcomeList.PrimaryOutcome.#.PrimaryOutcomeTimeFrame"},
	{"secondary_outcome_measure", "SecondaryOutcomeMeasure", "ProtocolSection.OutcomesModule.SecondaryOutcomeList.SecondaryOutcome.#.SecondaryOutcomeMeasure"},
	{"secondary_outcome_description", "SecondaryOutcomeDescription", "ProtocolSection.OutcomesModule.SecondaryOutcomeList.SecondaryOutcome.#.SecondaryOutcomeDescription"},
	{"secondary_outcome_time_frame", "SecondaryOutcomeTimeFrame", "ProtocolSection.OutcomesModule.SecondaryOutcomeList.SecondaryOutcome.#.SecondaryOutcomeTimeFrame"},
	{"arm_group_arm_group_label", "ArmGroupLabel", "ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup.#.ArmGroupLabel"},
	{"arm_group_arm_group_type", "ArmGroupType", "ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup.#.ArmGroupType"},
	{"arm_group_description", "ArmGroupDescription", "ProtocolSection.ArmsInterventionsModule.ArmGroupList.ArmGroup.#.ArmGroupDescription"},
	{"location_name", "LocationFacility", "ProtocolSection.ContactsLocationsModule.LocationList.Location.#.LocationFacility"},
	{"location_city", "LocationCity", "ProtocolSection.ContactsLocationsModule.LocationList.Location.#.LocationCity"},
	{"location_country", "LocationCountry", "ProtocolSection.ContactsLocationsModule.LocationList.Location.#.LocationCountry"},
	{"patient_data_sharing_ipd", "IPDSharing", "ProtocolSection.IPDSharingStatementModule.IPDSharing"},
	{"sponsors_agency", "LeadSponsorName", "ProtocolSection.SponsorCollaboratorsModule.LeadSponsor.LeadSponsorName"},
	{"sponsors_agency_class", "LeadSponsorClass", "ProtocolSection.SponsorCollaboratorsModule.LeadSponsor.LeadSponsorClass"},
	{"publications_reference", "ReferenceCitation", "ProtocolSection.ReferencesModule.ReferenceList.Reference.#.ReferenceCitation"},
	{"publications_PMID", "ReferencePMID", "ProtocolSection.ReferencesModule.ReferenceList.Reference.#.ReferencePMID"},
	{"", "OverallOfficialName", "ProtocolSection.ContactsLocationsModule.OverallOfficialList.OverallOfficial.#.OverallOfficialName"},
	{"", "OverallOfficialAffiliation", "ProtocolSection.ContactsLocationsModule.OverallOfficialList.OverallOfficial.#.OverallOfficialAffiliation"},
	{"", "OverallOfficialRole", "ProtocolSection.ContactsLocationsModule.OverallOfficialList.OverallOfficial.#.OverallOfficialRole"},
	{"", "CentralContactName", "ProtocolSection.ContactsLocationsModule.CentralContactList.CentralContact.#.CentralContactName"},
	{"", "CentralContactRole", "ProtocolSection.ContactsLocationsModule.CentralContactList.CentralContact.#.CentralContactPhone"},
	{"", "CentralContactPhone", "ProtocolSection.ContactsLocationsModule.CentralContactList.CentralContact.#.CentralContactRole"},
	{"", "CentralContactPhoneExt", "ProtocolSection.ContactsLocationsModule.CentralContactList.CentralContact.#.CentralContactPhoneExt"},
	{"", "CentralContactEMail", "ProtocolSection.ContactsLocationsModule.CentralContactList.CentralContact.#.CentralContactEMail"},
	{"", "PointOfContactTitle", "ResultsSection.MoreInfoModule.PointOfContact.PointOfContactTitle"},
	{"", "PointOfContactOrganization", "ResultsSection.MoreInfoModule.PointOfContact.PointOfContactOrganization"},
	{"", "PointOfContactPhone", "ResultsSection.MoreInfoModule.PointOfContact.PointOfContactPhone"},
	{"", "PointOfContactPhoneExt", "ResultsSection.MoreInfoModule.PointOfContact.PointOfContactPhoneExt"},
	{"", "PointOfContactEMail", "ResultsSection.MoreInfoModule.PointOfContact.PointOfContactEMail"},
}
