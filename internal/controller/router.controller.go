package controller

import "github.com/gin-gonic/gin"

func SetHealthCheckRoute(routes gin.IRoutes, ctrl *HealthCheck) {
	routes.GET("/health", ctrl.Health)
}

func SetUserRoutes(routes gin.IRoutes, ctrl *User) {

}
