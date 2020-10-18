
-- +migrate Up
CREATE TABLE `input_achievements` (
  `input_achievement_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `reference_url` varchar(300) DEFAULT NULL,
  `summary` varchar(1000) DEFAULT NULL,
  `input_time` varchar(30) DEFAULT NULL,
  `created_by` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `modified_by` int DEFAULT NULL,
  `modified_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`input_achievement_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `input_achievements_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE `input_achievements`;