
-- +migrate Up
CREATE TABLE `todo_details` (
  `todo_detail_id` int NOT NULL AUTO_INCREMENT,
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `created_by` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `checked` tinyint(1) NOT NULL,
  `todo_id` int DEFAULT NULL,
  `modified_by` int DEFAULT NULL,
  `modified_at` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`todo_detail_id`),
  KEY `todo_id` (`todo_id`),
  CONSTRAINT `todo_details_ibfk_1` FOREIGN KEY (`todo_id`) REFERENCES `todos` (`todo_id`)
) ENGINE=InnoDB AUTO_INCREMENT=89 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE `todo_details`;
