package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const polo = "polo"

//Polo test connection controller
func Polo(c *gin.Context) {
	c.String(http.StatusOK, polo)
}
