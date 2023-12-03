package appearance

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"

	FlareAuth "github.com/luckylu/flare/internal/auth"
	FlareData "github.com/luckylu/flare/internal/data"
	FlareModel "github.com/luckylu/flare/internal/model"
	FlareState "github.com/luckylu/flare/internal/state"
)

func RegisterRouting(router *gin.Engine) {
	router.GET(FlareState.SettingPages.Appearance.Path, FlareAuth.AuthRequired, pageAppearance)
	router.POST(FlareState.SettingPages.Appearance.Path, FlareAuth.AuthRequired, updateAppearanceOptions)
}

func updateAppearanceOptions(c *gin.Context) {

	type UpdateBody struct {
		OptionTitle  string `form:"title"`
		OptionFooter string `form:"footer"`

		OptionOpenAppNewTab      bool `form:"open-app-newtab"`
		OptionOpenBookmarkNewTab bool `form:"open-bookmark-newtab"`

		OptionShowTitle     bool   `form:"show-title"`
		OptionGreetings     string `form:"greetings"`
		OptionShowDateTime  bool   `form:"show-datetime"`
		OptionShowApps      bool   `form:"show-apps"`
		OptionShowBookmarks bool   `form:"show-bookmarks"`
		HideSettingsButton  bool   `form:"hide-settings-button"`
		HideHelpButton      bool   `form:"hide-help-button"`
		EnableEncryptedLink bool   `form:"enable-encrypted-link"`
		IconMode            string `form:"icon-mode"`
		KeepLetterCase      bool   `form:"keep-letter-case"`

		OptionCustomDay     string `form:"custom-day"`
		OptionCustomMonth   string `form:"custom-month"`
		BackgroundImageBlur string `form:"background-image-blur"`
	}

	var body UpdateBody
	if c.ShouldBind(&body) != nil {
		c.PureJSON(http.StatusForbidden, "提交数据缺失")
		return
	}

	uploadImage(c)

	var update FlareModel.Application
	update.Title = body.OptionTitle
	update.Footer = body.OptionFooter
	update.OpenAppNewTab = body.OptionOpenAppNewTab
	update.OpenBookmarkNewTab = body.OptionOpenBookmarkNewTab
	update.ShowTitle = body.OptionShowTitle
	update.Greetings = body.OptionGreetings
	update.ShowDateTime = body.OptionShowDateTime
	update.ShowApps = body.OptionShowApps
	update.ShowBookmarks = body.OptionShowBookmarks
	update.HideSettingsButton = body.HideSettingsButton
	update.HideHelpButton = body.HideHelpButton
	update.EnableEncryptedLink = body.EnableEncryptedLink
	update.KeepLetterCase = body.KeepLetterCase
	update.BackgroundImageBlur = body.BackgroundImageBlur

	requestIconMode := strings.ToUpper(body.IconMode)
	if requestIconMode != "DEFAULT" && requestIconMode != "FILLING" {
		update.IconMode = "DEFAULT"
	} else {
		update.IconMode = requestIconMode
	}

	FlareData.UpdateAppearance(update)
	FlareState.UpdatePagePalettes()

	pageAppearance(c)
}

func uploadImage(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println("error:" + err.Error())
		return
	}

	uploadedFile, err := file.Open()
	if err != nil {
		fmt.Println("error:" + err.Error())
		return
	}
	defer uploadedFile.Close()

	rootDir, _ := os.Getwd()
	destinationPath := filepath.Join(rootDir, "bg.jpg")
	destinationFile, err := os.Create(destinationPath)
	if err != nil {
		fmt.Println("error:" + err.Error())
		return
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, uploadedFile)
	if err != nil {
		fmt.Println("error:" + err.Error())
		return
	}

}

func pageAppearance(c *gin.Context) {
	options := FlareData.GetAllSettingsOptions()

	IconModeDefault := options.IconMode == "DEFAULT"
	IconModeFilling := options.IconMode == "FILLING"

	c.HTML(
		http.StatusOK,
		"settings.html",
		gin.H{
			"DebugMode":       FlareState.AppFlags.DebugMode,
			"PageInlineStyle": FlareState.GetPageInlineStyle(),

			"PageName":       "Appearance",
			"PageAppearance": FlareState.GetAppBodyStyle(),
			"SettingPages":   FlareState.SettingPages,
			"SettingsURI":    FlareState.RegularPages.Settings.Path,

			"OptionTitle":               options.Title,
			"OptionFooter":              template.HTML(options.Footer),
			"OptionOpenAppNewTab":       options.OpenAppNewTab,
			"OptionOpenBookmarkNewTab":  options.OpenBookmarkNewTab,
			"OptionShowTitle":           options.ShowTitle,
			"OptionGreetings":           options.Greetings,
			"OptionShowDateTime":        options.ShowDateTime,
			"OptionShowApps":            options.ShowApps,
			"OptionShowBookmarks":       options.ShowBookmarks,
			"OptionHideSettingsButton":  options.HideSettingsButton,
			"OptionHideHelpButton":      options.HideHelpButton,
			"OptionEnableEncryptedLink": options.EnableEncryptedLink,
			"OptionKeepLetterCase":      options.KeepLetterCase,
			"OptionIconModeDefault":     IconModeDefault,
			"OptionIconModeFilling":     IconModeFilling,
			"OptionBackgroundImageBlur": options.BackgroundImageBlur,
		},
	)
}
