package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	t "github.com/mrkhay/creative-quill-backend/models"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

type Storage interface {
	CreateTables() *u.ApiError
	Atuh
	Article
	User
}

type Atuh interface {
	CreateAccount(acc *t.NewAccount) *u.ApiError
	GetAccount(req *t.SigninRequest) (*t.Account, *u.ApiError)
	AltCreateAccount(acc *t.NewAccount) *u.ApiError
	AltGetAccount(req *t.AltSigninRequest) (*t.Account, *u.ApiError)
}

type Article interface {
	CreateArticle(req *t.NewArticleRequest) *u.ApiError
	GetAllArticles() ([]*t.Article, *u.ApiError)
	ModifieArticle(req *t.ModifieArticleRequest) *u.ApiError
	DeleteArticle(req *t.DeleteArticleRequest) *u.ApiError
	LikeArticle(req *t.Like) *u.ApiError
	DislikeArticle(req *t.Dislike) *u.ApiError
	UnlikeArticle(req *t.Like) *u.ApiError
	UndislikeArticle(req *t.Dislike) *u.ApiError
}

type User interface {
	FollowUser(req *t.FollowRequest) *u.ApiError
	UnFollowUser(req *t.UnFollowRequest) *u.ApiError
	GetFollowerDetails(req *t.GetFollowersDetail) ([]*t.Account, *u.ApiError)
	GetAccFollowingDetails(req *t.GetAccFollowingDetails) ([]*t.Account, *u.ApiError)
	UpdateAccount(req *t.UpdateAccountReq) *u.ApiError
	UpdateFirstName(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateLastName(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateDisplayName(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateUsername(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateWebsiteUrl(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateProfilePicUrl(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateDescription(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateHeaderPicUrl(req *t.UpdateAccDetailsRequest) *u.ApiError
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, *u.ApiError) {

	connStr := os.Getenv("POSTGRES_URI")

	if connStr == "" {
		return nil, u.NewError(fmt.Errorf("you must set your 'POSTGRES_URI' environmental variable"), http.StatusBadGateway)
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, u.NewError(err, http.StatusBadGateway)
	}

	if err := db.Ping(); err != nil {
		return nil, u.NewError(err, http.StatusBadGateway)
	}

	return &PostgresStorage{
		db: db,
	}, nil

}

func (s *PostgresStorage) Init() *u.ApiError {

	return s.CreateTables()

}

// Creats Table If table does not exists.
func (s *PostgresStorage) CreateTables() *u.ApiError {

	querey := `CREATE TABLE IF NOT EXISTS accounts (
	id uuid primary key,
	first_name varchar(50),
	last_name varchar(50),
	display_name varchar(50) DEFAULT '',
	user_name varchar(20) DEFAULT '',
	email varchar(30),
	website_url text DEFAULT '',
	description text DEFAULT '',
	profilepic_url text DEFAULT '',
	headerpic_url text DEFAULT '',
	followers int DEFAULT 0,
	following int DEFAULT 0,
	created_at timestamp,
    password text
    )`

	_, err := s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}
	querey = `CREATE TABLE IF NOT EXISTS articles (
		id uuid PRIMARY KEY,
		author_id uuid REFERENCES accounts(id) ON DELETE CASCADE,
		title VARCHAR(70),
		thumbnail_url text,
		likes int DEFAULT 0,
		dislikes int DEFAULT 0,
		comments int DEFAULT 0,
		last_modified time,
		date_created time,
		content json)`

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}
	querey = `CREATE TABLE IF NOT EXISTS article_likes (
		ID uuid PRIMARY KEY,
		user_id uuid REFERENCES accounts(id) ON DELETE CASCADE,
		article_id uuid REFERENCES articles(id) ON DELETE CASCADE)`

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}
	querey = `CREATE TABLE IF NOT EXISTS article_dislikes (
		ID uuid PRIMARY KEY,
		user_id uuid REFERENCES accounts(id) ON DELETE CASCADE,
		article_id uuid REFERENCES articles(id) ON DELETE CASCADE
		)`

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	querey = `CREATE TABLE IF NOT EXISTS following (
    follower_id uuid References accounts(id) ON DELETE CASCADE,
    following_id uuid References accounts(id) ON DELETE CASCADE,
	date_followed time,
	PRIMARY KEY (follower_id,following_id)
    )`

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	querey = `CREATE  OR REPLACE FUNCTION update_following_whenfollowed_func()
	RETURNS TRIGGER AS $$
	BEGIN
    UPDATE accounts
	SET followers = followers + 1
	WHERE accounts.id = NEW.following_id;

	RETURN NEW;
	
	END;
	$$ LANGUAGE plpgsql;
	
	
	CREATE OR REPLACE TRIGGER update_following_whenfollowed_func_trigger
	AFTER INSERT ON following
	FOR EACH ROW
	EXECUTE FUNCTION update_following_whenfollowed_func();


    CREATE  OR REPLACE FUNCTION update_followers_whenfollowed_func()
	RETURNS TRIGGER AS $$
	BEGIN
    UPDATE accounts
	SET following = following + 1
	WHERE accounts.id = NEW.follower_id;

	RETURN NEW;
	
	END;
	$$ LANGUAGE plpgsql;
	
	
	CREATE OR REPLACE TRIGGER  update_followers_whenfollowed_func_trigger
	AFTER INSERT ON following
	FOR EACH ROW
	EXECUTE FUNCTION  update_followers_whenfollowed_func();
			   `

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	querey = `CREATE  OR REPLACE FUNCTION update_following_whenunfollowed_func()
	RETURNS TRIGGER AS $$
	BEGIN
    UPDATE accounts
	SET followers = followers - 1
	WHERE accounts.id = OLD.following_id;

	RETURN OLD;
	
	END;
	$$ LANGUAGE plpgsql;
	
	
	CREATE OR REPLACE TRIGGER update_following_whenunfollowed_func_trigger
	AFTER DELETE ON following
	FOR EACH ROW
	EXECUTE FUNCTION update_following_whenunfollowed_func();


    CREATE  OR REPLACE FUNCTION update_followers_whenunfollowed_func()
	RETURNS TRIGGER AS $$
	BEGIN
    UPDATE accounts
	SET following = following - 1
	WHERE accounts.id = OLD.follower_id;

	RETURN OLD;
	
	END;
	$$ LANGUAGE plpgsql;
	
	
	CREATE OR REPLACE TRIGGER  update_followers_whenunfollowed_func_trigger
	AFTER DELETE ON following
	FOR EACH ROW
	EXECUTE FUNCTION  update_followers_whenunfollowed_func();`

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	querey = `CREATE OR REPLACE FUNCTION article_like_func()
               RETURNS TRIGGER AS $$
               BEGIN
               UPDATE articles
               SET likes = likes +1
               WHERE articles.id = New.article_id;
               RETURN NEW;
               END;
               $$ LANGUAGE plpgsql;
               
               CREATE OR REPLACE TRIGGER article_like_func_trigger
               AFTER INSERT ON article_likes
               FOR EACH ROW
               EXECUTE FUNCTION article_like_func();
			   
			   CREATE OR REPLACE FUNCTION article_unlike_func()
               RETURNS TRIGGER AS $$
               BEGIN
               UPDATE articles
               SET likes = likes - 1
               WHERE articles.id = OLD.article_id;
               RETURN OLD;
               END;
               $$ LANGUAGE plpgsql;
               
               CREATE OR REPLACE TRIGGER article_unlike_func_trigger
               AFTER DELETE ON article_likes
               FOR EACH ROW
               EXECUTE FUNCTION article_unlike_func();

			   `

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	querey = `CREATE OR REPLACE FUNCTION article_dislike_func()
               RETURNS TRIGGER AS $$
               BEGIN
               UPDATE articles
               SET dislikes = dislikes +1
               WHERE articles.id = New.article_id;
               RETURN NEW;
               END;
               $$ LANGUAGE plpgsql;
               
               CREATE OR REPLACE TRIGGER article_dislike_func_trigger
               AFTER INSERT ON article_dislikes
               FOR EACH ROW
               EXECUTE FUNCTION article_dislike_func();
			   

			   CREATE OR REPLACE FUNCTION article_undislike_func()
               RETURNS TRIGGER AS $$
               BEGIN
               UPDATE articles
               SET dislikes = dislikes -1
               WHERE articles.id = OLD.article_id;
               RETURN OLD;
               END;
               $$ LANGUAGE plpgsql;
               
               CREATE OR REPLACE TRIGGER article_undislike_func_trigger
               AFTER DELETE ON article_dislikes
               FOR EACH ROW
               EXECUTE FUNCTION article_undislike_func();
			   `

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil
}

func scanIntoAccount(acc *t.Account, rows *sql.Rows) *u.ApiError {

	err := rows.Scan(
		&acc.ID,
		&acc.Firstname,
		&acc.Lastname,
		&acc.Displayname,
		&acc.Username,
		&acc.Email,
		&acc.WebsiteUrl,
		&acc.Description,
		&acc.ProfilePicUrl,
		&acc.HeaderPicUrl,
		&acc.AccFolloweringCount,
		&acc.FollowersCount,
		&acc.CreatedAt,
		&acc.Password,
	)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil

}

func scanIntoArtice(acc *t.Article, rows *sql.Rows) *u.ApiError {

	err := rows.Scan(
		&acc.ID,
		&acc.AuthorID,
		&acc.Title,
		&acc.ThumbnailUrl,
		&acc.Likes,
		&acc.Dislikes,
		&acc.Comments,
		&acc.DateCreated,
		&acc.LastModified,
		&acc.Content,
	)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil

}
