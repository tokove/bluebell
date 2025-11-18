package mysql

import (
	"bluebell_backend/model"
	"bluebell_backend/setting"
	"testing"
)

func init() {
	cfg := &setting.MySQLConfig{
		Host:         "127.0.0.1",
		User:         "root",
		Password:     "123456",
		DB:           "bluebell",
		Port:         3306,
		MaxOpenConns: 200,
		MaxIdleConns: 50,
	}
	err := Init(cfg)
	if err != nil {
		panic("mysql Init failed")
	}
}
func TestCreatePost(t *testing.T) {
	p := &model.Post{
		ID:          1,
		AuthorID:    2,
		CommunityID: 3,
		Title:       "go test",
		Content:     "初学gotest",
	}
	err := CreatePost(p)
	if err != nil {
		t.Fatalf("CreatePost failed, err:%v\n", err)
	}
	t.Log("CreatePost success")
}
