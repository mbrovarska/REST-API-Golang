package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)


type Response struct {
	Status    string    `json:"status" example:"OK"`
	Timestamp time.Time `json:"timestamp" example:"2024-01-01T12:00:00Z"`
	Version   string    `json:"version" example:"1.0.0"` 
}

// CheckHealth godoc
// @Summary      Health Check
// @Description  Check if the API is running and return version info
// @Tags         system
// @Produce      json
// @Success      200  {object}  Response
// @Router       /health [get]
func CheckHealth(version string) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(http.StatusOK, Response{
            Status: "OK",
			Timestamp: time.Now(),
            Version: version,
        })
    }
}