CREATE TABLE `projects` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `status` varchar(20) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'draft',
  `owner` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `created_on` timestamp NULL DEFAULT NULL,
  `modified_by` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `modified_on` timestamp NULL DEFAULT NULL,
  `name` mediumtext COLLATE utf8mb4_unicode_ci,
  `identifier` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `description` mediumtext COLLATE utf8mb4_unicode_ci,
  PRIMARY KEY (`id`),
  KEY `projects_modified_by_foreign` (`modified_by`),
  CONSTRAINT `projects_modified_by_foreign` FOREIGN KEY (`modified_by`) REFERENCES `directus_users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `cove_basic_projects` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `cove_basic_id` int(10) unsigned DEFAULT NULL,
  `projects_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `cove_basic_projects_cove_basic_id_foreign` (`cove_basic_id`),
  KEY `cove_basic_projects_projects_id_foreign` (`projects_id`),
  CONSTRAINT `cove_basic_projects_cove_basic_id_foreign` FOREIGN KEY (`cove_basic_id`) REFERENCES `cove_basic` (`id`) ON DELETE SET NULL,
  CONSTRAINT `cove_basic_projects_projects_id_foreign` FOREIGN KEY (`projects_id`) REFERENCES `projects` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;