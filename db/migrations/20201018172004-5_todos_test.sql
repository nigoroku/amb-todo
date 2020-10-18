
-- +migrate Up
CREATE TABLE `todos` (
  `todo_id` int NOT NULL AUTO_INCREMENT,
  `created_by` int NOT NULL,
  `created_at` timestamp NOT NULL,
  `modified_by` int DEFAULT NULL,
  `modified_at` timestamp NULL DEFAULT NULL,
  `user_id` int NOT NULL,
  PRIMARY KEY (`todo_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `todos_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=81 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- +migrate Down
DROP TABLE `todos`;
