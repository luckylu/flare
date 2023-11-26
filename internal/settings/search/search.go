package search

import (
	"net/http"

	"github.com/gin-gonic/gin"

	FlareAuth "github.com/luckylu/flare/internal/auth"
	FlareData "github.com/luckylu/flare/internal/data"
	FlareState "github.com/luckylu/flare/internal/state"
)

func RegisterRouting(router *gin.Engine) {

	router.GET(FlareState.SettingPages.Search.Path, FlareAuth.AuthRequired, pageSearch)
	router.POST(FlareState.SettingPages.Search.Path, FlareAuth.AuthRequired, updateSearchOptions)

}

func updateSearchOptions(c *gin.Context) {

	type UpdateBody struct {
		ShowSearchComponent     bool `form:"show-search-component"`
		DisabledSearchAutoFocus bool `form:"disabled-search-auto-focus"`
	}

	var body UpdateBody
	if c.ShouldBind(&body) != nil {
		c.PureJSON(http.StatusForbidden, "提交数据缺失")
		return
	}

	FlareData.UpdateSearch(body.ShowSearchComponent, body.DisabledSearchAutoFocus)

	pageSearch(c)
}

func pageSearch(c *gin.Context) {
	options := FlareData.GetAllSettingsOptions()

	c.HTML(
		http.StatusOK,
		"settings.html",
		gin.H{
			"DebugMode":       FlareState.AppFlags.DebugMode,
			"PageInlineStyle": FlareState.GetPageInlineStyle(),

			"PageName":       "Search",
			"PageAppearance": FlareState.GetAppBodyStyle(),
			"SettingPages":   FlareState.SettingPages,
			"SettingsURI":    FlareState.RegularPages.Settings.Path,

			"ShowSearchComponent":     options.ShowSearchComponent,
			"DisabledSearchAutoFocus": options.DisabledSearchAutoFocus,

			"OptionTitle": options.Title,
		},
	)
}
