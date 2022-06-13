package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getFontFamily(c *gin.Context) string {
	ff, err := c.Cookie("ff")
	if err == http.ErrNoCookie {
		ff = "Times New Roman"
		c.SetCookie("ff", "Times New Roman", 60*60*24*365*100, "/", "", false, false)
	}
	return ff
}

func handleSettings(c *gin.Context) {
	c.HTML(http.StatusOK, "settings.html", gin.H{
		"FontFamily": getFontFamily(c),
	})
}