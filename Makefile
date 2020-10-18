build: task
	go version
task:
	sql-migrate new 1_users
	sql-migrate new 2_m_categories
	sql-migrate new 3_input_achievements
	sql-migrate new 4_input_achievement_tags
	sql-migrate new 5_todos
	sql-migrate new 6_todo_details
	sql-migrate new 7_output_achievements
	sql-migrate new 8_output_achievement_tags
