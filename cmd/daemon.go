package cmd

import (
	"context"
	"net/http"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/luckylu/flare/pkg/logger"

	FlareModel "github.com/luckylu/flare/internal/model"
	FlareState "github.com/luckylu/flare/internal/state"

	FlareAuth "github.com/luckylu/flare/internal/auth"
	FlareAssets "github.com/luckylu/flare/internal/resources/assets"
	FlareMDI "github.com/luckylu/flare/internal/resources/mdi"
	FlareTemplates "github.com/luckylu/flare/internal/resources/templates"

	FlareDeprecated "github.com/luckylu/flare/internal/deprecated"
	FlareEditor "github.com/luckylu/flare/internal/editor"
	FlareGuide "github.com/luckylu/flare/internal/guide"
	FlareHealth "github.com/luckylu/flare/internal/health"
	FlareHome "github.com/luckylu/flare/internal/pages/home"
	FlareRedir "github.com/luckylu/flare/internal/redir"
	FlareSettings "github.com/luckylu/flare/internal/settings"
	FlareAppearance "github.com/luckylu/flare/internal/settings/appearance"
	FlareOthers "github.com/luckylu/flare/internal/settings/others"
	FlareSearch "github.com/luckylu/flare/internal/settings/search"
	FlareTheme "github.com/luckylu/flare/internal/settings/theme"
	FlareWeather "github.com/luckylu/flare/internal/settings/weather"
)

func startDaemon(AppFlags *FlareModel.Flags) {

	if !AppFlags.DebugMode {
		gin.SetMode(gin.ReleaseMode)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	log := logger.GetLogger()

	router.Use(logger.Logger(log), gin.Recovery())

	if !AppFlags.DisableLoginMode {
		FlareAuth.RequestHandle(router)
	}

	if !AppFlags.DebugMode {
		router.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	FlareState.Init()
	FlareAssets.RegisterRouting(router)

	FlareMDI.Init()
	FlareMDI.RegisterRouting(router)

	FlareTemplates.RegisterRouting(router)
	FlareAppearance.RegisterRouting(router)
	FlareDeprecated.RegisterRouting(router)
	FlareHealth.RegisterRouting(router)
	FlareWeather.RegisterRouting(router)
	FlareHome.RegisterRouting(router)
	FlareOthers.RegisterRouting(router)
	FlareRedir.RegisterRouting(router)
	FlareSearch.RegisterRouting(router)
	FlareSettings.RegisterRouting(router)
	FlareTheme.RegisterRouting(router)

	if AppFlags.EnableEditor {
		FlareEditor.RegisterRouting(router)
		log.Println("在线编辑模块启用，可以访问 " + FlareState.RegularPages.Editor.Path + " 来进行数据编辑。")
	}

	if AppFlags.EnableGuide {
		FlareGuide.RegisterRouting(router)
		log.Println("向导模块启用，可以访问 " + FlareState.RegularPages.Guide.Path + " 来获取程序使用帮助。")
	}

	srv := &http.Server{
		Addr:              ":" + strconv.Itoa(AppFlags.Port),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       5 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("程序启动出错: %s\n", err)
		}
	}()
	log.Println("程序已启动完毕 🚀")

	<-ctx.Done()

	stop()
	log.Println("程序正在关闭中，如需立即结束请按 CTRL+C")

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("程序强制关闭: ", err)
	}

	log.Println("期待与你的再次相遇 ❤️")
}
