package database

import (
	"fmt"
	"net/http"

	t "github.com/mrkhay/creative-quill-backend/models"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

func (s *PostgresStorage) GetAllArticles() ([]*t.Article, *u.ApiError) {

	query := `SELECT * FROM articles`

	row, err := s.db.Query(query)

	articles := []*t.Article{}

	for row.Next() {
		a := t.Article{}

		err := scanIntoArtice(&a, row)

		if err != nil {
			return nil, err
		}

		articles = append(articles, &a)
	}

	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}
	return articles, nil
}

func (s *PostgresStorage) CreateArticle(req *t.NewArticleRequest) *u.ApiError {

	query := `INSERT INTO articles (id,author_id, title,thumbnail_url,last_modified,date_created,content)
	values($1,$2,$3,$4,$5,$6,$7)`

	_, err := s.db.Exec(query,
		req.ID,
		req.AuthorID,
		req.Title,
		req.ThumbnailUrl,
		req.LastModified,
		req.DateCreated,
		req.Content,
	)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}
	return nil
}

func (s *PostgresStorage) ModifieArticle(req *t.ModifieArticleRequest) *u.ApiError {

	tx, err := s.db.Begin()

	if err != nil {
		tx.Rollback()
		return u.NewError(err, http.StatusExpectationFailed)
	}

	ch := make(chan *u.ApiError)

	go func() {
		res, err := tx.Query(`SELECT EXISTS (SELECT 1 FROM articles WHERE author_id = $1 and id = $2)`, req.AuthorID, req.ArticleID)

		if err != nil {
			ch <- u.NewError(err, http.StatusConflict)
		}

		isMatch := false
		for res.Next() {
			err := res.Scan(&isMatch)
			if err != nil {
				ch <- u.NewError(err, http.StatusConflict)
			}
		}

		if !isMatch {

			ch <- u.NewError(fmt.Errorf("unauthorized"), http.StatusUnauthorized)
		}

		ch <- nil
	}()

	res := <-ch

	if res != nil {
		tx.Rollback()
		return res
	}

	query := `UPDATE articles SET title = $1,thumbnail_url=$2,last_modified=$3,content=$4 WHERE id = $5 AND author_id = $6`

	_, err = tx.Exec(query,
		req.Title,
		req.ThumbnailUrl,
		req.LastModified,
		req.Content,
		req.ArticleID,
		req.AuthorID,
	)

	if err != nil {
		tx.Rollback()
		return u.NewError(err, http.StatusConflict)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return u.NewError(err, http.StatusConflict)
	}
	return nil
}

func (s *PostgresStorage) DeleteArticle(req *t.DeleteArticleRequest) *u.ApiError {

	tx, err := s.db.Begin()

	if err != nil {
		tx.Rollback()
		return u.NewError(err, http.StatusExpectationFailed)
	}

	ch := make(chan *u.ApiError)

	go func() {
		res, err := tx.Query(`SELECT EXISTS (SELECT 1 FROM articles WHERE author_id = $1 and id = $2)`, req.AuthorID, req.ArticleID)

		if err != nil {
			ch <- u.NewError(err, http.StatusConflict)
		}

		isMatch := false
		for res.Next() {
			err := res.Scan(&isMatch)
			if err != nil {
				ch <- u.NewError(err, http.StatusConflict)
			}
		}

		if !isMatch {

			ch <- u.NewError(fmt.Errorf("unauthorized"), http.StatusUnauthorized)
		}

		ch <- nil
	}()

	res := <-ch

	if res != nil {
		tx.Rollback()
		return res
	}

	query := `DELETE FROM articles WHERE id = $1 AND author_id = $2`

	_, err = tx.Exec(
		query,
		req.ArticleID,
		req.AuthorID,
	)

	if err != nil {
		tx.Rollback()
		return u.NewError(err, http.StatusConflict)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return u.NewError(err, http.StatusConflict)
	}
	return nil
}

func (s *PostgresStorage) LikeArticle(req *t.Like) *u.ApiError {

	ch := make(chan *u.ApiError)

	go func() {

		ch <- s.UnlikeArticle(req)
	}()

	res := <-ch

	if res != nil {
		return res
	}

	query := `INSERT INTO article_likes (id,user_id, article_id)
	values($1,$2,$3)`

	_, err := s.db.Exec(query,
		req.ID,
		req.UserID,
		req.ContentID,
	)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}
	return nil

}
func (s *PostgresStorage) DislikeArticle(req *t.Dislike) *u.ApiError {

	ch := make(chan *u.ApiError)

	go func() {

		ch <- s.UndislikeArticle(req)
	}()

	res := <-ch

	if res != nil {
		return res
	}

	query := `INSERT INTO article_dislikes (id,user_id, article_id)
	values($1,$2,$3)`

	_, err := s.db.Exec(query,
		req.ID,
		req.UserID,
		req.ContentID,
	)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil
}
func (s *PostgresStorage) UnlikeArticle(req *t.Like) *u.ApiError {

	_, err := s.db.Query(`DELETE FROM article_likes WHERE article_id = $1 AND user_id = $2`, req.ContentID, req.UserID)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil

}
func (s *PostgresStorage) UndislikeArticle(req *t.Dislike) *u.ApiError {

	_, err := s.db.Query(`DELETE FROM article_dislikes WHERE article_id = $1 AND user_id = $2`, req.ContentID, req.UserID)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil
}

func (s *PostgresStorage) CreateArticleNewComment(req *t.NewArticleCommentRequest) *u.ApiError {

	query := `INSERT INTO article_comments (id,user_id,article_id,date_created,content)
	values($1,$2,$3,$4,$5)`

	_, err := s.db.Exec(query,
		req.ID,
		req.UserID,
		req.ArticleID,
		req.DateCreated,
		req.Content,
	)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}
	return nil
}

func (s *PostgresStorage) GetArticleComments(req *t.ArticleCommentsRequest) ([]*t.Comment, *u.ApiError) {

	query :=
		`SELECT * FROM article_comments WHERE article_id = $1 AND parent_comment_id IS NULL`

	row, err := s.db.Query(
		query,
		req.ArticleID,
	)
	comments := []*t.Comment{}

	for row.Next() {

		comment := new(t.Comment)

		err := scanIntoComment(comment, row)

		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return comments, nil

}
