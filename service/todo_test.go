package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTodos(t *testing.T) {
	array := []string{"todo1", "todo2", "todo3"}

	todoService := &TodoService{}

	err := todoService.AddTodos(array[:])

	assert.Nil(t, err)
}
