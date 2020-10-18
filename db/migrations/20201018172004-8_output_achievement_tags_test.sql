
-- +migrate Up
CREATE TABLE `output_achievement_tags` (
  `output_achievement_tag_id` int NOT NULL AUTO_INCREMENT,
  `category_id` int NOT NULL,
  `created_by` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `modified_by` int DEFAULT NULL,
  `modified_at` timestamp NULL DEFAULT NULL,
  `output_achievement_id` int NOT NULL,
  PRIMARY KEY (`output_achievement_tag_id`),
  KEY `output_achievement_id` (`output_achievement_id`),
  KEY `category_id` (`category_id`),
  CONSTRAINT `output_achievement_tags_ibfk_1` FOREIGN KEY (`output_achievement_id`) REFERENCES `output_achievements` (`output_achievement_id`),
  CONSTRAINT `output_achievement_tags_ibfk_2` FOREIGN KEY (`category_id`) REFERENCES `m_categories` (`category_id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE `output_achievement_tags`;