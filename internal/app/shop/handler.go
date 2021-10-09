/*
Package shop is a Banana Duck shop
Handles requests for managing ducks and orders
*/
package shop

import (
	"net/http"

	"github.com/labstack/echo/v4"
	. "github.com/patriya-piyawiroj/banana-ducks/internal/app/bnnd"
)

type Handler interface {
	GetAllBananaDucks(c echo.Context) error
	CreateBananaDuck(c echo.Context) error
}

type Shop struct {
	bnnds BNNDService
	// orders OrderService
}

func NewShop(bnnds BNNDService) Handler {
	return &Shop{
		bnnds: bnnds,
		// orders: orders,
	}
}

// GetAllBananaDucks is a handler for getting all banana ducks
// Takes query, sort, and pagination params
func (s *Shop) GetAllBananaDucks(c echo.Context) error {
	ducks, err := s.bnnds.GetAllBananaDucks()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ducks)
}

// CreateBananaDuck is a handler for creating a banana duck
func (s *Shop) CreateBananaDuck(c echo.Context) error {
	var d BananaDuck
	if err := c.Bind(d); err != nil {
		// TODO: Write errors to external service
		return err
	}
	if err := s.bnnds.CreateBananaDuck(d); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, d)
}
