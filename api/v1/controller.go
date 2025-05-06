package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonBack(c *gin.Context, message string, ret int, data interface{}) {
	if ret == 0 { // 业务正常走完流程
		if data != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": message,
				"data":    data,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": message,
			})
		}
	} else if ret == -2 { // 业务数据问题导致未正常走完业务流程，返回400
		c.JSON(http.StatusOK, gin.H{
			"code":    400,
			"message": message,
		})
	} else if ret == -1 { // 系统错误，比如序列化失败，radis缓存失败等，返回500
		c.JSON(http.StatusOK, gin.H{
			"code":    500,
			"message": message,
		})
	}
}
