package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	UC UseCase
}

func NewHandler(e *echo.Echo, uc UseCase) {
	handler := &Handler{
		UC: uc,
	}

	e.GET("/:Path", handler.SlugHandler)
}

func (h *Handler) SlugHandler(c echo.Context) error {
	Path := c.Param("Path")

	url, err := h.UC.GetRedirectUrl(Path)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "something went wrong",
		})
	}

	if url != nil {
		return c.Redirect(http.StatusFound, *url)
	}
	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "no redirect url found",
	})
}
