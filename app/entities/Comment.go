package entities

import (
	"time"
)

// コメントの内容を保持する
type Comment struct {
	Id        int       `json:"id,omitempty"`         // コメントID
	Owner     string    `json:"owner,omitempty"`      // 対象リポジトリのowner
	Repo      string    `json:"repo,omitempty"`       // 対象リポジトリ
	PullId    int       `json:"pull_id,omitempty"`    // プルリクのID
	Body      string    `json:"body,omitempty"`       // コメント本文
	Title     string    `json:"title,omitempty"`      // コメントのタイトル
	UserName  string    `json:"user_name,omitempty"`  // コメントをしたユーザー名
	FilePath  string    `json:"file_path,omitempty"`  // コメントがついたファイル名。issueの場合は空
	CreatedAt time.Time `json:"created_at,omitempty"` // コメントが最初に投稿された時間
	UpdatedAt time.Time `json:"updated_at,omitempty"` // コメントが最後に更新された時間
	Status    int       `json:"status,omitempty"`     // そのコメントの状態
}
