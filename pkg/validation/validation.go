package validation

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Validator interface {
	ValidateBody(c *fiber.Ctx, input interface{}) []map[string]string
}

type validatorImpl struct {
	validate *validator.Validate
	logger   *zap.SugaredLogger
}

func NewValidator(logger *zap.SugaredLogger) Validator { // ✅ Mengembalikan interface Validator
	return &validatorImpl{
		validate: validator.New(),
		logger:   logger,
	}
}

func (v *validatorImpl) ValidateBody(c *fiber.Ctx, input interface{}) []map[string]string {

	// Parsing request body
	if err := c.BodyParser(input); err != nil {
		return []map[string]string{
			{
				"field":   "body",
				"message": err.Error(),
			},
		}
	}

	err := v.validate.Struct(input)
	if err != nil {
		validationErrors := []map[string]string{}
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, map[string]string{
				"field":   err.Field(),
				"message": getErrorMessage(err.Tag()),
			})
		}
		return validationErrors
	}

	return nil
}

func getErrorMessage(tag string) string {
	messages := map[string]string{
		"required": "This field is required",
		"email":    "Must be a valid email",
		"min":      "Value is too short",
		"max":      "Value is too long",
	}

	if msg, ok := messages[tag]; ok {
		return msg
	}
	return "Invalid value"
}

var Module = fx.Module("validation", fx.Provide(NewValidator)) // ✅ Sesuaikan dengan perubahan
