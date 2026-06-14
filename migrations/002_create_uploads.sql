CREATE TABLE `uploads` (
                           `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                           `created_at` datetime(3) DEFAULT NULL,
                           `updated_at` datetime(3) DEFAULT NULL,
                           `user_id` bigint(20) unsigned NOT NULL,
                           `filename` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                           `url` varchar(512) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                           `path` varchar(512) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                           `size` bigint(20) DEFAULT NULL,
                           `mime_type` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                           PRIMARY KEY (`id`),
                           KEY `idx_uploads_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci