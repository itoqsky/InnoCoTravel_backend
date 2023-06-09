package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

// type statusOkResponse struct {
// 	Status string `json:"status"`
// }

type dataResponse struct {
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
