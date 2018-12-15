package handler

import (
	"net/http"

	"github.com/gobuffalo/uuid"

	"github.com/uqichi/goffee-server/models"

	"github.com/labstack/echo"
	"github.com/uqichi/goffee-server/usecases"
)

type Handler struct {
	useCase usecases.CoffeeUseCase
}

func NewHandler(useCase usecases.CoffeeUseCase) *Handler {
	return &Handler{useCase}
}

func (h *Handler) GetCoffee(c echo.Context) error {
	id := c.Param("id")
	coffee, err := h.useCase.Find(c.Request().Context(), uuid.FromStringOrNil(id))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, coffee)
}

func (h *Handler) CreateCoffee(c echo.Context) error {
	req := new(models.Coffee)
	err := c.Bind(req)
	if err != nil {
		return err
	}
	coffee, err := h.useCase.Create(c.Request().Context(), req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, coffee)
}
