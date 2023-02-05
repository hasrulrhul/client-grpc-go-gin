package utils

import (
	"encoding/json"
	"net/http"

	"github.com/hasrulrhul/client-grpc-go-gin/model"

	"github.com/gin-gonic/gin"
)

// HandleSuccess is function wrap in respon success
func HandleSuccess(c *gin.Context, data interface{}) {
	var returnData = model.ResponWrapper{
		Success: true,
		Message: "Success",
		Data:    data,
	}
	c.JSON(http.StatusOK, returnData)
}

// HandleError is function wrap in respon failed
func HandleError(c *gin.Context, status int, message string) {
	var returnData = model.ResponWrapper{
		Success: false,
		Message: message,
	}
	c.JSON(status, returnData)
}

func HandleSuccessReturn(data interface{}) string {
	var returnData = model.ResponWrapper{
		Success: true,
		Message: "Success",
		Data:    data,
	}
	res, _ := json.Marshal(&returnData)
	return string(res)
}
