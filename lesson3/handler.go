package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	UC UseCase
}

func NewHandler(e *echo.Echo, uc UseCase) {
	handler := &Handler{
		UC: uc,
	}

	e.GET("/story", handler.GetStory)
}

func (h *Handler) GetStory(c echo.Context) error {
	option := c.QueryParam("option")
	if option == "" {
		option = "intro"
	}

	story, err := h.UC.GetStory(option)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "<p>something went wrong</p>")
	}
	if story != nil {
		return c.HTML(http.StatusOK, *story)
	}
	return c.HTML(http.StatusNotFound, "<p>sorry! no new story found with your choice</p>")
}
