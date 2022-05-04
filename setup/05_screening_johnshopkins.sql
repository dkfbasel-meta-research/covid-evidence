CREATE TABLE `johnshopkins` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` timestamp NULL DEFAULT NULL,
  `modified_by` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `modified_on` timestamp NULL DEFAULT NULL,
  `date` date DEFAULT NULL,
  `country` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `cases` int(11) DEFAULT NULL,
  `week_number` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `johnshopkins_modified_by_foreign` (`modified_by`),
  CONSTRAINT `johnshopkins_modified_by_foreign` FOREIGN KEY (`modified_by`) REFERENCES `directus_users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;