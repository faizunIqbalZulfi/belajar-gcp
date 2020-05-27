package presenter

import (
	echo "github.com/labstack/echo/v4"
)

// Mount routes for Lender domain API.
func (h *HTTPHandler) Mount(g *echo.Group) {
	g.GET("/users", h.FindAllUser)
}
