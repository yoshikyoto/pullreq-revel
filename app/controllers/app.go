package controllers

import (
	"github.com/revel/revel"
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
	services.Get(owner, repo, number)
	json := JsonResponse{
		Status:  202,
		Message: "Accepted",
	}
	return c.RenderJson(json)
}
