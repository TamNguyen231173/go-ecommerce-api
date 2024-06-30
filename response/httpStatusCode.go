package response

import "github.com/gin-gonic/gin"

const (
	ErrorCodeStatusOK = 0
	ErrorCodeStatusBadRequest = 400
	ErrorCodeStatusUnauthorized = 401
	ErrorCodeStatusForbidden = 403
	ErrorCodeStatusNotFound = 404
	ErrorCodeStatusInternalServerError = 500
	ErrorCodeInvalidParam = 1001
)

var Msg = map[int]string{
	ErrorCodeStatusOK: "OK",
	ErrorCodeStatusBadRequest: "Bad Request",
	ErrorCodeStatusUnauthorized: "Unauthorized",
	ErrorCodeStatusForbidden: "Forbidden",
	ErrorCodeStatusNotFound: "Not Found",
	ErrorCodeStatusInternalServerError: "Internal Server Error",
	ErrorCodeInvalidParam: "Invalid Param",
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code": code,
		"message": message,
	})
}