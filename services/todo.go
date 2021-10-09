package services

import (
	"net/http"
	"strconv"

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
	todo.GET("/:id", viewTodo)

	// `/todo`
	todo.POST("", createTodo)

	// `/todo/:id`
	todo.PATCH("/:id", updateTodo)
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

func viewTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	response := Response{
		Error:   0,
		Message: "Data sukses diterima",
		Result:  nil,
	}
	if err != nil {
		response.Error = 1
		response.Message = "Id invalid"
		c.JSON(400, response)
		return
	}

	rows, err := models.DB.Query("SELECT * FROM todos WHERE id = ?", id)
	if err != nil {
		response.Error = 1
		response.Message = err.Error()

		c.JSON(400, response)
		return
	}
	defer rows.Close()

	var result []models.Todo
	for rows.Next() {
		var todo models.Todo
		err := rows.Scan(&todo.ID, &todo.Body, &todo.Done)
		if err != nil {
			continue
		}

		result = append(result, todo)
	}
	if len(result) == 0 {
		response.Error = 1
		response.Message = "Data tidak ditemukan"
		c.JSON(400, response)
		return
	}

	response.Result = result
	c.JSON(200, response)
}

func createTodo(c *gin.Context) {
	var model models.TodoForm
	resp := Response{
		Error:   0,
		Message: "Data sukses disimpan",
		Result:  nil,
	}

	if err := c.ShouldBindJSON(&model); err != nil {
		resp.Error = 1
		resp.Message = "Gagal parsing data json"
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	_, err := models.DB.Exec("INSERT INTO todos(body, done) VALUES(?, ?)", model.Body, model.Done)
	if err != nil {
		resp.Error = 1
		resp.Message = "Data gagal disimpan"
		c.JSON(400, resp)
		return
	}

	c.JSON(200, resp)
}

func updateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	var model models.TodoForm
	resp := Response{
		Error:   0,
		Message: "Data sukses diupdate",
		Result:  nil,
	}

	if err != nil {
		resp.Error = 1
		resp.Message = "Invalid id"
		return
	}

	if err := c.ShouldBindJSON(&model); err != nil {
		resp.Error = 1
		resp.Message = "Gagal parsing data json " + err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	_, err = models.DB.Exec("UPDATE todos SET body = ?, done = ? WHERE id = ?", model.Body, model.Done, id)
	if err != nil {
		resp.Error = 1
		resp.Message = err.Error()
		c.JSON(400, resp)
		return
	}

	c.JSON(200, resp)
}
