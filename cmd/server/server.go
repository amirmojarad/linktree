package server

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	gormTracing "github.com/kostyay/gorm-opentelemetry"
	"github.com/pressly/goose/v3"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"linktree/config"
	"linktree/database"
	"linktree/internal/controller"
	"linktree/internal/logger"
	"linktree/internal/repository"
	"linktree/internal/service"
	"linktree/tracing"
)

func Run() error {
	cfg, err := config.NewConfig()
	if err != nil {
		return err
	}

	sqlDb, err := database.ConnectToPostgres(cfg)
	if err != nil {
		return err
	}

	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err = goose.Up(sqlDb, cfg.PostgresDatabase.MigrationPath); err != nil {
		return err
	}

	tracerProvider, err := tracing.GetTracerProvider(cfg)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	tracing.SetTracerProvider(tracerProvider)

	routerEngine, err := setupRouter(cfg, tracerProvider, sqlDb)
	if err != nil {
		return err
	}

	logger.GetLogger().Infof("running server on port %s ...", cfg.App.Port)

	return routerEngine.Run(fmt.Sprintf(":%s", cfg.App.Port))
}

func setupRouter(cfg *config.AppConfig, tracerProvider *trace.TracerProvider, postgresDb *sql.DB) (*gin.Engine, error) {
	psql, err := getGormDB(cfg, postgresDb)
	if err != nil {
		return nil, err
	}

	gin.SetMode(cfg.Gin.Mode)

	ginEngine := gin.New()
	ginEngine.Use(gin.Recovery())
	ginEngine.Use(gin.LoggerWithConfig(gin.LoggerConfig{SkipPaths: []string{"/health"}}))

	ginEngine.Use(otelgin.Middleware(cfg.App.Name))

	controller.SetHealthCheckRoute(ginEngine, controller.NewHealthCheck(postgresDb))

	v1Group := ginEngine.Group("/v1")

	setUserWorkspace(v1Group, psql, cfg, tracerProvider)

	return ginEngine, nil
}

func setUserWorkspace(routes gin.IRoutes, db *gorm.DB, cfg *config.AppConfig, tracer *trace.TracerProvider) {
	userRepo := repository.NewUser(db)

	userSvc := service.NewUser(cfg,
		logger.GetLogger().WithField("name", "user.service"),
		tracer.Tracer("service.user"),
		userRepo,
	)

	userCtrl := controller.NewUser(logger.GetLogger().WithField("name", "user.controller"), userSvc)

	controller.SetUserRoutes(routes, userCtrl)
}

func getGormDB(_ *config.AppConfig, postgresDb *sql.DB) (*gorm.DB, error) {
	psql, err := gorm.Open(postgres.New(postgres.Config{Conn: postgresDb}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = psql.Use(gormTracing.NewPlugin()); err != nil {
		return nil, err
	}

	return psql, nil
}
