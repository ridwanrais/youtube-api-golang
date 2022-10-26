// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/youtube-api-golang/internal/apis/operations/auth"
	"github.com/youtube-api-golang/internal/apis/operations/comment"
	"github.com/youtube-api-golang/internal/apis/operations/health"
	"github.com/youtube-api-golang/internal/apis/operations/like"
	"github.com/youtube-api-golang/internal/apis/operations/subscription"
	"github.com/youtube-api-golang/internal/apis/operations/user"
	"github.com/youtube-api-golang/internal/apis/operations/video"
)

// NewYoutubeAPIAPI creates a new YoutubeAPI instance
func NewYoutubeAPIAPI(spec *loads.Document) *YoutubeAPIAPI {
	return &YoutubeAPIAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		CommentDeleteCommentHandler: comment.DeleteCommentHandlerFunc(func(params comment.DeleteCommentParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation comment.DeleteComment has not yet been implemented")
		}),
		UserDeleteUserIDHandler: user.DeleteUserIDHandlerFunc(func(params user.DeleteUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUserID has not yet been implemented")
		}),
		VideoDeleteVideoIDHandler: video.DeleteVideoIDHandlerFunc(func(params video.DeleteVideoIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation video.DeleteVideoID has not yet been implemented")
		}),
		CommentGetCommentsVideoIDHandler: comment.GetCommentsVideoIDHandlerFunc(func(params comment.GetCommentsVideoIDParams) middleware.Responder {
			return middleware.NotImplemented("operation comment.GetCommentsVideoID has not yet been implemented")
		}),
		HealthGetHealthHandler: health.GetHealthHandlerFunc(func(params health.GetHealthParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation health.GetHealth has not yet been implemented")
		}),
		UserGetUserIDHandler: user.GetUserIDHandlerFunc(func(params user.GetUserIDParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserID has not yet been implemented")
		}),
		VideoGetVideoIDHandler: video.GetVideoIDHandlerFunc(func(params video.GetVideoIDParams) middleware.Responder {
			return middleware.NotImplemented("operation video.GetVideoID has not yet been implemented")
		}),
		VideoGetVideoRandomHandler: video.GetVideoRandomHandlerFunc(func(params video.GetVideoRandomParams) middleware.Responder {
			return middleware.NotImplemented("operation video.GetVideoRandom has not yet been implemented")
		}),
		VideoGetVideoSearchHandler: video.GetVideoSearchHandlerFunc(func(params video.GetVideoSearchParams) middleware.Responder {
			return middleware.NotImplemented("operation video.GetVideoSearch has not yet been implemented")
		}),
		VideoGetVideoSubHandler: video.GetVideoSubHandlerFunc(func(params video.GetVideoSubParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation video.GetVideoSub has not yet been implemented")
		}),
		VideoGetVideoTagsHandler: video.GetVideoTagsHandlerFunc(func(params video.GetVideoTagsParams) middleware.Responder {
			return middleware.NotImplemented("operation video.GetVideoTags has not yet been implemented")
		}),
		VideoGetVideoTrendHandler: video.GetVideoTrendHandlerFunc(func(params video.GetVideoTrendParams) middleware.Responder {
			return middleware.NotImplemented("operation video.GetVideoTrend has not yet been implemented")
		}),
		LikePatchDislikeVideoIDHandler: like.PatchDislikeVideoIDHandlerFunc(func(params like.PatchDislikeVideoIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation like.PatchDislikeVideoID has not yet been implemented")
		}),
		LikePatchLikeVideoIDHandler: like.PatchLikeVideoIDHandlerFunc(func(params like.PatchLikeVideoIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation like.PatchLikeVideoID has not yet been implemented")
		}),
		SubscriptionPatchSubUserIDHandler: subscription.PatchSubUserIDHandlerFunc(func(params subscription.PatchSubUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation subscription.PatchSubUserID has not yet been implemented")
		}),
		SubscriptionPatchUnsubUserIDHandler: subscription.PatchUnsubUserIDHandlerFunc(func(params subscription.PatchUnsubUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation subscription.PatchUnsubUserID has not yet been implemented")
		}),
		VideoPatchVideoViewIDHandler: video.PatchVideoViewIDHandlerFunc(func(params video.PatchVideoViewIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation video.PatchVideoViewID has not yet been implemented")
		}),
		CommentPostCommentHandler: comment.PostCommentHandlerFunc(func(params comment.PostCommentParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation comment.PostComment has not yet been implemented")
		}),
		AuthPostUserLoginHandler: auth.PostUserLoginHandlerFunc(func(params auth.PostUserLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostUserLogin has not yet been implemented")
		}),
		UserPostUserRegisterHandler: user.PostUserRegisterHandlerFunc(func(params user.PostUserRegisterParams) middleware.Responder {
			return middleware.NotImplemented("operation user.PostUserRegister has not yet been implemented")
		}),
		VideoPostVideoHandler: video.PostVideoHandlerFunc(func(params video.PostVideoParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation video.PostVideo has not yet been implemented")
		}),
		UserPutUserIDHandler: user.PutUserIDHandlerFunc(func(params user.PutUserIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.PutUserID has not yet been implemented")
		}),
		VideoPutVideoIDHandler: video.PutVideoIDHandlerFunc(func(params video.PutVideoIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation video.PutVideoID has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*YoutubeAPIAPI Youtube API */
type YoutubeAPIAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator

	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator

	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// BearerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	BearerAuth func(string) (interface{}, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// CommentDeleteCommentHandler sets the operation handler for the delete comment operation
	CommentDeleteCommentHandler comment.DeleteCommentHandler
	// UserDeleteUserIDHandler sets the operation handler for the delete user ID operation
	UserDeleteUserIDHandler user.DeleteUserIDHandler
	// VideoDeleteVideoIDHandler sets the operation handler for the delete video ID operation
	VideoDeleteVideoIDHandler video.DeleteVideoIDHandler
	// CommentGetCommentsVideoIDHandler sets the operation handler for the get comments video ID operation
	CommentGetCommentsVideoIDHandler comment.GetCommentsVideoIDHandler
	// HealthGetHealthHandler sets the operation handler for the get health operation
	HealthGetHealthHandler health.GetHealthHandler
	// UserGetUserIDHandler sets the operation handler for the get user ID operation
	UserGetUserIDHandler user.GetUserIDHandler
	// VideoGetVideoIDHandler sets the operation handler for the get video ID operation
	VideoGetVideoIDHandler video.GetVideoIDHandler
	// VideoGetVideoRandomHandler sets the operation handler for the get video random operation
	VideoGetVideoRandomHandler video.GetVideoRandomHandler
	// VideoGetVideoSearchHandler sets the operation handler for the get video search operation
	VideoGetVideoSearchHandler video.GetVideoSearchHandler
	// VideoGetVideoSubHandler sets the operation handler for the get video sub operation
	VideoGetVideoSubHandler video.GetVideoSubHandler
	// VideoGetVideoTagsHandler sets the operation handler for the get video tags operation
	VideoGetVideoTagsHandler video.GetVideoTagsHandler
	// VideoGetVideoTrendHandler sets the operation handler for the get video trend operation
	VideoGetVideoTrendHandler video.GetVideoTrendHandler
	// LikePatchDislikeVideoIDHandler sets the operation handler for the patch dislike video ID operation
	LikePatchDislikeVideoIDHandler like.PatchDislikeVideoIDHandler
	// LikePatchLikeVideoIDHandler sets the operation handler for the patch like video ID operation
	LikePatchLikeVideoIDHandler like.PatchLikeVideoIDHandler
	// SubscriptionPatchSubUserIDHandler sets the operation handler for the patch sub user ID operation
	SubscriptionPatchSubUserIDHandler subscription.PatchSubUserIDHandler
	// SubscriptionPatchUnsubUserIDHandler sets the operation handler for the patch unsub user ID operation
	SubscriptionPatchUnsubUserIDHandler subscription.PatchUnsubUserIDHandler
	// VideoPatchVideoViewIDHandler sets the operation handler for the patch video view ID operation
	VideoPatchVideoViewIDHandler video.PatchVideoViewIDHandler
	// CommentPostCommentHandler sets the operation handler for the post comment operation
	CommentPostCommentHandler comment.PostCommentHandler
	// AuthPostUserLoginHandler sets the operation handler for the post user login operation
	AuthPostUserLoginHandler auth.PostUserLoginHandler
	// UserPostUserRegisterHandler sets the operation handler for the post user register operation
	UserPostUserRegisterHandler user.PostUserRegisterHandler
	// VideoPostVideoHandler sets the operation handler for the post video operation
	VideoPostVideoHandler video.PostVideoHandler
	// UserPutUserIDHandler sets the operation handler for the put user ID operation
	UserPutUserIDHandler user.PutUserIDHandler
	// VideoPutVideoIDHandler sets the operation handler for the put video ID operation
	VideoPutVideoIDHandler video.PutVideoIDHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *YoutubeAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *YoutubeAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *YoutubeAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *YoutubeAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *YoutubeAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *YoutubeAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *YoutubeAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *YoutubeAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *YoutubeAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the YoutubeAPIAPI
func (o *YoutubeAPIAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BearerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.CommentDeleteCommentHandler == nil {
		unregistered = append(unregistered, "comment.DeleteCommentHandler")
	}
	if o.UserDeleteUserIDHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserIDHandler")
	}
	if o.VideoDeleteVideoIDHandler == nil {
		unregistered = append(unregistered, "video.DeleteVideoIDHandler")
	}
	if o.CommentGetCommentsVideoIDHandler == nil {
		unregistered = append(unregistered, "comment.GetCommentsVideoIDHandler")
	}
	if o.HealthGetHealthHandler == nil {
		unregistered = append(unregistered, "health.GetHealthHandler")
	}
	if o.UserGetUserIDHandler == nil {
		unregistered = append(unregistered, "user.GetUserIDHandler")
	}
	if o.VideoGetVideoIDHandler == nil {
		unregistered = append(unregistered, "video.GetVideoIDHandler")
	}
	if o.VideoGetVideoRandomHandler == nil {
		unregistered = append(unregistered, "video.GetVideoRandomHandler")
	}
	if o.VideoGetVideoSearchHandler == nil {
		unregistered = append(unregistered, "video.GetVideoSearchHandler")
	}
	if o.VideoGetVideoSubHandler == nil {
		unregistered = append(unregistered, "video.GetVideoSubHandler")
	}
	if o.VideoGetVideoTagsHandler == nil {
		unregistered = append(unregistered, "video.GetVideoTagsHandler")
	}
	if o.VideoGetVideoTrendHandler == nil {
		unregistered = append(unregistered, "video.GetVideoTrendHandler")
	}
	if o.LikePatchDislikeVideoIDHandler == nil {
		unregistered = append(unregistered, "like.PatchDislikeVideoIDHandler")
	}
	if o.LikePatchLikeVideoIDHandler == nil {
		unregistered = append(unregistered, "like.PatchLikeVideoIDHandler")
	}
	if o.SubscriptionPatchSubUserIDHandler == nil {
		unregistered = append(unregistered, "subscription.PatchSubUserIDHandler")
	}
	if o.SubscriptionPatchUnsubUserIDHandler == nil {
		unregistered = append(unregistered, "subscription.PatchUnsubUserIDHandler")
	}
	if o.VideoPatchVideoViewIDHandler == nil {
		unregistered = append(unregistered, "video.PatchVideoViewIDHandler")
	}
	if o.CommentPostCommentHandler == nil {
		unregistered = append(unregistered, "comment.PostCommentHandler")
	}
	if o.AuthPostUserLoginHandler == nil {
		unregistered = append(unregistered, "auth.PostUserLoginHandler")
	}
	if o.UserPostUserRegisterHandler == nil {
		unregistered = append(unregistered, "user.PostUserRegisterHandler")
	}
	if o.VideoPostVideoHandler == nil {
		unregistered = append(unregistered, "video.PostVideoHandler")
	}
	if o.UserPutUserIDHandler == nil {
		unregistered = append(unregistered, "user.PutUserIDHandler")
	}
	if o.VideoPutVideoIDHandler == nil {
		unregistered = append(unregistered, "video.PutVideoIDHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *YoutubeAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *YoutubeAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, o.BearerAuth)

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *YoutubeAPIAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *YoutubeAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *YoutubeAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *YoutubeAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the youtube API API
func (o *YoutubeAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *YoutubeAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/comment"] = comment.NewDeleteComment(o.context, o.CommentDeleteCommentHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/user/{id}"] = user.NewDeleteUserID(o.context, o.UserDeleteUserIDHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/video/{id}"] = video.NewDeleteVideoID(o.context, o.VideoDeleteVideoIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/comments/{videoID}"] = comment.NewGetCommentsVideoID(o.context, o.CommentGetCommentsVideoIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/health"] = health.NewGetHealth(o.context, o.HealthGetHealthHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/user/{id}"] = user.NewGetUserID(o.context, o.UserGetUserIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/video/{id}"] = video.NewGetVideoID(o.context, o.VideoGetVideoIDHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/video/random"] = video.NewGetVideoRandom(o.context, o.VideoGetVideoRandomHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/video/search"] = video.NewGetVideoSearch(o.context, o.VideoGetVideoSearchHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/video/sub"] = video.NewGetVideoSub(o.context, o.VideoGetVideoSubHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/video/tags"] = video.NewGetVideoTags(o.context, o.VideoGetVideoTagsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/video/trend"] = video.NewGetVideoTrend(o.context, o.VideoGetVideoTrendHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/dislike/{video_id}"] = like.NewPatchDislikeVideoID(o.context, o.LikePatchDislikeVideoIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/like/{video_id}"] = like.NewPatchLikeVideoID(o.context, o.LikePatchLikeVideoIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/sub/{user_id}"] = subscription.NewPatchSubUserID(o.context, o.SubscriptionPatchSubUserIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/unsub/{user_id}"] = subscription.NewPatchUnsubUserID(o.context, o.SubscriptionPatchUnsubUserIDHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/video/view/{id}"] = video.NewPatchVideoViewID(o.context, o.VideoPatchVideoViewIDHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/comment"] = comment.NewPostComment(o.context, o.CommentPostCommentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/login"] = auth.NewPostUserLogin(o.context, o.AuthPostUserLoginHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/user/register"] = user.NewPostUserRegister(o.context, o.UserPostUserRegisterHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/video"] = video.NewPostVideo(o.context, o.VideoPostVideoHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/user/{id}"] = user.NewPutUserID(o.context, o.UserPutUserIDHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/video/{id}"] = video.NewPutVideoID(o.context, o.VideoPutVideoIDHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *YoutubeAPIAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *YoutubeAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *YoutubeAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *YoutubeAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *YoutubeAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
