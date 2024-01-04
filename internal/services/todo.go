package services

import db "todoapp/internal/models"

func GetTodo(id uint) (*db.Todo, error) {
	todo := &db.Todo{ID: id}
	if err := db.DB.Find(todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

// TODO: add pagination
func GetAllTodos() []db.Todo {
	var todos []db.Todo
	db.DB.Order("id ASC").Find(&todos)

	return todos
}

func CreateTodo(content string) (*db.Todo, error) {
	todo := &db.Todo{Content: content, IsComplete: false}
	if err := db.DB.Create(&todo).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func UpdateTodoContent(id uint, content string) (*db.Todo, error) {
	todo := &db.Todo{ID: id}

	if err := db.DB.Model(todo).Updates(db.Todo{Content: content}).Error; err != nil {
		return nil, err
	}

	return todo, nil
}

func MarkTodoAsCompleted(id uint) (*db.Todo, error) {
	return changeTodoCompleteness(id, true)
}

func MarkTodoAsInCompleted(id uint) (*db.Todo, error) {
	return changeTodoCompleteness(id, false)
}

// private
func changeTodoCompleteness(id uint , isComplete bool) (*db.Todo, error) {
todo := &db.Todo{ID: id}

	if err := db.DB.Model(todo).Updates(db.Todo{IsComplete: isComplete}).Error; err != nil {
		return nil, err
	}

	return todo, nil
}