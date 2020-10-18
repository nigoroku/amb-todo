package service

import (
	"bytes"
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/boil"
	"local.packages/models"
)

var (
	todoDBTypes = map[string]string{`TodoID`: `int`, `CreatedBy`: `int`, `CreatedAt`: `timestamp`, `ModifiedBy`: `int`, `ModifiedAt`: `timestamp`, `UserID`: `int`}
	_           = bytes.MinRead
)

func Init() {
	// DB接続
	db, err := sql.Open("mysql", "moizumi:jamyuki0210@tcp(localhost:3306)/ambitious_test?parseTime=true")
	if err != nil {
		log.Fatalf("Cannot connect database: %v", err)
	}
	boil.SetDB(db)
}

func MustTx(transactor boil.ContextTransactor, err error) boil.ContextTransactor {
	if err != nil {
		panic(fmt.Sprintf("Cannot create a transactor: %s", err))
	}
	return transactor
}

func TestAddTodos(t *testing.T) {
	// DB接続
	Init()

	t.Parallel()
	d1 := &models.TodoDetail{Content: "content1", Checked: false}
	d2 := &models.TodoDetail{Content: "content2", Checked: true}
	ds := models.TodoDetailSlice{d1, d2}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	// TODO: 事前に登録
	user := models.User{UserID: 153}

	todoService := &TodoService{ctx, tx}
	_, err2 := todoService.AddTodos(ds, user.UserID)

	if err2 != nil {
		t.Error(err2)
	}

	got2, err3 := todoService.FindTodosByUser(user.UserID)

	if err3 != nil {
		t.Error(err3)
	}

	assert.Equal(t, len(got2), 2)
	assert.Equal(t, got2[0].Content, "content1")
	assert.Equal(t, got2[1].Content, "content2")
}
