package templates

import (
	"embed"
	"html/template"

	"github.com/gin-gonic/gin"

	FlareState "github.com/luckylu/flare/internal/state"
)

//go:embed html
var TPL embed.FS

func RegisterRouting(router *gin.Engine) {

	if FlareState.AppFlags.DebugMode {
		router.LoadHTMLGlob("embed/templates/*.html")
		return
	}

	templ := template.Must(template.New("").ParseFS(TPL, "html/*.html"))
	router.SetHTMLTemplate(templ)
}
