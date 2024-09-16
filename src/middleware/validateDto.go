package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
)

// Validator instance (singleton)
var validate = validator.New()

// ContextKey to store DTO in request context
type ContextKey string

// ValidateDTO is the middleware that dynamically validates incoming requests against a DTO
func ValidateDTO(dto interface{}) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Create a new instance of the provided DTO
			dtoInstance := reflect.New(reflect.TypeOf(dto).Elem()).Interface()

			// Decode the request body into the DTO
			if err := json.NewDecoder(r.Body).Decode(dtoInstance); err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}

			// Validate the DTO
			if err := validate.Struct(dtoInstance); err != nil {
				http.Error(w, "Validation failed: "+err.Error(), http.StatusBadRequest)
				return
			}

			// Add the validated DTO to the context
			ctx := context.WithValue(r.Context(), ContextKey("dto"), dtoInstance)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// Retrieve the DTO from context
func GetDTOFromContext(r *http.Request) interface{} {
	return r.Context().Value(ContextKey("dto"))
}
