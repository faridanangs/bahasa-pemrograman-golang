package app

import (
	"net/http"
	"restfullapi_news/middleware"
)

func NewServerNangs(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}
}
