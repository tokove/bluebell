package model

type Community struct {
	ID   uint64 `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}
