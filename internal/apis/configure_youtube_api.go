// This file is safe to edit. Once it exists it will not be overwritten

package apis

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/youtube-api-golang/internal/apis/operations"
	"github.com/youtube-api-golang/internal/handler"
	"github.com/youtube-api-golang/internal/middleware"
)

//go:generate swagger generate server --target ../../../youtube-api-golang --name YoutubeAPI --spec ../../api/swagger.yml --model-package internal/models --server-package internal/apis --principal interface{}

func configureFlags(api *operations.YoutubeAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.YoutubeAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	//mounting handler
	mountHandler := handler.NewHandler()

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BearerAuth = middleware.ValidateHeader

	// health check
	api.HealthGetHealthHandler = mountHandler.GetHealthHandler()

	// user
	api.UserPostUserRegisterHandler = mountHandler.RegisterUserHandler()
	api.UserPutUserIDHandler = mountHandler.UpdateUserHandler()
	api.UserDeleteUserIDHandler = mountHandler.DeleteUserHandler()
	api.UserGetUserIDHandler = mountHandler.GetUserByIDHandler()

	// auth
	api.AuthPostUserLoginHandler = mountHandler.LoginHandler()

	// subscription
	api.SubscriptionPatchSubUserIDHandler = mountHandler.SubscribeHandler()
	api.SubscriptionPatchUnsubUserIDHandler = mountHandler.UnsubscribeHandler()

	// video
	api.VideoPostVideoHandler = mountHandler.AddVideoHandler()
	api.VideoPutVideoIDHandler = mountHandler.UpdateVideoHandler()
	api.VideoDeleteVideoIDHandler = mountHandler.DeleteVideoHandler()
	api.VideoGetVideoIDHandler = mountHandler.GetVideoByIDHandler()
	api.VideoPatchVideoViewIDHandler = mountHandler.UpdateViewHandler()
	api.VideoGetVideoRandomHandler = mountHandler.GetRandomVideosHandler()
	api.VideoGetVideoTrendHandler = mountHandler.GetTrendingVideosHandler()
	api.VideoGetVideoSubHandler = mountHandler.GetVideosFromSubscribedChannelsHandler()
	api.VideoGetVideoTagsHandler = mountHandler.GetVideosByTagsHandler()
	api.VideoGetVideoSearchHandler = mountHandler.SearchVideosHandler()

	// comment
	api.CommentPostCommentHandler = mountHandler.AddCommentHandler()
	api.CommentDeleteCommentHandler = mountHandler.DeleteCommentHandler()
	api.CommentGetCommentsVideoIDHandler = mountHandler.GetCommentsByVideoIDHandler()

	api.PreServerShutdown = func() {}

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
