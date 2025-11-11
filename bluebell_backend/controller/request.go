package controller

import (
	"bluebell_backend/middleware"
	"bluebell_backend/pkg/errcode"

	"github.com/gin-gonic/gin"
)

func getCurrentUser(c *gin.Context) (userID uint64, err error) {
	uid, ok := c.Get(middleware.ContextUserID)
	if !ok {
		err = errcode.ErrorNotLogin
		return
	}
	userID, ok = uid.(uint64)
	if !ok {
		err = errcode.ErrorNotLogin
		return
	}
	return userID, nil
}
