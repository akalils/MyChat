package v1

import (
	"LishengChat-go/internal/dto/request"
	"LishengChat-go/internal/service/gorm"
	"LishengChat-go/pkg/constants"
	"LishengChat-go/pkg/zlog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCurContactListInChatRoom 获取当前聊天室联系人列表
func GetCurContactListInChatRoom(c *gin.Context) {
	var req request.GetCurContactListInChatRoomRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, rspList, ret := gorm.ChatRoomService.GetCurContactListInChatRoom(req.OwnerId, req.ContactId)
	JsonBack(c, message, ret, rspList)
}
