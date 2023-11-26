package others

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"

	FlareAuth "github.com/luckylu/flare/internal/auth"
	FlareData "github.com/luckylu/flare/internal/data"
	FlareState "github.com/luckylu/flare/internal/state"
	FlareVersion "github.com/luckylu/flare/internal/version"
)

func RegisterRouting(router *gin.Engine) {
	router.GET(FlareState.SettingPages.Others.Path, pageOthers)
}

func pageOthers(c *gin.Context) {
	options := FlareData.GetAllSettingsOptions()

	isLogined := false

	if !FlareState.AppFlags.DisableLoginMode {
		isLogined = FlareAuth.CheckUserIsLogin(c)
	} else {
		isLogined = true
	}

	c.HTML(
		http.StatusOK,
		"settings.html",
		gin.H{
			"DebugMode":        FlareState.AppFlags.DebugMode,
			"DisableLoginMode": FlareState.AppFlags.DisableLoginMode,
			"UserIsLogin":      isLogined,
			"UserName":         FlareAuth.GetUserName(c),
			"LoginDate":        FlareAuth.GetUserLoginDate(c),

			"PageInlineStyle": FlareState.GetPageInlineStyle(),
			"PageAppearance":  FlareState.GetAppBodyStyle(),
			"SettingsURI":     FlareState.RegularPages.Settings.Path,
			"LoginURI":        FlareState.MiscPages.Login.Path,
			"LogoutURI":       FlareState.MiscPages.Logout.Path,

			"PageName":     "Others",
			"SettingPages": FlareState.SettingPages,
			"OptionTitle":  options.Title,
			"Version":      FlareVersion.Version,
			"BuildDate":    FlareVersion.BuildDate,
			"COMMIT":       FlareVersion.Commit,

			"OptionFooter": template.HTML(options.Footer),
		},
	)
}
