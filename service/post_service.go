package service

//go:generate mockgen -source=post_service.go -destination=./../mocks/mock_post_service.go -package=mocks

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"post-api/constants"
	"post-api/models/db"
	"post-api/models/response"
	"post-api/repository"
	"post-api/utils"
	"strings"

	"github.com/gola-glitch/gola-utils/golaerror"
	"github.com/gola-glitch/gola-utils/logging"
)

type PostService interface {
	PublishPost(ctx context.Context, draftUID, userId string) (string, *golaerror.Error)
	LikePost(userID string, postID string, ctx context.Context) (response.LikedByCount, *golaerror.Error)
	UnlikePost(userId string, postId string, ctx context.Context) (response.LikedByCount, *golaerror.Error)
	CommentPost(ctx context.Context, userId, postId, comment string) *golaerror.Error
	GetPost(ctx context.Context, postId, userId string) (response.Post, *golaerror.Error)
	MarkReadLater(ctx context.Context, postId, userId string) *golaerror.Error
	RemoveReadLater(ctx context.Context, postId, userId string) *golaerror.Error
}

type postService struct {
	repository      repository.PostsRepository
	draftRepository repository.DraftRepository
	validator       utils.PostValidator
	postTransaction neo4j.Session
}

func (service postService) PublishPost(ctx context.Context, draftUID string, userId string) (string, *golaerror.Error) {
	logger := logging.GetLogger(ctx).WithField("class", "PostService").WithField("method", "PublishPost")
	dbDraft, err := service.draftRepository.GetDraft(ctx, draftUID, userId)

	if err != nil {
		logger.Errorf("Error occurred while fetching draft from draft repository %v", err)
		if err.Error() == constants.NoDraftFoundCode {
			logger.Errorf("Error occurred while getting draft data, no draft found for draft id %v .%v", draftUID, err)
			return "", &constants.NoDraftFoundError
		}
		return "", constants.StoryInternalServerError(err.Error())
	}

	logger.Infof("Validating draft and Generated read time for post %v", draftUID)

	draft := db.Draft{
		DraftID:      dbDraft.DraftID,
		PostData:     dbDraft.PostData,
		PreviewImage: &dbDraft.PreviewImage,
		Tagline:      &dbDraft.Tagline,
		Interest:     dbDraft.Interest,
	}

	metaData, validationErr := service.validator.ValidateAndGetMetaData(draft, ctx)

	if validationErr != nil {
		logger.Errorf("Error occurred while validating draft of id %v .%v", draftUID, validationErr)
		return "", validationErr
	}

	if *draft.Tagline == "" {
		draft.Tagline = &metaData.Tagline
	}

	if *draft.PreviewImage == "" {
		draft.PreviewImage = &metaData.PreviewImage
	}

	url := utils.GenerateUrl(metaData.Title)

	post := db.PublishPost{
		PUID:         draftUID,
		UserID:       userId,
		PostData:     draft.PostData,
		ReadTime:     metaData.ReadTime,
		Interest:     draft.Interest,
		Title:        metaData.Title,
		Tagline:      *draft.Tagline,
		PreviewImage: *draft.PreviewImage,
		PostUrl:      url,
	}

	logger.Infof("Saving post in post repository for post id %v", draftUID)

	_, transactionErr := service.postTransaction.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		createPostErr := service.repository.CreatePost(ctx, post, transaction)
		if createPostErr != nil {
			logger.Errorf("Error occurred while publishing post in post repository %v", createPostErr)
			closeTransactionErr := transaction.Close()
			if closeTransactionErr != nil {
				logger.Errorf("Error occurred while rolling back create post transaction %v", closeTransactionErr)
				return nil, closeTransactionErr
			}
			return nil, createPostErr
		}
		updateStatusErr := service.draftRepository.UpdatePublishedStatus(ctx, draftUID, userId, transaction)
		if updateStatusErr != nil {
			closeTransactionErr := transaction.Close()
			if closeTransactionErr != nil {
				logger.Errorf("Error occurred while rolling back update published status for draft transaction %v", closeTransactionErr)
				return nil, closeTransactionErr
			}
			logger.Errorf("Error occurred while updating draft publish status for draft id %v, Error %v", draftUID, updateStatusErr)
			return nil, updateStatusErr
		}
		return nil, nil
	})

	if transactionErr != nil {
		logger.Errorf("Error occurred while publishing draft for draft id %v, Error %v", draftUID, transactionErr)
		return "", constants.StoryInternalServerError(transactionErr.Error())
	}

	logger.Infof("Successfully saved story for post id %v", draftUID)
	logger.Info("generating final url for post")

	finalPostUrl := strings.Join([]string{url, draftUID}, "-")

	return finalPostUrl, nil
}

func (service postService) LikePost(userID string, postUID string, ctx context.Context) (response.LikedByCount, *golaerror.Error) {
	logger := logging.GetLogger(ctx).WithField("class", "PostService").WithField("method", "LikePost")
	logger.Infof("Saving post data to draft repository")

	var likedByCount response.LikedByCount

	err := service.repository.LikePost(postUID, userID, ctx)
	if err != nil {
		logger.Errorf("Error occurred while Updating likedby in likes repository %v", err)
		return likedByCount, constants.StoryInternalServerError(err.Error())
	}

	likeCount, err := service.repository.GetLikesCountByPostID(ctx, postUID)
	if err != nil {
		logger.Errorf("Error occurred while Getting like id in likes repository %v", err)
		return likedByCount, constants.StoryInternalServerError(err.Error())
	}

	likedByCount.LikeCount = likeCount

	return likedByCount, nil
}

func (service postService) UnlikePost(userId string, postId string, ctx context.Context) (response.LikedByCount, *golaerror.Error) {
	logger := logging.GetLogger(ctx).WithField("class", "PostService").WithField("method", "LikePost")
	logger.Infof("Saving post data to draft repository")

	var likedByCount response.LikedByCount

	err := service.repository.UnlikePost(ctx, userId, postId)
	if err != nil {
		logger.Errorf("Error occurred while Updating likedby in likes repository %v", err)
		return likedByCount, constants.StoryInternalServerError(err.Error())
	}

	likeCount, err := service.repository.GetLikesCountByPostID(ctx, postId)
	if err != nil {
		logger.Errorf("Error occurred while Getting like id in likes repository %v", err)
		return likedByCount, constants.StoryInternalServerError(err.Error())
	}

	likedByCount.LikeCount = likeCount

	return likedByCount, nil
}

func (service postService) CommentPost(ctx context.Context, userId, postId, comment string) *golaerror.Error {
	logger := logging.GetLogger(ctx).WithField("class", "PostService").WithField("method", "CommentPost")
	logger.Infof("Commenting on post %v by user %v", postId, userId)

	err := service.repository.CommentPost(ctx, userId, comment, postId)

	if err != nil {
		logger.Errorf("Error occurred while user %v commenting on post %v", userId, postId)
		return constants.StoryInternalServerError(err.Error())
	}

	logger.Infof("Successfully commented on post %v by user %v", postId, userId)
	return nil
}

func (service postService) GetPost(ctx context.Context, postId, userId string) (response.Post, *golaerror.Error) {
	logger := logging.GetLogger(ctx).WithField("class", "PostService").WithField("method", "GetPost")
	logger.Infof("Fetching post for post id %v", postId)

	post, err := service.repository.FetchPost(ctx, postId, userId)

	if err != nil {
		logger.Errorf("Error occurred while fetching post for given post id %v, Error %v", postId, err)
		if err.Error() == constants.NoPostFoundCode {
			logger.Errorf("Error No post found for given post id %v", postId)
			return response.Post{}, &constants.PostNotFoundErr
		}
		return response.Post{}, constants.StoryInternalServerError(err.Error())
	}

	logger.Infof("Successfully fetching post from post repository for given post id %v", postId)

	return post, nil
}

func (service postService) MarkReadLater(ctx context.Context, postId, userId string) *golaerror.Error {
	logger := logging.GetLogger(ctx).WithField("class", "PostService").WithField("method", "MarkReadLater")
	logger.Infof("Calling repository to update the read later status for post %v", postId)

	err := service.repository.MarkPostAsReadLater(ctx, postId, userId)
	if err != nil {
		logger.Errorf("Error occurred while updating read later status for post %v, Error %v", postId, err)
		if err.Error() == constants.NoPostFoundCode {
			logger.Errorf("Error post not found for requested post id %v, Error %v", postId, err)
			return &constants.PostNotFoundErr
		}
		return constants.StoryInternalServerError(err.Error())
	}

	logger.Infof("Successfully added post to read later for post id %v", postId)
	return nil
}

func (service postService) RemoveReadLater(ctx context.Context, postId, userId string) *golaerror.Error {
	logger := logging.GetLogger(ctx).WithField("class", "PostService").WithField("method", "RemoveReadLater")
	logger.Infof("Calling repository to remove post %v from read later", postId)

	err := service.repository.RemoveReadLater(ctx, postId, userId)
	if err != nil {
		logger.Errorf("Error occurred while removing post %v from read later for user %v, Error %v", postId, userId, err)
		if err.Error() == constants.PostNotInReadLaterCode {
			logger.Errorf("either post %v not found or post not found in read later for user %v", postId, userId)
			return &constants.PostNotInReadLaterErr
		}
		logger.Error("Error Generic error occurred on repository")
		return &constants.PostServiceFailureError
	}
	logger.Infof("Successfully removed post %v from read later for user %v", postId, userId)

	return nil
}

func NewPostService(postsRepository repository.PostsRepository, draftRepository repository.DraftRepository, validator utils.PostValidator, session neo4j.Session) PostService {
	return postService{
		repository:      postsRepository,
		draftRepository: draftRepository,
		validator:       validator,
		postTransaction: session,
	}
}
