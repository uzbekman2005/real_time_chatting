package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetPage(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Query("page"))
}

func GetLimit(c *gin.Context) (int, error) {
	return strconv.Atoi(c.Query("limit"))
}
