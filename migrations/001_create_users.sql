CREATE TABLE `users` (
                         `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                         `created_at` datetime(3) DEFAULT NULL,
                         `updated_at` datetime(3) DEFAULT NULL,
                         `username` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
                         `password` varchar(128) COLLATE utf8mb4_unicode_ci NOT NULL,
                         `nickname` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `avatar` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `email` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `role` tinyint(3) unsigned NOT NULL DEFAULT '2',
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `idx_users_username` (`username`),
                         UNIQUE KEY `idx_users_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci