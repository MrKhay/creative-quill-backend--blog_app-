package database

import (
	"fmt"
	"net/http"

	t "github.com/mrkhay/creative-quill-backend/models"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

func (s *PostgresStorage) CreateArticleSubComment(req *t.NewArticleSubCommentRequest) *u.ApiError {

	query := `INSERT INTO article_comments (id,user_id,article_id,parent_comment_id,date_created,content)
	values($1,$2,$3,$4,$5,$6)`

	_, err := s.db.Exec(
		query,
		req.ID,
		req.UserID,
		req.ArticleID,
		req.ParentCommentID,
		req.DateCreated,
		req.Content,
	)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}
	return nil
}

func (s *PostgresStorage) LikeComment(req *t.NewCommentLikeRequest) *u.ApiError {

	ch := make(chan *u.ApiError)

	go func() {

		ch <- s.UndislikeComment(&t.NewCommentDislikeRequest{UserID: req.UserID, CommentID: req.CommentID})
		fmt.Println("Called")
	}()

	res := <-ch

	if res != nil {
		return res
	}

	query := `INSERT INTO article_comment_likes (user_id,comment_id)
	values($1,$2)`

	_, err := s.db.Exec(query,
		req.UserID,
		req.CommentID,
	)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}
	return nil
}

func (s *PostgresStorage) DislikeComment(req *t.NewCommentDislikeRequest) *u.ApiError {

	ch := make(chan *u.ApiError)

	go func() {

		ch <- s.UnlikeComment(&t.NewCommentLikeRequest{UserID: req.UserID, CommentID: req.CommentID})

	}()

	res := <-ch

	if res != nil {
		return res
	}

	query := `INSERT INTO article_comment_dislikes (user_id,comment_id)
	values($1,$2)`

	_, err := s.db.Exec(query,
		req.UserID,
		req.CommentID,
	)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}
	return nil
}

func (s *PostgresStorage) UnlikeComment(req *t.NewCommentLikeRequest) *u.ApiError {

	_, err := s.db.Query(`DELETE FROM article_comment_likes WHERE user_id = $1 AND comment_id = $2`, req.UserID, req.CommentID)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil

}
func (s *PostgresStorage) UndislikeComment(req *t.NewCommentDislikeRequest) *u.ApiError {

	_, err := s.db.Query(`DELETE FROM article_comment_dislikes WHERE user_id = $1 AND comment_id = $2`, req.UserID, req.CommentID)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil
}
func (s *PostgresStorage) DeleteComment(req *t.DeleteCommentRequest) *u.ApiError {

	tx, err := s.db.Begin()

	if err != nil {
		tx.Rollback()
		return u.NewError(err, http.StatusExpectationFailed)
	}

	ch := make(chan *u.ApiError)

	go func() {
		res, err := tx.Query(`SELECT EXISTS (SELECT 1 FROM article_comments WHERE id = $1 and user_id = $2)`, req.ID, req.UserID)

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

	query := `DELETE FROM article_comments WHERE id = $1 AND user_id = $2`

	_, err = tx.Exec(
		query,
		req.ID,
		req.UserID,
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

func (s *PostgresStorage) GetCommentSubComments(req *t.GetCommentSubCommentsRequest) ([]*t.Comment, *u.ApiError) {

	query :=
		`SELECT * FROM article_comments WHERE article_id = $1 AND parent_comment_id = $2`

	row, err := s.db.Query(
		query,
		req.ArticleID,
		req.ParentCommentID,
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
