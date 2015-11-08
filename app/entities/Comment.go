package entities

import (
	"time"
)

// コメントの内容を保持する
type Comment struct {
	Id        int       `json:"id,omitempty"`         // コメントID
	Body      string    `json:"body,omitempty"`       // コメント本文
	UserName  string    `json:"user_name,omitempty"`  // コメントをしたユーザー名
	FilePath  string    `json:"file_path,omitempty"`  // コメントがついたファイル名。issueの場合は空
	CreatedAt time.Time `json:"created_at,omitempty"` // コメントが最初に投稿された時間
	UpdatedAt time.Time `json:"updated_at,omitempty"` // コメントが最後に更新された時間
}
