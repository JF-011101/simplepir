package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendResponse(c *gin.Context, response interface{}) {
	c.JSON(http.StatusOK, response)
}

// 查询用户名是否存在
type PirQueryParam struct {
	PhoneNumber uint64 `json:"phonenumber"`
}
