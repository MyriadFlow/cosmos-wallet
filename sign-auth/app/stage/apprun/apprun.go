// Package apprun provides method to Start http server of gin
package apprun

import (
	"time"

	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/api"
	"github.com/MyriadFlow/cosmos-wallet/sign-auth/pkg/env"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run() {
	ginApp := gin.Default()

	corsM := cors.New(cors.Config{AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{env.MustGetEnv("ALLOWED_ORIGIN")}})
	ginApp.Use(corsM)
	api.ApplyRoutes(ginApp)
	port := env.MustGetEnv("APP_PORT")
	err := ginApp.Run(":" + port)
	if err != nil {
		logo.Fatalf("failed to serve app on port %s: %s", port, err)
	}
}
