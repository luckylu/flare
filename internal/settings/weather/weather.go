package weather

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	FlareAuth "github.com/luckylu/flare/internal/auth"
	FlareData "github.com/luckylu/flare/internal/data"
	FlareModel "github.com/luckylu/flare/internal/model"
	FlareState "github.com/luckylu/flare/internal/state"
	weather "github.com/soulteary/funny-china-weather"
)

type RemoteWeatherResponse struct {
	Data struct {
		Observe struct {
			Degree        string `json:"degree"`
			Humidity      string `json:"humidity"`
			Precipitation string `json:"precipitation"`
			Pressure      string `json:"pressure"`
			UpdateTime    string `json:"update_time"`
			Weather       string `json:"weather"`
			WeatherCode   string `json:"weather_code"`
			WeatherShort  string `json:"weather_short"`
			WindDirection string `json:"wind_direction"`
			WindPower     string `json:"wind_power"`
		} `json:"observe"`
	} `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type WeatherQueryParams struct {
	Location string `form:"location"`
}

func GetWeatherInfo(location string) (response FlareModel.Weather, desc string, err error) {

	code, degree, humidity, lastUpdate, fetchRemoteErr := weather.GetWeatherByLocation(location)

	if fetchRemoteErr != nil {
		return response, "获取远程数据失败", errors.New("获取远程数据失败")
	}

	hour, _, _ := time.Now().Clock()
	isDay := hour >= 5 && hour <= 18

	conditionCode, conditionText := weather.GetWeatherIconByCode(code)

	const _WEATHER_DATA_CACHE_TIME = 60 * 10 // 10 minutes

	response.ExternalLastUpdate = lastUpdate
	response.Degree = degree
	response.IsDay = isDay
	response.ConditionCode = conditionCode
	response.ConditionText = conditionText
	response.Humidity = humidity
	response.Expires = time.Now().Unix() + _WEATHER_DATA_CACHE_TIME

	return response, "接口正常", nil
}

func RegisterRouting(router *gin.Engine) {
	router.GET(FlareState.SettingPages.Weather.Path, FlareAuth.AuthRequired, pageHome)
	if !FlareState.AppFlags.EnableOfflineMode {
		router.POST(FlareState.SettingPages.Weather.Path, FlareAuth.AuthRequired, updateWeatherOptions)
		router.POST(FlareState.SettingPagesAPI.WeatherTest.Path, FlareAuth.AuthRequired, testWeatherFetch)
	}
}

func pageHome(c *gin.Context) {

	render(c, "")

}

func updateWeatherOptions(c *gin.Context) {

	type UpdateBody struct {
		Location    string `form:"location"`
		ShowWeather bool   `form:"show"`
	}

	var body UpdateBody
	if c.ShouldBind(&body) != nil {
		c.PureJSON(http.StatusForbidden, "提交数据缺失")
		return
	}

	FlareData.UpdateWeatherAndLocation(body.ShowWeather, body.Location)

	render(c, "")
}

func testWeatherFetch(c *gin.Context) {
	location, _ := FlareData.GetLocationAndWeatherShow()

	_, desc, _ := GetWeatherInfo(location)

	render(c, desc)
}

func render(c *gin.Context, testResult string) {
	location, showWeather := FlareData.GetLocationAndWeatherShow()
	options := FlareData.GetAllSettingsOptions()

	c.HTML(
		http.StatusOK,
		"settings.html",
		gin.H{
			"DebugMode":         FlareState.AppFlags.DebugMode,
			"PageInlineStyle":   FlareState.GetPageInlineStyle(),
			"ShowWeatherModule": !FlareState.AppFlags.EnableOfflineMode && showWeather,
			"ShowWeather":       showWeather,

			"PageName":       "Weather",
			"PageAppearance": FlareState.GetAppBodyStyle(),
			"SettingPages":   FlareState.SettingPages,
			"SettingsURI":    FlareState.RegularPages.Settings.Path,
			"OptionTitle":    options.Title,

			"SettingPagesAPI": FlareState.SettingPagesAPI,
			"Location":        location,
			"TestResult":      testResult,
		},
	)
}
