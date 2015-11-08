package controllers

import (
	"github.com/revel/revel"
	"pullreq/app/repos"
	"pullreq/app/services"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

type JsonResponse struct {
	Status  int
	Message string
}

func (c App) GetCommentsAsync(owner, repo string, number int) revel.Result {
	services.GetAsync(owner, repo, number)
	json := JsonResponse{
		Status:  202,
		Message: "Accepted",
	}
	return c.RenderJson(json)
}

func (c App) Recent(owner, repo string, numer int) revel.Result {
	comments, _ := repos.Recent()
	return c.RenderJson(comments)
}
