package models

import (
	"encoding/json"
	"fmt"
	"time"
)

type Comment struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (c *Comment) Key() string {
	return fmt.Sprintf("comment:%d", c.ID)
}

func (c *Comment) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
