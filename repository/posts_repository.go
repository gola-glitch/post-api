package repository

//go:generate mockgen -source=posts_repository.go -destination=./../mocks/mock_posts_repository.go -package=mocks

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"post-api/models/db"

	"github.com/gola-glitch/gola-utils/logging"
	"github.com/jmoiron/sqlx"
)

type PostsRepository interface {
	CreatePost(ctx context.Context, post db.PublishPost) (int64, error)
	GetLikeCountByPost(ctx context.Context, postID string) (string, error)
	AppendOrRemoveUserFromLikedBy(postID string, userID string, ctx context.Context) error
	SaveInitialLike(ctx context.Context, postID int64) error
}

type postRepository struct {
	db *sqlx.DB
}

const (
	PublishPost           = "INSERT INTO POSTS (puid, user_id, post_data, read_time, view_count) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	GetLikedByCount       = "SELECT array_length(liked_by, 1) FROM likes WHERE post_id=$1 "
	UpdateOrRemoveLikedBy = "UPDATE likes SET liked_by = case when (SELECT count(id) as id  FROM likes WHERE post_id=$1 AND $2=ANY(liked_by)) = '1' then array_remove(liked_by, $2) else array_append(liked_by, $2) end where post_id = $1"
	InitialLike           = "INSERT INTO likes (liked_by, post_id) VALUES ('{}', $1)"
)

func (repository postRepository) CreatePost(ctx context.Context, post db.PublishPost) (int64, error) {
	logger := logging.GetLogger(ctx).WithField("class", "PostsRepository").WithField("method", "CreatePost")

	logger.Infof("Publishing the post in posts table for post postID %v", post.PUID)

	var postID int64
	err := repository.db.QueryRowContext(ctx, PublishPost, post.PUID, post.UserID, post.PostData, post.ReadTime, post.ViewCount).Scan(&postID)

	if err != nil {
		logger.Errorf("Error occurred while publishing user post in posts table %v", err)
		return 0, err
	}

	logger.Infof("Successfully posted the post in posts table for post postID %v", post.PUID)

	return postID, nil
}

func (repository postRepository) GetLikeCountByPost(ctx context.Context, postID string) (string, error) {
	logger := logging.GetLogger(ctx).WithField("class", "PostsRepository").WithField("method", "GetLikeCountByPost")

	logger.Infof("Fetching likedby count from likes table for the given post id %v", postID)

	var likeCount sql.NullString

	err := repository.db.GetContext(ctx, &likeCount, GetLikedByCount, postID)

	if err != nil && err.Error() != sql.ErrNoRows.Error() {
		logger.Errorf("Error occurred while fetching likedby count from likes table %v", err.Error())
		return likeCount.String, err
	}

	logger.Infof("Successfully fetching likedby count from likes table for given post id %v", postID)

	return likeCount.String, nil
}

func (repository postRepository) AppendOrRemoveUserFromLikedBy(postID string, userID string, ctx context.Context) error {
	logger := logging.GetLogger(ctx)
	log := logger.WithField("class", "PostsRepository").WithField("method", "AppendOrRemoveUserFromLikedBy")

	log.Infof("updating the likedby in like table for post %v", postID)

	_, err := repository.db.ExecContext(ctx, UpdateOrRemoveLikedBy, postID, userID)

	if err != nil {
		log.Errorf("Error occurred while updating liked by in likes table %v", err)
		return err
	}

	return nil
}

func (repository postRepository) SaveInitialLike(ctx context.Context, postID int64) error {
	logger := logging.GetLogger(ctx).WithField("class", "PostRepository").WithField("method", "SaveInitialLike")

	logger.Infof("saving initial empty like for post id %v", postID)
	result, err := repository.db.ExecContext(ctx, InitialLike, postID)

	if err != nil {
		logger.Errorf("Error occurred while saving initial like for post id %v", postID)
		return err
	}

	affected, err := result.RowsAffected()

	if err != nil {
		logger.Errorf("Error occurred while getting the affected rows for post initial like insert %v", postID)
		return err
	}

	if affected != 1 {
		return errors.New(fmt.Sprintf("more than one row or nor row affected for post id %v ,affected rows %v", postID, affected))
	}

	logger.Infof("One row affected for inserting initial like for post id %v", postID)

	return nil
}

func NewPostsRepository(db *sqlx.DB) PostsRepository {
	return postRepository{db: db}
}
