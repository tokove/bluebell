package mysql

import (
	"bluebell_backend/model"
	"database/sql"
)

func CreatePost(post *model.Post) error {
	sqlStr := `insert into post(post_id, title, content, author_id, community_id) values(?, ?, ?, ?, ?)`
	_, err := db.Exec(sqlStr, post.ID, post.Title, post.Content, post.AuthorID, post.CommunityID)
	return err
}

func GetPostDetailByID(id uint64) (p *model.Post, err error) {
	p = new(model.Post)
	sqlStr := `select post_id, title, content, author_id, community_id, create_time from post where post_id = ? `
	
	if err = db.Get(p, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return
}
