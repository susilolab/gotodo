package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/susilolab/gotodo/models"
)

type Response struct {
	Error   int         `json:"error"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func TodoServices(app *gin.Engine) {
	todo := app.Group("/todo")

	// `/todo`
	todo.GET("", todoIndex)

	// `/todo/1`
	todo.GET("/:id", func(c *gin.Context) {
		c.String(200, "Todo"+c.Param("id"))
	})

	// `/todo`
	todo.POST("", createTodo)
}

func todoIndex(c *gin.Context) {
	response := Response{
		Error:   0,
		Message: "Data sukses diterima",
		Result:  nil,
	}

	rows, err := models.DB.Query("SELECT * FROM todos")
	if err != nil {
		response.Error = 1
		response.Message = err.Error()

		c.JSON(400, response)
		return
	}
	defer rows.Close()

	var result []models.Todo
	count := 0
	for rows.Next() {
		var todo models.Todo

		err := rows.Scan(&todo.ID, &todo.Body, &todo.Done)
		if err != nil {
			continue
		}

		count += 1
		result = append(result, todo)
	}

	if count == 0 {
		response.Error = 1
		response.Message = "Tabel masih kosong!"

		c.JSON(200, response)
		return
	}

	response.Result = result
	c.JSON(200, response)
}

func createTodo(c *gin.Context) {
	var model models.TodoForm
	resp := Response{
		Error:   0,
		Message: "Data sukses diterima",
		Result:  nil,
	}

	if err := c.ShouldBindJSON(&model); err != nil {
		resp.Error = 1
		resp.Message = "Gagal parsing data json"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
}
