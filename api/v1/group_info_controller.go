package v1

import (
	"LishengChat-go/internal/dto/request"
	"LishengChat-go/internal/service/gorm"
	"LishengChat-go/pkg/constants"
	"LishengChat-go/pkg/zlog"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateGroup 创建群聊
func CreateGroup(c *gin.Context) {
	var createGroupReq request.CreateGroupRequest
	if err := c.BindJSON(&createGroupReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.CreateGroup(createGroupReq)
	JsonBack(c, message, ret, nil)
}

// LoadMyGroup 获取我创建的群聊
func LoadMyGroup(c *gin.Context) {
	var loadMyGroupReq request.OwnlistRequest
	if err := c.BindJSON(&loadMyGroupReq); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, groupList, ret := gorm.GroupInfoService.LoadMyGroup(loadMyGroupReq.OwnerId)
	JsonBack(c, message, ret, groupList)
}

// CheckGroupAddMode 检查群聊加群方式
func CheckGroupAddMode(c *gin.Context) {
	var req request.CheckGroupAddModeRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, addMode, ret := gorm.GroupInfoService.CheckGroupAddMode(req.GroupId)
	JsonBack(c, message, ret, addMode)
}

// EnterGroupDirectly 直接进群
func EnterGroupDirectly(c *gin.Context) {
	var req request.EnterGroupDirectlyRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.EnterGroupDirectly(req.OwnerId, req.ContactId)
	JsonBack(c, message, ret, nil)
}

// LeaveGroup 退群
func LeaveGroup(c *gin.Context) {
	var req request.LeaveGroupRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.LeaveGroup(req.UserId, req.GroupId)
	JsonBack(c, message, ret, nil)
}

// DismissGroup 解散群聊
func DismissGroup(c *gin.Context) {
	// 前端根据用户操作生成一个 DismissGroupRequest 请求体，其中包含解散群聊所需的信息，例如 OwnerId 和 GroupId
	var req request.DismissGroupRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{ // 前端将生成的请求体绑定到HTTP请求的上下文中，并通过JSON格式发送给后端
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	// 后端调用 gorm.GroupInfoService.DismissGroup(req.OwnerId, req.GroupId) 方法，在服务层进行具体的业务逻辑处理
	message, ret := gorm.GroupInfoService.DismissGroup(req.OwnerId, req.GroupId)
	// 处理完成后，后端通过 JsonBack 函数将结果以JSON格式返回给前端，前端根据返回的结果进行相应的处理和展示
	JsonBack(c, message, ret, nil)
}

// GetGroupInfo 获取群聊详情
func GetGroupInfo(c *gin.Context) {
	var req request.GetGroupInfoRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, groupInfo, ret := gorm.GroupInfoService.GetGroupInfo(req.GroupId)
	JsonBack(c, message, ret, groupInfo)
}

// UpdateGroupInfo 更新群聊消息
func UpdateGroupInfo(c *gin.Context) {
	var req request.UpdateGroupInfoRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.UpdateGroupInfo(req)
	JsonBack(c, message, ret, nil)
}

// GetGroupMemberList 获取群聊成员列表
func GetGroupMemberList(c *gin.Context) {
	var req request.GetGroupMemberListRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, groupMemberList, ret := gorm.GroupInfoService.GetGroupMemberList(req.GroupId)
	JsonBack(c, message, ret, groupMemberList)
}

// RemoveGroupMembers 移除群聊成员
func RemoveGroupMembers(c *gin.Context) {
	var req request.RemoveGroupMembersRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.RemoveGroupMembers(req)
	JsonBack(c, message, ret, nil)
}

// GetGroupInfoList 获取群聊列表 - 管理员
func GetGroupInfoList(c *gin.Context) {
	message, groupList, ret := gorm.GroupInfoService.GetGroupInfoList()
	JsonBack(c, message, ret, groupList)
}

// DeleteGroups 删除列表中群聊 - 管理员
func DeleteGroups(c *gin.Context) {
	var req request.DeleteGroupsRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.DeleteGroups(req.UuidList)
	JsonBack(c, message, ret, nil)
}

// SetGroupsStatus 设置群聊是否启用
func SetGroupsStatus(c *gin.Context) {
	var req request.SetGroupsStatusRequest
	if err := c.BindJSON(&req); err != nil {
		zlog.Error(err.Error())
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": constants.SYSTEM_ERROR,
		})
		return
	}
	message, ret := gorm.GroupInfoService.SetGroupsStatus(req.UuidList, req.Status)
	JsonBack(c, message, ret, nil)
}