package xvalidator

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/ubaidillahhf/dating-service/app/infra/utility/helper"
)

type Lv1Error struct {
	Param   string
	Message string
}

func TranslateError(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error() // default error
}

func GenerateHumanizeError(payload any, err error) []Lv1Error {
	out := make([]Lv1Error, 0)

	var ve validator.ValidationErrors
	if errors.As(err, &ve) {

		for _, fe := range ve {

			var reflectValue = reflect.TypeOf(payload)

			/**
			* handling source data payload with type interface/any.
			* if payload type using pointer use Elem() instead, no pointer will be error panic when
			* using Elem().
			* note: pointer using new, not pointer using var
			* Why about unuse Elem() for all? for pointer will throw error: reflect: FieldByName of non-struct type *domain
			 */
			if reflectValue.Kind() == reflect.Ptr {
				reflectValue = reflectValue.Elem()
			}

			fieldName := fe.Field()
			fieldNameNs := fe.Namespace()

			field, _ := reflectValue.FieldByName(fieldName)
			fieldJSONName, ok := field.Tag.Lookup("json")
			if !ok {
				fieldJSONName = helper.ConvLastStructNameToCamelCase(fieldNameNs)
			}

			out = append(out, Lv1Error{
				Param:   fieldJSONName,
				Message: TranslateError(fe),
			})
		}
	}

	return out
}
