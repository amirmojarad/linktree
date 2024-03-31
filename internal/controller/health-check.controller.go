package controller

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheck struct {
	postgresDb *sql.DB
}

func NewHealthCheck(postgresDb *sql.DB) *HealthCheck {
	return &HealthCheck{postgresDb: postgresDb}
}

func (h HealthCheck) Health(ctx *gin.Context) {
	if err := h.postgresDb.PingContext(ctx); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "Not OK", "error": err.Error()})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "OK"})
}
