package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/harshau007/listmanager/internal/listmanager"
	"github.com/harshau007/listmanager/internal/models"
)

func New(lm *listmanager.Manager) *gin.Engine {
	r := gin.Default()
	r.POST("/add", func(c *gin.Context) {
		var req models.AddRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "invalid JSON"})
			return
		}
		list := lm.Add(float64(req.Number))
		c.JSON(http.StatusOK, models.AddResponse{Success: true, List: list, Message: "processed"})
	})
	r.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.ListResponse{List: lm.List()})
	})
	r.POST("/reset", func(c *gin.Context) {
		lm = listmanager.New()
		c.JSON(http.StatusOK, gin.H{"success": true, "list": []int{}})
	})
	return r
}
