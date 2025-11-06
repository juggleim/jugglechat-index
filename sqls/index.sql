CREATE TABLE `appnavs` (
  `id` int NOT NULL AUTO_INCREMENT,
  `app_key` varchar(50) DEFAULT NULL,
  `alias_no` varchar(50) DEFAULT NULL,
  `admin_url` varchar(500) DEFAULT NULL,
  `api_url` varchar(500) DEFAULT NULL,
  `ws_url` varchar(500) DEFAULT NULL,
  `app_url` varchar(500) DEFAULT NULL,
  `created_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3),
  `updated_time` datetime(3) DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3),
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_appkey` (`app_key`),
  UNIQUE KEY `uniq_alias` (`alias_no`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;