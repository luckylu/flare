package assets

import (
	"crypto/md5" //#nosec
	"embed"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	FlareState "github.com/luckylu/flare/internal/state"
)

//go:embed favicon.ico
var Favicon embed.FS

func RegisterRouting(router *gin.Engine) {

	router.Use(optimizeResourceCacheTime())

	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		c.FileFromFS("favicon.ico", http.FS(Favicon))
	})

	router.GET("/bg.jpg", func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age=31536000")
		rootDir, _ := os.Getwd()
		filePath := filepath.Join(rootDir, "bg.jpg")

		// 返回静态文件
		c.File(filePath)
	})

	if FlareState.AppFlags.DebugMode {
		router.StaticFS("/assets/css", http.Dir("embed/assets/css"))
		return
	}
}

// ViewHandler support dist handler from UI
// https://github.com/gin-gonic/gin/issues/1222
func optimizeResourceCacheTime() gin.HandlerFunc {
	data := []byte(time.Now().String())
	/* #nosec */
	etag := fmt.Sprintf("W/%x", md5.Sum(data))
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RequestURI, "/assets/") ||
			strings.HasPrefix(c.Request.RequestURI, "/favicon.ico") {
			c.Header("Cache-Control", "public, max-age=31536000")
			c.Header("ETag", etag)

			if match := c.GetHeader("If-None-Match"); match != "" {
				if strings.Contains(match, etag) {
					c.Status(http.StatusNotModified)
					return
				}
			}
		}
	}
}
