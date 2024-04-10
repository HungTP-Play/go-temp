package app

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type AppConfig struct {
	Port           string
	LogLevel       string
	MetricsEnabled bool
	MetricsPort    string
	MetricsPath    string
}

var defaultAppConfig = &AppConfig{
	Port:           "3000",
	LogLevel:       "debug",
	MetricsEnabled: false,
	MetricsPort:    "9090",
	MetricsPath:    "/metrics",
}

func MergeAppConfigsWithEnv() *AppConfig {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultAppConfig.Port
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = defaultAppConfig.LogLevel
	}

	metricEnabled := os.Getenv("METRICS_ENABLE") == "true"

	metricPort := os.Getenv("METRICS_PORT")
	if metricPort == "" {
		metricPort = defaultAppConfig.MetricsPort
	}

	metricPath := os.Getenv("METRICS_PATH")
	if metricPath == "" {
		metricPath = defaultAppConfig.MetricsPath
	}

	return &AppConfig{
		Port:           port,
		LogLevel:       logLevel,
		MetricsEnabled: metricEnabled,
		MetricsPort:    metricPort,
		MetricsPath:    metricPath,
	}
}

type App struct {
	Configs *AppConfig
	Context context.Context
}

func NewApp(configs *AppConfig) *App {
	return &App{Configs: configs, Context: context.Background()}
}

func (a *App) Start() {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	fmt.Println("Listening on port " + a.Configs.Port)
	http.ListenAndServe(":"+a.Configs.Port, r)
}
