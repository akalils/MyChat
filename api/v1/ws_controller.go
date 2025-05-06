package v1

import (
	"LishengChat-go/internal/dto/request"
	"LishengChat-go/internal/service/chat"
	"LishengChat-go/pkg/constants"
	"LishengChat-go/pkg/zlog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// WsLogin wss登录 Get
func WsLogin(c *gin.Context) {
	// GET请求获取客户端id，如果客户端id为空应该是业务问题，可能用户输入有误或没有输入，返回400
	clientId := c.Query("client_id")
	if clientId == "" {
		zlog.Error("clientId获取失败")
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": "clientId获取失败",
		})
		return
	}
	chat.NewClientInit(c, clientId)
}

// WsLogout wss登出
func WsLogout(c *gin.Context) {
	var req request.WsLogoutRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := chat.ClientLogout(req.OwnerId)
	JsonBack(c, message, ret, nil)
}
