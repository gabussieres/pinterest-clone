// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/rs/cors"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"pinterest-clone/controllers"
	"pinterest-clone/restapi/operations"
	"pinterest-clone/restapi/operations/pinterest"
)

//go:generate swagger generate server --target ../../pinterest-clone --name PinterestClone --spec ../swagger.yml

func configureFlags(api *operations.PinterestCloneAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PinterestCloneAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.PinterestDeletePinHandler == nil {
		api.PinterestDeletePinHandler = pinterest.DeletePinHandlerFunc(func(params pinterest.DeletePinParams) middleware.Responder {
			err := controllers.HandleDeletePinCall(params)
			if err != nil {
				return pinterest.NewFeedDefault(500)
			}
			return pinterest.NewDeletePinOK()
		})
	}
	if api.PinterestFeedHandler == nil {
		api.PinterestFeedHandler = pinterest.FeedHandlerFunc(func(params pinterest.FeedParams) middleware.Responder {
			feed, err := controllers.HandleFeedCall(params)
			if err != nil {
				return pinterest.NewFeedDefault(500)
			}
			return pinterest.NewFeedOK().WithPayload(feed)
		})
	}
	if api.PinterestPinHandler == nil {
		api.PinterestPinHandler = pinterest.PinHandlerFunc(func(params pinterest.PinParams) middleware.Responder {
			pin, err := controllers.HandlePinCall(params)
			if err != nil {
				return pinterest.NewPinDefault(500)
			}
			return pinterest.NewPinOK().WithPayload(pin)
		})
	}
	if api.PinterestPinsHandler == nil {
		api.PinterestPinsHandler = pinterest.PinsHandlerFunc(func(params pinterest.PinsParams) middleware.Responder {
			pins, err := controllers.HandlePinsCall(params)
			if err != nil {
				return pinterest.NewPinsDefault(500)
			}
			return pinterest.NewPinsOK().WithPayload(pins)
		})
	}
	if api.PinterestUserHandler == nil {
		api.PinterestUserHandler = pinterest.UserHandlerFunc(func(params pinterest.UserParams) middleware.Responder {
			user, err := controllers.HandleUserCall(params)
			if err != nil {
				return pinterest.NewUserDefault(500)
			}
			return pinterest.NewUserOK().WithPayload(user)
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}
