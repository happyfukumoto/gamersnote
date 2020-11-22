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

	"gamersnote.com/v1/restapi/operations/article"
	"gamersnote.com/v1/restapi/operations/comment"
	"gamersnote.com/v1/restapi/operations/follow"
	"gamersnote.com/v1/restapi/operations/like"
	"gamersnote.com/v1/restapi/operations/notification"
	"gamersnote.com/v1/restapi/operations/stock"
	"gamersnote.com/v1/restapi/operations/user"
	"gamersnote.com/v1/utils"
)

// NewGamersnoteAPIAPI creates a new GamersnoteAPI instance
func NewGamersnoteAPIAPI(spec *loads.Document) *GamersnoteAPIAPI {
	return &GamersnoteAPIAPI{
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

		ArticleDeleteArticleHandler: article.DeleteArticleHandlerFunc(func(params article.DeleteArticleParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation article.DeleteArticle has not yet been implemented")
		}),
		CommentDeleteCommentHandler: comment.DeleteCommentHandlerFunc(func(params comment.DeleteCommentParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation comment.DeleteComment has not yet been implemented")
		}),
		LikeDeleteLikeHandler: like.DeleteLikeHandlerFunc(func(params like.DeleteLikeParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation like.DeleteLike has not yet been implemented")
		}),
		StockDeleteStockHandler: stock.DeleteStockHandlerFunc(func(params stock.DeleteStockParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation stock.DeleteStock has not yet been implemented")
		}),
		UserDeleteUserHandler: user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		}),
		ArticleGetArticleHandler: article.GetArticleHandlerFunc(func(params article.GetArticleParams) middleware.Responder {
			return middleware.NotImplemented("operation article.GetArticle has not yet been implemented")
		}),
		ArticleGetArticlesHandler: article.GetArticlesHandlerFunc(func(params article.GetArticlesParams) middleware.Responder {
			return middleware.NotImplemented("operation article.GetArticles has not yet been implemented")
		}),
		CommentGetCommentsHandler: comment.GetCommentsHandlerFunc(func(params comment.GetCommentsParams) middleware.Responder {
			return middleware.NotImplemented("operation comment.GetComments has not yet been implemented")
		}),
		FollowGetFollowsHandler: follow.GetFollowsHandlerFunc(func(params follow.GetFollowsParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation follow.GetFollows has not yet been implemented")
		}),
		ArticleGetFollowsArticlesHandler: article.GetFollowsArticlesHandlerFunc(func(params article.GetFollowsArticlesParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation article.GetFollowsArticles has not yet been implemented")
		}),
		LikeGetLikedArticlesHandler: like.GetLikedArticlesHandlerFunc(func(params like.GetLikedArticlesParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation like.GetLikedArticles has not yet been implemented")
		}),
		UserGetMeHandler: user.GetMeHandlerFunc(func(params user.GetMeParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation user.GetMe has not yet been implemented")
		}),
		NotificationGetNotificationsHandler: notification.GetNotificationsHandlerFunc(func(params notification.GetNotificationsParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation notification.GetNotifications has not yet been implemented")
		}),
		StockGetStockedArticlesHandler: stock.GetStockedArticlesHandlerFunc(func(params stock.GetStockedArticlesParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation stock.GetStockedArticles has not yet been implemented")
		}),
		ArticleGetTheUsersArticlesHandler: article.GetTheUsersArticlesHandlerFunc(func(params article.GetTheUsersArticlesParams) middleware.Responder {
			return middleware.NotImplemented("operation article.GetTheUsersArticles has not yet been implemented")
		}),
		UserGetUserHandler: user.GetUserHandlerFunc(func(params user.GetUserParams) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUser has not yet been implemented")
		}),
		ArticlePostArticleHandler: article.PostArticleHandlerFunc(func(params article.PostArticleParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation article.PostArticle has not yet been implemented")
		}),
		CommentPostCommentHandler: comment.PostCommentHandlerFunc(func(params comment.PostCommentParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation comment.PostComment has not yet been implemented")
		}),
		UserPostUserHandler: user.PostUserHandlerFunc(func(params user.PostUserParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation user.PostUser has not yet been implemented")
		}),
		ArticlePutArticleHandler: article.PutArticleHandlerFunc(func(params article.PutArticleParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation article.PutArticle has not yet been implemented")
		}),
		CommentPutCommentHandler: comment.PutCommentHandlerFunc(func(params comment.PutCommentParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation comment.PutComment has not yet been implemented")
		}),
		FollowPutFollowHandler: follow.PutFollowHandlerFunc(func(params follow.PutFollowParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation follow.PutFollow has not yet been implemented")
		}),
		LikePutLikeHandler: like.PutLikeHandlerFunc(func(params like.PutLikeParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation like.PutLike has not yet been implemented")
		}),
		NotificationPutNotificationReadHandler: notification.PutNotificationReadHandlerFunc(func(params notification.PutNotificationReadParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation notification.PutNotificationRead has not yet been implemented")
		}),
		StockPutStockHandler: stock.PutStockHandlerFunc(func(params stock.PutStockParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation stock.PutStock has not yet been implemented")
		}),
		UserPutUserHandler: user.PutUserHandlerFunc(func(params user.PutUserParams, principal *utils.TokenPayload) middleware.Responder {
			return middleware.NotImplemented("operation user.PutUser has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		BearerAuth: func(token string) (*utils.TokenPayload, error) {
			return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*GamersnoteAPIAPI the gamersnote API API */
type GamersnoteAPIAPI struct {
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
	BearerAuth func(string) (*utils.TokenPayload, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ArticleDeleteArticleHandler sets the operation handler for the delete article operation
	ArticleDeleteArticleHandler article.DeleteArticleHandler
	// CommentDeleteCommentHandler sets the operation handler for the delete comment operation
	CommentDeleteCommentHandler comment.DeleteCommentHandler
	// LikeDeleteLikeHandler sets the operation handler for the delete like operation
	LikeDeleteLikeHandler like.DeleteLikeHandler
	// StockDeleteStockHandler sets the operation handler for the delete stock operation
	StockDeleteStockHandler stock.DeleteStockHandler
	// UserDeleteUserHandler sets the operation handler for the delete user operation
	UserDeleteUserHandler user.DeleteUserHandler
	// ArticleGetArticleHandler sets the operation handler for the get article operation
	ArticleGetArticleHandler article.GetArticleHandler
	// ArticleGetArticlesHandler sets the operation handler for the get articles operation
	ArticleGetArticlesHandler article.GetArticlesHandler
	// CommentGetCommentsHandler sets the operation handler for the get comments operation
	CommentGetCommentsHandler comment.GetCommentsHandler
	// FollowGetFollowsHandler sets the operation handler for the get follows operation
	FollowGetFollowsHandler follow.GetFollowsHandler
	// ArticleGetFollowsArticlesHandler sets the operation handler for the get follows articles operation
	ArticleGetFollowsArticlesHandler article.GetFollowsArticlesHandler
	// LikeGetLikedArticlesHandler sets the operation handler for the get liked articles operation
	LikeGetLikedArticlesHandler like.GetLikedArticlesHandler
	// UserGetMeHandler sets the operation handler for the get me operation
	UserGetMeHandler user.GetMeHandler
	// NotificationGetNotificationsHandler sets the operation handler for the get notifications operation
	NotificationGetNotificationsHandler notification.GetNotificationsHandler
	// StockGetStockedArticlesHandler sets the operation handler for the get stocked articles operation
	StockGetStockedArticlesHandler stock.GetStockedArticlesHandler
	// ArticleGetTheUsersArticlesHandler sets the operation handler for the get the users articles operation
	ArticleGetTheUsersArticlesHandler article.GetTheUsersArticlesHandler
	// UserGetUserHandler sets the operation handler for the get user operation
	UserGetUserHandler user.GetUserHandler
	// ArticlePostArticleHandler sets the operation handler for the post article operation
	ArticlePostArticleHandler article.PostArticleHandler
	// CommentPostCommentHandler sets the operation handler for the post comment operation
	CommentPostCommentHandler comment.PostCommentHandler
	// UserPostUserHandler sets the operation handler for the post user operation
	UserPostUserHandler user.PostUserHandler
	// ArticlePutArticleHandler sets the operation handler for the put article operation
	ArticlePutArticleHandler article.PutArticleHandler
	// CommentPutCommentHandler sets the operation handler for the put comment operation
	CommentPutCommentHandler comment.PutCommentHandler
	// FollowPutFollowHandler sets the operation handler for the put follow operation
	FollowPutFollowHandler follow.PutFollowHandler
	// LikePutLikeHandler sets the operation handler for the put like operation
	LikePutLikeHandler like.PutLikeHandler
	// NotificationPutNotificationReadHandler sets the operation handler for the put notification read operation
	NotificationPutNotificationReadHandler notification.PutNotificationReadHandler
	// StockPutStockHandler sets the operation handler for the put stock operation
	StockPutStockHandler stock.PutStockHandler
	// UserPutUserHandler sets the operation handler for the put user operation
	UserPutUserHandler user.PutUserHandler
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
func (o *GamersnoteAPIAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *GamersnoteAPIAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *GamersnoteAPIAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *GamersnoteAPIAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *GamersnoteAPIAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *GamersnoteAPIAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *GamersnoteAPIAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *GamersnoteAPIAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *GamersnoteAPIAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the GamersnoteAPIAPI
func (o *GamersnoteAPIAPI) Validate() error {
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

	if o.ArticleDeleteArticleHandler == nil {
		unregistered = append(unregistered, "article.DeleteArticleHandler")
	}
	if o.CommentDeleteCommentHandler == nil {
		unregistered = append(unregistered, "comment.DeleteCommentHandler")
	}
	if o.LikeDeleteLikeHandler == nil {
		unregistered = append(unregistered, "like.DeleteLikeHandler")
	}
	if o.StockDeleteStockHandler == nil {
		unregistered = append(unregistered, "stock.DeleteStockHandler")
	}
	if o.UserDeleteUserHandler == nil {
		unregistered = append(unregistered, "user.DeleteUserHandler")
	}
	if o.ArticleGetArticleHandler == nil {
		unregistered = append(unregistered, "article.GetArticleHandler")
	}
	if o.ArticleGetArticlesHandler == nil {
		unregistered = append(unregistered, "article.GetArticlesHandler")
	}
	if o.CommentGetCommentsHandler == nil {
		unregistered = append(unregistered, "comment.GetCommentsHandler")
	}
	if o.FollowGetFollowsHandler == nil {
		unregistered = append(unregistered, "follow.GetFollowsHandler")
	}
	if o.ArticleGetFollowsArticlesHandler == nil {
		unregistered = append(unregistered, "article.GetFollowsArticlesHandler")
	}
	if o.LikeGetLikedArticlesHandler == nil {
		unregistered = append(unregistered, "like.GetLikedArticlesHandler")
	}
	if o.UserGetMeHandler == nil {
		unregistered = append(unregistered, "user.GetMeHandler")
	}
	if o.NotificationGetNotificationsHandler == nil {
		unregistered = append(unregistered, "notification.GetNotificationsHandler")
	}
	if o.StockGetStockedArticlesHandler == nil {
		unregistered = append(unregistered, "stock.GetStockedArticlesHandler")
	}
	if o.ArticleGetTheUsersArticlesHandler == nil {
		unregistered = append(unregistered, "article.GetTheUsersArticlesHandler")
	}
	if o.UserGetUserHandler == nil {
		unregistered = append(unregistered, "user.GetUserHandler")
	}
	if o.ArticlePostArticleHandler == nil {
		unregistered = append(unregistered, "article.PostArticleHandler")
	}
	if o.CommentPostCommentHandler == nil {
		unregistered = append(unregistered, "comment.PostCommentHandler")
	}
	if o.UserPostUserHandler == nil {
		unregistered = append(unregistered, "user.PostUserHandler")
	}
	if o.ArticlePutArticleHandler == nil {
		unregistered = append(unregistered, "article.PutArticleHandler")
	}
	if o.CommentPutCommentHandler == nil {
		unregistered = append(unregistered, "comment.PutCommentHandler")
	}
	if o.FollowPutFollowHandler == nil {
		unregistered = append(unregistered, "follow.PutFollowHandler")
	}
	if o.LikePutLikeHandler == nil {
		unregistered = append(unregistered, "like.PutLikeHandler")
	}
	if o.NotificationPutNotificationReadHandler == nil {
		unregistered = append(unregistered, "notification.PutNotificationReadHandler")
	}
	if o.StockPutStockHandler == nil {
		unregistered = append(unregistered, "stock.PutStockHandler")
	}
	if o.UserPutUserHandler == nil {
		unregistered = append(unregistered, "user.PutUserHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *GamersnoteAPIAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *GamersnoteAPIAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "Bearer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.BearerAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *GamersnoteAPIAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *GamersnoteAPIAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
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
func (o *GamersnoteAPIAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
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
func (o *GamersnoteAPIAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the gamersnote API API
func (o *GamersnoteAPIAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *GamersnoteAPIAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/articles/{article_id}"] = article.NewDeleteArticle(o.context, o.ArticleDeleteArticleHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/articles/{article_id}/comments/{comment_id}"] = comment.NewDeleteComment(o.context, o.CommentDeleteCommentHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/me/likes/{article_id}"] = like.NewDeleteLike(o.context, o.LikeDeleteLikeHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/me/stocks/{article_id}"] = stock.NewDeleteStock(o.context, o.StockDeleteStockHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/users/me"] = user.NewDeleteUser(o.context, o.UserDeleteUserHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/articles/{article_id}"] = article.NewGetArticle(o.context, o.ArticleGetArticleHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/articles"] = article.NewGetArticles(o.context, o.ArticleGetArticlesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/articles/{article_id}/comments"] = comment.NewGetComments(o.context, o.CommentGetCommentsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me/follows"] = follow.NewGetFollows(o.context, o.FollowGetFollowsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me/follows/articles"] = article.NewGetFollowsArticles(o.context, o.ArticleGetFollowsArticlesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me/likes"] = like.NewGetLikedArticles(o.context, o.LikeGetLikedArticlesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me"] = user.NewGetMe(o.context, o.UserGetMeHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me/notifications"] = notification.NewGetNotifications(o.context, o.NotificationGetNotificationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/me/stocks"] = stock.NewGetStockedArticles(o.context, o.StockGetStockedArticlesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{gamersnote_id}/articles"] = article.NewGetTheUsersArticles(o.context, o.ArticleGetTheUsersArticlesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/users/{gamersnote_id}"] = user.NewGetUser(o.context, o.UserGetUserHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/articles"] = article.NewPostArticle(o.context, o.ArticlePostArticleHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/articles/{article_id}/comments"] = comment.NewPostComment(o.context, o.CommentPostCommentHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/users"] = user.NewPostUser(o.context, o.UserPostUserHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/articles/{article_id}"] = article.NewPutArticle(o.context, o.ArticlePutArticleHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/articles/{article_id}/comments/{comment_id}"] = comment.NewPutComment(o.context, o.CommentPutCommentHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/me/follows/{user_id}"] = follow.NewPutFollow(o.context, o.FollowPutFollowHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/me/likes/{article_id}"] = like.NewPutLike(o.context, o.LikePutLikeHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/me/notifications/{notification_id}/read"] = notification.NewPutNotificationRead(o.context, o.NotificationPutNotificationReadHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/me/stocks/{article_id}"] = stock.NewPutStock(o.context, o.StockPutStockHandler)
	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/users/me"] = user.NewPutUser(o.context, o.UserPutUserHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *GamersnoteAPIAPI) Serve(builder middleware.Builder) http.Handler {
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
func (o *GamersnoteAPIAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *GamersnoteAPIAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *GamersnoteAPIAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *GamersnoteAPIAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
