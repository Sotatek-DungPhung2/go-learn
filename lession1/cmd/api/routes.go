package main

import (
	middlewares "go-learn/internal/middlewares"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	// Initialize a new httprouter router instance.
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/login", app.authHandler)
	// Return the httprouter instance.

	// cors
	// logging
	return middlewares.RecoverPanicMiddleware(middlewares.RateLimiterMiddleware(router))
}
