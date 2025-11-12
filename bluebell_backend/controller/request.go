package controller

import (
	"bluebell_backend/dao/mysql"

	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "userID"

func getCurrentUserID(c *gin.Context) (userID uint64, err error) {
	uid, ok := c.Get(ContextUserIDKey)
	if !ok {
		err = mysql.ErrorNotLogin
		return
	}
	userID, ok = uid.(uint64)
	if !ok {
		err = mysql.ErrorNotLogin
		return
	}
	return userID, nil
}
