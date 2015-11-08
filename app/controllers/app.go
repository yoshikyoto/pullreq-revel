package controllers

import (
	"github.com/revel/revel"
	"pullreq/app/repos"
	"pullreq/app/services"
)

type App struct {
	*revel.Controller
}

// トップページ
func (c App) Index() revel.Result {
	return c.Render()
}

// ステータスだけ返したいときのJson
type JsonResponse struct {
	Status  int
	Message string
}

// GitHubからコメントを取得してくる
func (c App) GetCommentsAsync(owner, repo string, number int) revel.Result {
	services.GetAsync(owner, repo, number)
	json := JsonResponse{
		Status:  202,
		Message: "Accepted",
	}
	return c.RenderJson(json)
}

// リポジトリのnumberに該当するコメントをjsonで返す
func (c App) GetCommentList(owner, repo string, number int) revel.Result {
	revel.TRACE.Printf("GetCommentList %s/%s/%d", owner, repo, number)
	comments, _ := repos.FindCommentList(owner, repo, number)
	return c.RenderJson(comments)
}
