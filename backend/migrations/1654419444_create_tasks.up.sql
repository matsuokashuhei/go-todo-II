CREATE TABLE `tasks` (`id` bigint NOT NULL AUTO_INCREMENT, `create_time` timestamp NULL, `update_time` timestamp NULL, `title` varchar(255) NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;