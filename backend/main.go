package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"

	_ "github.com/gmemcc/oopstack/docs"          // Import generated docs
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title Echo OpenAPI Example
// @version 1.0
// @description This is a sample server.
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:1323
// @BasePath /

// @swagger 3.0

// NameRequest represents the expected request body.
// swagger:model NameRequest
type NameRequest struct {
	// Name of the person to greet.
	// Required: true
	Name string `json:"name" example:"Alex"`
}

// GreetingResponse represents the successful response.
// swagger:model GreetingResponse
type GreetingResponse struct {
	// The personalized greeting message.
	// Example: Hello Alex
	Greeting string `json:"greeting" example:"Hello Alex"`
}

// ErrorResponse represents an error response.
// swagger:model ErrorResponse
type ErrorResponse struct {
	// Error message
	Message string `json:"message" example:"Invalid request"`
}

func main() {
	e := echo.New()

	// Routes
	e.POST("/api/hello", helloHandler)
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Start server
	e.Logger.Fatal(e.Start(":8085"))
}

// helloHandler godoc
// @Summary      Greet the user
// @Description  Responds with a personalized greeting message
// @Tags         greeting
// @Accept       json
// @Produce      json
// @Param        request  body      NameRequest       true  "Name in JSON"
// @Success      200      {object}  GreetingResponse  "Successful response"
// @Failure      400      {object}  ErrorResponse     "Bad request"
// @Router       /api/hello [post]
func helloHandler(c echo.Context) error {
	req := new(NameRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Invalid request"})
	}
	if req.Name == "" {
		return c.JSON(http.StatusBadRequest, ErrorResponse{Message: "Name is required"})
	}
	greeting := fmt.Sprintf("Hello %s", req.Name)
	return c.JSON(http.StatusOK, GreetingResponse{Greeting: greeting})
}
