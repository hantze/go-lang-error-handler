package middleware

import (
	"fmt"
	"gosqlx/internal/gosqlx/helper"

	"gopkg.in/go-playground/validator.v8"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"reflect"
)

func respondWithError(code int, ve interface{}, c *gin.Context) {
	c.JSON(code, &ve)
	c.Abort()
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			for _, e := range c.Errors {
				fmt.Println(reflect.TypeOf(e.Err), e.Type)
				switch e.Err.(type) {
					case validator.ValidationErrors:
						log.Println("Bind Error")
						errs := e.Err.(validator.ValidationErrors)
						ve := helper.NewValidationError()
						if errs != nil  {
							for _, value := range errs {
								ve.GenerateError(value.Field, value.ActualTag, fmt.Sprintf("%v", value.Value))
							}
						}
						respondWithError(http.StatusBadRequest, ve, c)
						break

					default:
						respondWithError(http.StatusInternalServerError, "Server Error", c)
						break
				}
			}
		}
	}
}