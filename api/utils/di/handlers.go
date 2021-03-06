package di

import (
	"gamersnote.com/v1/handlers/article"
	"gamersnote.com/v1/handlers/comment"
	"gamersnote.com/v1/handlers/image"
	"gamersnote.com/v1/handlers/like"
	"gamersnote.com/v1/handlers/user"
)

// article
var GetArticlesHandler *article.GetArticlesHandler
var GetArticleByArticleIDHandler *article.GetArticleByArticleIDHandler
var PostArticleHandler *article.PostArticleHandler
var GetArticleByUsernameHandler *article.GetArticleByUsernameHandler
var GetLikedArticleHandler *article.GetArticleByUserLikedHandler

// image
var UploadImageHandler *image.UploadImageHandler

// user
var DeleteMeHandler *user.DeleteMeHandler
var GetMeHandler *user.GetMeHandler
var GetUserHandler *user.GetUserHandler
var PatchSigninedHandler *user.PatchSigninedHandler
var PatchVerifiedHandler *user.PatchVefifiedHandler
var PostUserHandler *user.PostUserHandler
var PutMeHandler *user.PutMeHandler
var PatchSignoutedHandler *user.PatchSignoutedHandler
var PatchEmailVefifiedHandler *user.PatchEmailVefifiedHandler
var PatchPasswordResetHandler *user.PatchPasswordResetHandler
var PutPasswordHandler *user.PutPasswordHandler
var DeleteArticleHandler *article.DeleteArticleHandler
var PutArticleHandler *article.PutArticleHandler

// like
var PutLikeHandler *like.PutLikeHandler
var DeleteLikeHandler *like.DeleteLikeHandler
var GetLikeHandler *like.GetLikeHandler

// comment
var GetCommentsHandler *comment.GetCommentsHandler
var PostCommentHandler *comment.PostCommentHandler
var PutCommentHandler *comment.PutCommentHandler
var DeleteCommentHandler *comment.DeleteCommentHandler

func initHdlr() {
	GetArticlesHandler = article.NewGetArticlesHandler(ArticleRepository)
	GetArticleByArticleIDHandler = article.NewGetArticleByArticleIDHandler(ArticleRepository)
	PostArticleHandler = article.NewPostArticleHandler(ArticleRepository, CtxuidService)
	GetArticleByUsernameHandler = article.NewGetArticleByUsernameHandler(ArticleRepository, UserRepository)
	DeleteArticleHandler = article.NewDeleteArticleHandler(ArticleRepository, CtxuidService)
	UploadImageHandler = image.NewUploadImageHandler(ImageService)
	DeleteMeHandler = user.NewDeleteMeHandler(UserRepository, CtxuidService)
	GetMeHandler = user.NewGetMeHandler(UserRepository, CtxuidService)
	GetUserHandler = user.NewGetUserHandler(UserRepository)
	PatchSigninedHandler = user.NewPatchSigninedHandler(UserRepository, SessionService)
	PatchVerifiedHandler = user.NewPatchVefifiedHandler(UserRepository, TmpuserService, SessionService)
	PostUserHandler = user.NewPostUserHandler(UserRepository, TmpuserService, EmailSender)
	PutMeHandler = user.NewPutMeHandler(UserRepository, CtxuidService, TmpemailService, EmailSender)
	PatchSignoutedHandler = user.NewPatchSignoutedHandler(CtxuidService, SessionService)
	PatchEmailVefifiedHandler = user.NewPatchEmailVefifiedHandler(UserRepository, TmpemailService)
	PatchPasswordResetHandler = user.NewPatchPasswordResetHandler(UserRepository, TmpsessionService, EmailSender)
	PutPasswordHandler = user.NewPutPasswordHandler(UserRepository, CtxuidService, TmpsessionService)
	PutArticleHandler = article.NewPutArticleHandler(ArticleRepository, CtxuidService)
	PutLikeHandler = like.NewPutLikeHandler(LikeRepository, CtxuidService)
	DeleteLikeHandler = like.NewDeleteLikeHandler(LikeRepository, CtxuidService)
	GetLikedArticleHandler = article.NewGetArticleByUserLikedHandler(ArticleRepository, CtxuidService)
	GetLikeHandler = like.NewGetLikeHandler(LikeRepository, CtxuidService)
	GetCommentsHandler = comment.NewGetCommentsHandler(CommentRepository)
	PostCommentHandler = comment.NewPostCommentHandler(CommentRepository, CtxuidService)
	PutCommentHandler = comment.NewPutCommentHandler(CommentRepository, CtxuidService)
	DeleteCommentHandler = comment.NewDeleteCommentHandler(CommentRepository, CtxuidService)
}
