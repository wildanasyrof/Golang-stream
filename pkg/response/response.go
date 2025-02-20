package response

import (
	"github.com/gofiber/fiber/v2"
)

// SuccessResponse struct for successful responses
type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse struct for error responses
type ErrorResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error,omitempty"`
}

// ValidationError struct for validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// NewValidationError membuat instance error validasi baru
func NewValidationError(field, message string) ValidationError {
	return ValidationError{Field: field, Message: message}
}

// Success returns a standard success response
func Success(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Created returns a standard created response
func Created(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// Error returns a standard error response with optional error details
func Error(c *fiber.Ctx, statusCode int, message string, errDetails ...interface{}) error {
	var errData interface{} = nil
	if len(errDetails) > 0 {
		errData = errDetails[0]
	}

	return c.Status(statusCode).JSON(ErrorResponse{
		Status:  "error",
		Message: message,
		Error:   errData,
	})
}
