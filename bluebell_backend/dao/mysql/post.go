package mysql

import (
	"bluebell_backend/model"
	"database/sql"
	"strings"

	"github.com/jmoiron/sqlx"
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

func GetPostList(page, size int64) (posts []*model.Post, err error) {
	sqlStr := `select 
	post_id, title, content, author_id, community_id, create_time from post 
	order by create_time desc 
	limit ?, ?`
	// 从第几条读，读多少
	posts = make([]*model.Post, 0, size)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

func GetPostListByIDs(ids []string) (posts []*model.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, community_id, create_time 
	from post 
	where post_id in (?) 
	order by find_in_set(post_id, ?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&posts, query, args...)
	return 
}