package repos

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"pullreq/app/entities"
)

// データベースにコメントをInsertする。
// エラーが発生した場合は戻り値としてエラーを返す
func Create(comment entities.Comment) error {
	revel.TRACE.Println("Create")
	db, err := gorm.Open("mysql", revel.Config.StringDefault("db.uri", ""))
	if err != nil {
		revel.ERROR.Println("データベースに接続できませんでした")
		return err
	}
	db.DB()
	db.Create(comment)
	return nil
}

// 該当するプルリクエストのデータをデータベースから取得する
func FindCommentList(owner, repo string, number int) ([]entities.Comment, error) {
	revel.TRACE.Println("repos.FindCommentLis")
	db, err := gorm.Open("mysql", revel.Config.StringDefault("db.uri", ""))
	if err != nil {
		revel.ERROR.Println("データベースに接続できませんでした")
		return nil, err
	}
	db.DB()
	comments := []entities.Comment{}
	db.Where(&entities.Comment{
		Owner:  owner,
		Repo:   repo,
		PullId: number,
	}).Find(&comments)
	return comments, nil
}
