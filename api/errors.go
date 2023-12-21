package api

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func Errorhandler(c *fiber.Ctx, err error) error {
	fmt.Println(reflect.TypeOf(err))
	if apiError, ok := err.(Error); ok {
		return c.Status(apiError.Code).JSON(apiError)
	}
	apiError := NewError(http.StatusInternalServerError, err.Error())
	return c.Status(apiError.Code).JSON(apiError)
}

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

// Error implements the error interface
func (e Error) Error() string {
	return e.Err
}

func NewError(code int, err string) Error {
	return Error{
		Code: code,
		Err:  err,
	}
}

func ErrUnAuthorized() Error {
	return Error{
		Code: http.StatusUnauthorized,
		Err:  "Unauthorized request ",
	}
}

func ErrResourceNotFound(res string) Error {
	return Error{
		Code: http.StatusNotFound,
		Err:  res + " resource not found",
	}
}

func ErrBadRequest() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "Invalid JSON request",
	}
}

func ErrInValidID() Error {
	return Error{
		Code: http.StatusBadRequest,
		Err:  "Invalid ID given",
	}
}
