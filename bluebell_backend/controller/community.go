package controller

import (
	"bluebell_backend/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CommunityHandler(c *gin.Context) {
	communityList, err := service.GetCommunityList()
	if err != nil {
		zap.L().Error("service.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServerBusy) // 不要轻易将服务端错误暴漏
		return
	}
	ResponseSuccess(c, communityList)
}