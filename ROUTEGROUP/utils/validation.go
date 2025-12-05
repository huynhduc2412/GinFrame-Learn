package utils

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ValidationRequired(fieldName , value string) error {
	if value == "" {
		return fmt.Errorf("%s is empty" , fieldName)
	}
	return nil
}

func ValidationStringLength(fieldName , value string , min , max int) error {
	l := len(value)
	if l < min || l > max {
		return fmt.Errorf("%s must be between %d and %d characters" , fieldName , min , max)
	}
	return nil
}

func ValidationRegex(fieldName , value string , re *regexp.Regexp , errorMessage string) error {
	if !re.MatchString(value) {
		return fmt.Errorf("%s %s" , fieldName , errorMessage)
	}
	return nil
}

func ValidationPositiveInt(fielName , value string) (int , error) {
	v , err := strconv.Atoi(value)
	if err != nil {
		return 0 , fmt.Errorf("%s must be a number" , fielName)
	}
	if v <= 0 {
		return 0 , fmt.Errorf("%s must be a positive number" , fielName)
	}
	return v , nil
}

func ValidationUuid(fieldName , value string) (uuid.UUID , error) {
	uid , err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil , fmt.Errorf("%s must be a valid UUID" , fieldName)
	}
	return uid , nil
}
func ValidationInList(fieldName , value string , allowedList map[string]struct{}) error {
	if _ , ok := allowedList[value] ; !ok {
		return	fmt.Errorf("%s must be one of : %v" , fieldName , keys(allowedList))
	}
	return nil
}
func keys(m map[string]struct{}) []string {
	var k []string
	for key := range m {
		k = append(k , key)
	}
	return k
}
func HandleValidationErrors(err error) gin.H {
	if validationError , ok := err.(validator.ValidationErrors) ; ok {
		errors := make(map[string]string)
		for _ , e := range validationError {
				errors[e.Field()] = e.Field() + e.Tag()
		}
		return gin.H{
			"error" : errors,
		}
	}
	return gin.H{
		"error" : err.Error(),
	}
}