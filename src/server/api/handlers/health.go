package handlers

import (
	"fmt"
	"net/http"

	"github.com/garrickedd/ReLibca/src/server/api/helper"
	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(c *gin.Context) {
	// c.JSON(http.StatusOK, "Working!")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("Working!", true, 0))
}

func (h *HealthHandler) HealthPost(c *gin.Context) {
	c.JSON(http.StatusOK, "Working post!")
}

func (h *HealthHandler) HealthPostById(c *gin.Context) {
	id := c.Params.ByName("id")
	c.JSON(http.StatusOK, fmt.Sprintf("Working post by id: %s!", id))
}
