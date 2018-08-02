package helper

import (
	"fmt"
)

type ValidationError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Fields   []FieldStruct `json:"fields"`
}

type FieldStruct struct{
	Field string `json:"field"`
	Message string `json:"message"`
}

func (ve *ValidationError) GenerateError(key string, messages ...string){
	switch messages[0] {
		case "min":
			temp := &FieldStruct{
				Field: key,
				Message: fmt.Sprintf("%s, minimum %s character(s)", key, messages[1]),
			}
			(*ve).Fields = append((*ve).Fields, *temp)
		case "required":
			temp := &FieldStruct{
				Field: key,
				Message: fmt.Sprintf("%s is required", key),
			}
			(*ve).Fields = append((*ve).Fields, *temp)
	}
}

func NewValidationError() *ValidationError{
	ve := ValidationError{
		Code: "ValidationError",
		Message: "validation failed",
	}
	return &ve
}