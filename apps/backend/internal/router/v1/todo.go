package v1

import (
	"github.com/SR-SHREYAS/Go-Custom-Boilerplate/internal/handler"
	"github.com/SR-SHREYAS/Go-Custom-Boilerplate/internal/middleware"
	"github.com/labstack/echo/v4"
)

func registerTodoRoutes(r *echo.Group, h *handler.TodoHandler, ch *handler.CommentHandler, auth *middleware.AuthMiddleware) {
	// Todo operations
	todos := r.Group("/todos")
	todos.Use(auth.RequireAuth)

	// Collection operations
	todos.POST("", h.CreateTodo)
	todos.GET("", h.GetTodos)
	todos.GET("/stats", h.GetTodoStats)

	// Individual todo operations
	dynamicTodo := todos.Group("/:id")
	dynamicTodo.GET("", h.GetTodoByID)
	dynamicTodo.PUT("", h.UpdateTodo)
	dynamicTodo.DELETE("", h.DeleteTodo)

	// Todo comments
	todoComments := dynamicTodo.Group("/comments")
	todoComments.POST("", ch.AddComment)
	todoComments.GET("", ch.GetCommentsByTodoID)
}
