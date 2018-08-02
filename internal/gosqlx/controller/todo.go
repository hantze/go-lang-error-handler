package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gosqlx/internal/gosqlx/model"
	"github.com/jmoiron/sqlx"
	"errors"
	)

// TodoController represents the todo controller
type TodoController struct {
	db *sqlx.DB
}

// List lists all todo
func (tc *TodoController) List(c *gin.Context) {

	todoModel := model.NewTodoModel(tc.db)
	listTodos := todoModel.AllTodos()


	if len(listTodos) > 0 {
		c.JSON(http.StatusOK, listTodos)
	}

	{
		c.Error(errors.New("Cannot find any data!"))
	}
}

// Create creates a new todo
func (tc *TodoController) Create(c *gin.Context) {

	todoModel := model.NewTodoModel(tc.db)

	err := c.ShouldBindJSON(todoModel)

	if err != nil {
		c.Error(err)
		return
	}

	result := todoModel.CreateTodo(todoModel)

	if result == true {
		c.JSON(http.StatusOK, &model.TodoModel{
			ID:        todoModel.ID,
			UserID:    todoModel.UserID,
			Title:     todoModel.Title,
			Completed: todoModel.Completed,
		})
	}
}

// Get gets a todo
func (tc *TodoController) Get(c *gin.Context) {

	todoModel := model.NewTodoModel(tc.db)
	result := todoModel.GetTodo(c.Params.ByName("id"))

	if result.ID != "" {
		c.JSON(http.StatusOK, result)
	}

	{
		c.JSON(http.StatusBadRequest, "Cannot Get the Data")
	}
}

// Update updates a todo
func (tc *TodoController) Update(c *gin.Context) {

	todoModel := model.NewTodoModel(tc.db)
	c.BindJSON(&todoModel)
	result := todoModel.UpdateTodo(c.Params.ByName("id"), todoModel)

	if result == true {
		c.JSON(http.StatusOK, "Success Update the Data")
	}

	{
		c.JSON(http.StatusBadRequest, "Cannot Update the Data")
	}
}

// Delete deletes a todo
func (tc *TodoController) Delete(c *gin.Context) {

	todoModel := model.NewTodoModel(tc.db)
	result := todoModel.DeleteTodo(c.Params.ByName("id"))

	if result == true {
		c.JSON(http.StatusOK, "Success Delete the Data")
	}

	{
		c.JSON(http.StatusBadRequest, "Cannot Delete the Data")
	}
}

// NewTodoController creates a new todo controller
func NewTodoController(db *sqlx.DB) *TodoController {

	return &TodoController{db: db}
}
