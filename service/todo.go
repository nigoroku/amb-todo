package service

import (
	"context"
	"fmt"
	"time"

	// "github.com/kzpolicy/amb-todo/models"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"local.packages/models"

	"log"
)

type TodoService struct {
	ctx context.Context
	db  boil.ContextExecutor
}

func NewTodoService() *TodoService {
	ctx := context.Background()
	// DB作成
	db := boil.GetContextDB()

	return &TodoService{ctx, db}
}

func (t *TodoService) FindTodosByUser(userId int) (models.TodoDetailSlice, error) {

	// データ取得
	var todoDetails models.TodoDetailSlice
	err := models.NewQuery(
		qm.Select("td.*"),
		qm.From("todos as t"),
		qm.InnerJoin("todo_details as td ON t.todo_id = td.todo_id"),
		qm.Where("t.user_id=?", userId),
		// qm.Where("t.user_id=? and t.created_at > DATE_SUB(NOW(), INTERVAL 1 DAY)", userId),
	).Bind(t.ctx, t.db, &todoDetails)

	if err != nil {
		fmt.Printf("error %v", err)
		return models.TodoDetailSlice{}, err
	}

	return todoDetails, err
}

func (ts *TodoService) AddTodos(todos models.TodoDetailSlice, userId int) (models.TodoDetailSlice, error) {
	jst := time.FixedZone("JST", 9*60*60)
	nowJST := time.Now().In(jst)

	// Todo登録
	var t models.Todo
	t.UserID = userId
	t.CreatedBy = userId
	t.CreatedAt = nowJST
	fmt.Println(t.CreatedAt)
	err1 := t.Insert(ts.ctx, ts.db, boil.Infer())

	if err1 != nil {
		log.Fatalf("error %v", err1)
		return nil, err1
	}

	// TodoDetails登録
	var todo_details models.TodoDetailSlice
	for _, todo := range todos {
		if todo.Content == "" {
			continue
		}
		td := &models.TodoDetail{}
		td.TodoID = t.TodoID
		td.Checked = false
		td.Content = todo.Content
		td.CreatedBy = userId
		td.CreatedAt = nowJST
		err2 := td.Insert(ts.ctx, ts.db, boil.Infer())

		todo_details = append(todo_details, td)

		if err2 != nil {
			log.Fatalf("error %v", err2)
			return nil, err2
		}
	}
	return todo_details, nil
}

func (t *TodoService) UpdateTodos(todoDetails models.TodoDetailSlice) error {
	jst := time.FixedZone("JST", 9*60*60)
	nowJST := time.Now().In(jst)

	for _, detail := range todoDetails {

		updCols := map[string]interface{}{
			models.TodoDetailColumns.TodoDetailID: detail.TodoDetailID,
			models.TodoDetailColumns.Checked:      detail.Checked,
			models.TodoDetailColumns.ModifiedAt:   nowJST,
		}

		query := qm.WhereIn(models.TodoDetailColumns.TodoDetailID+" = ?", detail.TodoDetailID)

		_, err := models.TodoDetails(query).UpdateAll(t.ctx, t.db, updCols)

		if err != nil {
			return err
		}
	}

	return nil
}
