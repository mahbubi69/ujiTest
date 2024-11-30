package helper

// import (
// 	"encoding/json"
// 	"errors"
// 	"strings"
// 	"ujiTest/models"

// 	"github.com/go-playground/validator"
// )

// type InputValidation struct {
// 	Validator *validator.Validate
// }

// func UniqueNonEmptyElementsOf(s []string) []string {
// 	unique := make(map[string]bool, len(s))
// 	us := make([]string, len(unique))
// 	for _, elem := range s {
// 		if len(elem) != 0 {
// 			if !unique[elem] {
// 				us = append(us, elem)
// 				unique[elem] = true
// 			}
// 		}
// 	}

// 	return us
// }

// // config/validation.go
// func (iv *InputValidation) Validate(data interface{}) error {
// 	var errFields []models.ErrorInputResponse

// 	err := iv.Validator.Struct(data)
// 	if err != nil {
// 		for _, err := range err.(validator.ValidationErrors) {
// 			var errField models.ErrorInputResponse
// 			switch err.Tag() {
// 			case "code":
// 				errField.FieldName = strings.ToLower(err.Field())
// 				errField.Message = "Email format is invalid"
// 			case "name":
// 				errField.FieldName = strings.ToLower(err.Field())
// 				errField.Message = err.Field() + " must be minimum " + err.Param() + " characters"
// 			case "model":
// 				errField.FieldName = strings.ToLower(err.Field())
// 				errField.Message = err.Field() + " maximum allowed is" + err.Param() + " characters"
// 			case "tech":
// 				errField.FieldName = strings.ToLower(err.Field())
// 				errField.Message = err.Field() + " cannot be blank"
// 			case "status":
// 				errField.FieldName = strings.ToLower(err.Field())
// 				errField.Message = err.Field() + " cannot be blank"
// 			}

// 			errFields = append(errFields, errField)
// 		}
// 	}
// 	// jika tidak ada error, kembalikan menjadi nil
// 	if len(errFields) == 0 {
// 		return nil
// 	}
// 	// sebaliknya, ubah errFields menjadi JSON Array of objects,
// 	// lalu diperlakukan sebagai error
// 	marshaledErr, _ := json.Marshal(errFields)
// 	return errors.New(string(marshaledErr))
// }
