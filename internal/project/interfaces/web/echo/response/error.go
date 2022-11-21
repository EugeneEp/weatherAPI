package response

import "github.com/labstack/echo/v4"

func ErrServerInternal(internal error) *echo.HTTPError {
	return &echo.HTTPError{Code: 500, Message: "Internal Server Error", Internal: internal}
}

func ErrBadRequest(internal error) *echo.HTTPError {
	return &echo.HTTPError{Code: 400, Message: "Bad Request", Internal: internal}
}

func ErrNotFound(internal error) *echo.HTTPError {
	return &echo.HTTPError{Code: 404, Message: "Not found", Internal: internal}
}
