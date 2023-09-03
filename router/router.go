package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrkhay/creative-quill-backend/handlers"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

type ApiError struct {
	Error string `json:"error"`
}
type ApiSuccess struct {
	Success string `json:"success"`
}

type apiFunc func(http.ResponseWriter, *http.Request) *u.ApiError

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if err := f(w, r); err != nil {
			u.WriteJson(w, err.Status, err.Error.Error())
		}

	}
}

func SetupRoutes(app *mux.Router, h handlers.Database) {

	/// u - Update
	/// d - Delete

	//  health
	app.HandleFunc("/health", makeHttpHandleFunc(handlers.HealthCheckHandler))

	// auth
	app.HandleFunc("/account/register", makeHttpHandleFunc(h.Register)).Methods("POST")
	app.HandleFunc("/account/login", makeHttpHandleFunc(h.Login)).Methods("POST")
	app.HandleFunc("/account/altregister", makeHttpHandleFunc(h.AltRegister)).Methods("POST")
	app.HandleFunc("/account/altlogin", makeHttpHandleFunc(h.AltLogin)).Methods("POST")

	// user
	app.HandleFunc("/user/follow", makeHttpHandleFunc(h.FollowUser)).Methods("POST")
	app.HandleFunc("/user/unfollow", makeHttpHandleFunc(h.UnFollowUser)).Methods("POST")
	app.HandleFunc("/user/followers", makeHttpHandleFunc(h.GetFollowersAccDetail)).Methods("GET")
	app.HandleFunc("/user/following", makeHttpHandleFunc(h.GetFolloweringAccDetails)).Methods("GET")
	app.HandleFunc("/user/u/firstname", makeHttpHandleFunc(h.UpdateFirstName)).Methods("PUT")
	app.HandleFunc("/user/u/lastname", makeHttpHandleFunc(h.UpdateLastName)).Methods("PUT")
	app.HandleFunc("/user/u/displayname", makeHttpHandleFunc(h.UpdateDisplayName)).Methods("PUT")
	app.HandleFunc("/user/u/username", makeHttpHandleFunc(h.UpdateUsername)).Methods("PUT")
	app.HandleFunc("/user/u/website", makeHttpHandleFunc(h.UpdateWebsiteUrl)).Methods("PUT")
	app.HandleFunc("/user/u/profilepic", makeHttpHandleFunc(h.UpdateProfilePicUrl)).Methods("PUT")
	app.HandleFunc("/user/u/description", makeHttpHandleFunc(h.UpdateDescription)).Methods("PUT")
	app.HandleFunc("/user/u/headerpic", makeHttpHandleFunc(h.UpdateHeaderPicUrl)).Methods("PUT")
	app.HandleFunc("/user/u/account", makeHttpHandleFunc(h.UpdateHeaderPicUrl)).Methods("PUT")

	// article
	app.HandleFunc("/articles", makeHttpHandleFunc(h.GetArticle)).Methods("GET")
	app.HandleFunc("/new/article", makeHttpHandleFunc(h.CreateNewArticle)).Methods("POST")
	app.HandleFunc("/u/article", makeHttpHandleFunc(h.ModifieArticle)).Methods("PUT")
	app.HandleFunc("/d/article", makeHttpHandleFunc(h.DeleteArticle)).Methods("DELETE")
	app.HandleFunc("/like/article", makeHttpHandleFunc(h.LikeArticle)).Methods("POST")
	app.HandleFunc("/dislike/article", makeHttpHandleFunc(h.DisLikeArticle)).Methods("POST")
	app.HandleFunc("/article/comments", makeHttpHandleFunc(h.GetArticleComments)).Methods("GET")

	// comment
	app.HandleFunc("/new/comment", makeHttpHandleFunc(h.CreateNewArticleComment)).Methods("POST")
	app.HandleFunc("/d/comment", makeHttpHandleFunc(h.DeleteComment)).Methods("DELETE")
	app.HandleFunc("/like/comment", makeHttpHandleFunc(h.LikeComment)).Methods("POST")
	app.HandleFunc("/dislike/comment", makeHttpHandleFunc(h.DislikeComment)).Methods("POST")
	app.HandleFunc("/sub/comment", makeHttpHandleFunc(h.CreateNewArticleSubComment)).Methods("POST")
	app.HandleFunc("/sub/comment", makeHttpHandleFunc(h.GetCommentSubComments)).Methods("GET")

}
