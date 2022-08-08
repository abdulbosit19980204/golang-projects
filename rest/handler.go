package rest

import (
	"transaction/repo"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	r     *gin.Engine
	store repo.Repo
}

func New(s repo.Repo) Handler {
	r := gin.Default()

	h := Handler{
		r:     r,
		store: s,
	}

	h.initRoute()

	return h
}

func (h *Handler) Run() {
	h.r.Run()
}

func (h *Handler) initRoute() {
	h.r.GET("/users", h.user)
}

/*
func (h *Handler) initRoute() {
	h.r.GET("/user-filter", h.user)
}
*/
