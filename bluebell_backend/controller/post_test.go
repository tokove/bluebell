package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	body := `{
		"title": "go test",
		"content": "写test测试单元",
		"community_id": "1"
	}`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	// // 方法一
	// assert.Equal(t, 200, w.Code)
	// assert.Contains(t, w.Body.String(), "未登录")

	// 方法二
	res := new(ResponseData)
	if err := json.Unmarshal(w.Body.Bytes(), res); err != nil {
		t.Fatalf("json.Unmarshal w.Body failed, err:%v\n", err)
	}
	assert.Equal(t, res.Code, CodeNotLogin)
}
