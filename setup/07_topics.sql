CREATE TABLE `topic` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `status` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'draft',
  `user_created` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `date_created` timestamp NULL DEFAULT NULL,
  `user_updated` char(36) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `date_updated` timestamp NULL DEFAULT NULL,
  `name` mediumtext COLLATE utf8mb4_unicode_ci,
  `order` int(11) NOT NULL,
  `description` mediumtext COLLATE utf8mb4_unicode_ci,
  `keywords` mediumtext COLLATE utf8mb4_unicode_ci,
  `keyword_search_columns` varchar(1024) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `topic_user_created_foreign` (`user_created`),
  KEY `topic_user_updated_foreign` (`user_updated`),
  CONSTRAINT `topic_user_created_foreign` FOREIGN KEY (`user_created`) REFERENCES `directus_users` (`id`) ON DELETE SET NULL,
  CONSTRAINT `topic_user_updated_foreign` FOREIGN KEY (`user_updated`) REFERENCES `directus_users` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `topic_cove_basic` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `topic_id` int(10) unsigned DEFAULT NULL,
  `cove_basic_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `topic_cove_basic_topic_id_foreign` (`topic_id`),
  KEY `topic_cove_basic_cove_basic_id_foreign` (`cove_basic_id`),
  CONSTRAINT `topic_cove_basic_cove_basic_id_foreign` FOREIGN KEY (`cove_basic_id`) REFERENCES `cove_basic` (`id`) ON DELETE SET NULL,
  CONSTRAINT `topic_cove_basic_topic_id_foreign` FOREIGN KEY (`topic_id`) REFERENCES `topic` (`id`) ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;