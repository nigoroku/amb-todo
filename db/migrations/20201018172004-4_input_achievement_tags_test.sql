
-- +migrate Up
CREATE TABLE `input_achievement_tags` (
  `input_achievement_tag_id` int NOT NULL AUTO_INCREMENT,
  `input_achievement_id` int NOT NULL,
  `category_id` int NOT NULL,
  `created_by` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `modified_by` int DEFAULT NULL,
  `modified_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`input_achievement_tag_id`),
  KEY `input_achievement_id` (`input_achievement_id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `input_achievement_tags_ibfk_1` FOREIGN KEY (`input_achievement_id`) REFERENCES `input_achievements` (`input_achievement_id`),
  CONSTRAINT `input_achievement_tags_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `m_categories` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=59 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE `input_achievement_tags`;