package controller

import (
	"github.com/google/uuid"
	"net/http"
	"post-api/story/constants"
	"post-api/story/models"
	"post-api/story/models/request"
	"post-api/story/service"
	"post-api/story/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gola-glitch/gola-utils/logging"
)

type DraftController struct {
	service service.DraftService
}

func (controller DraftController) CreateDraft(ctx *gin.Context) {
	logger := logging.GetLogger(ctx).WithField("class", "DraftController").WithField("method", "CreateDraft")

	token, err := utils.GetIDToken(ctx)
	if err != nil {
		logger.Error("id token not found", err)
		ctx.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userUUID, _ := uuid.Parse(token.UserId)
	logger.Infof("Entered controller to upsert draft request for user %v", userUUID)

	var upsertPost models.CreateDraft
	err = ctx.ShouldBindBodyWith(&upsertPost, binding.JSON)
	if err != nil {
		logger.Errorf("Unable to bind upsert draft request for user %v. Error %v", userUUID, err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}
	upsertPost.UserID = userUUID
	draftUUID, draftSaveErr := controller.service.CreateDraft(ctx, upsertPost)
	if draftSaveErr != nil {
		logger.Errorf("Error occurred in draft service while saving draft for user %v. Error %v", userUUID, draftSaveErr)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	logger.Infof("writing response to draft request for user %v", userUUID)
	ctx.JSON(http.StatusOK, gin.H{
		"draft_id": draftUUID.String(),
	})
}

// SaveDraft godoc
// @Tags draft
// @Summary SaveDraft
// @Description Save new draft or update existing draft
// @Accept json
// @Param request body models.UpsertDraft true "Request Body"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/post/v1/draft/upsertDraft [post]
func (controller DraftController) SaveDraft(ctx *gin.Context) {
	logger := logging.GetLogger(ctx).WithField("class", "DraftController").WithField("method", "SavePostDraft")

	token, err := utils.GetIDToken(ctx)
	if err != nil {
		logger.Error("id token not found", err)
		ctx.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userUUID, _ := uuid.Parse(token.UserId)
	logger.Infof("Entered controller to upsert draft request for user %v", userUUID)

	draftUUID := ctx.Query("draft")
	draftID, err := uuid.Parse(draftUUID)
	if err != nil {
		logger.Errorf("invalid draft id request for user %v. Error %v", userUUID, err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}
	var upsertPost models.UpsertDraft
	err = ctx.ShouldBindBodyWith(&upsertPost, binding.JSON)
	if err != nil {
		logger.Errorf("Unable to bind upsert draft request for user %v. Error %v", userUUID, err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}

	logger.Infof("Request body bind successful with upsert draft request for user %v", userUUID)
	upsertPost.UserID = userUUID
	upsertPost.DraftID = draftID
	draftSaveErr := controller.service.UpdateDraft(upsertPost, ctx)
	if draftSaveErr != nil {
		logger.Errorf("Error occurred in draft service while saving draft for user %v. Error %v", userUUID, draftSaveErr)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	logger.Infof("writing response to draft request for user %v", userUUID)
	ctx.Status(http.StatusOK)
}

func (controller DraftController) SaveTagline(ctx *gin.Context) {
	logger := logging.GetLogger(ctx).WithField("class", "DraftController").WithField("method", "SaveTagline")

	token, err := utils.GetIDToken(ctx)
	if err != nil {
		logger.Error("id token not found", err)
		ctx.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userUUID, _ := uuid.Parse(token.UserId)
	logger.Infof("Entered controller to save tagline request for user %v", userUUID)

	draftUUID := ctx.Query("draft")
	draftID, err := uuid.Parse(draftUUID)
	if err != nil {
		logger.Errorf("invalid draft id request for user %v. Error %v", userUUID, err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}
	var upsertTagline request.TaglineSaveRequest
	err = ctx.ShouldBindBodyWith(&upsertTagline, binding.JSON)
	if err != nil {
		logger.Errorf("Unable to bind upsert draft request for user %v. Error %v", "12", err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}

	logger.Infof("Request body bind successful with save tagline request for user %v", userUUID)
	upsertTagline.DraftID = draftID
	upsertTagline.UserID = userUUID
	draftSaveErr := controller.service.UpsertTagline(upsertTagline, ctx)

	if draftSaveErr != nil {
		logger.Errorf("Error occurred in draft service while saving tagline for user %v. Error %v", "12", draftSaveErr)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	logger.Infof("writing response to tagline request for user %v", userUUID)

	ctx.Status(http.StatusOK)
}

func (controller DraftController) SaveInterests(ctx *gin.Context) {
	logger := logging.GetLogger(ctx)
	log := logger.WithField("class", "DraftController").WithField("method", "SaveInterests")

	token, err := utils.GetIDToken(ctx)
	if err != nil {
		logger.Error("id token not found", err)
		ctx.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userUUID, _ := uuid.Parse(token.UserId)
	log.Infof("Entered controller to save Interests request for user %v", userUUID)

	draftUUID := ctx.Query("draft")
	draftID, err := uuid.Parse(draftUUID)
	if err != nil {
		log.Errorf("invalid draft id request for user %v. Error %v", userUUID, err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}

	var upsertInterests request.InterestsSaveRequest
	err = ctx.ShouldBindBodyWith(&upsertInterests, binding.JSON)
	if err != nil {
		log.Errorf("Unable to bind upsert interests request for user %v. Error %v", "12", err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}

	log.Infof("Request body bind successful with save interests request for user %v", userUUID)
	upsertInterests.DraftID = draftID
	upsertInterests.UserID = userUUID
	draftSaveErr := controller.service.UpsertInterests(upsertInterests, ctx)

	if draftSaveErr != nil {
		log.Errorf("Error occurred in draft service while saving interests for user %v. Error %v", "12", draftSaveErr)
		constants.RespondWithGolaError(ctx, draftSaveErr)
		return
	}

	log.Infof("writing response to interests request for user %v", userUUID)

	ctx.Status(http.StatusOK)
}

func (controller DraftController) GetDraft(ctx *gin.Context) {
	logger := logging.GetLogger(ctx).WithField("class", "DraftController").WithField("method", "GetDraft")
	token, err := utils.GetIDToken(ctx)
	if err != nil {
		logger.Error("id token not found", err)
		ctx.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userUUID, _ := uuid.Parse(token.UserId)
	logger.Infof("Entered controller to get draft request for user %v", userUUID)

	draftUUID := ctx.Query("draft")
	draftID, err := uuid.Parse(draftUUID)
	if err != nil {
		logger.Errorf("invalid draft id request for user %v. Error %v", userUUID, err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}
	logger.Infof("Entered controller to get draft request for user %v", userUUID)

	draftData, draftGetErr := controller.service.GetDraft(ctx, draftID, userUUID)
	if draftGetErr != nil {
		logger.Errorf("Error occurred in draft service while saving tagline for user %v. Error %v", userUUID, draftGetErr)
		constants.RespondWithGolaError(ctx, draftGetErr)
		return
	}
	logger.Infof("writing response to draft data request for user %v %s", userUUID, draftID)
	ctx.JSON(http.StatusOK, draftData)
}

func (controller DraftController) SavePreviewImage(ctx *gin.Context) {
	logger := logging.GetLogger(ctx).WithField("class", "DraftController").WithField("method", "SavePreviewImage")
	token, err := utils.GetIDToken(ctx)
	if err != nil {
		logger.Error("id token not found", err)
		ctx.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userUUID, _ := uuid.Parse(token.UserId)
	logger.Infof("Entered controller to update preview image request for user %v", userUUID)

	draftUUID := ctx.Query("draft")
	draftID, err := uuid.Parse(draftUUID)
	if err != nil {
		logger.Errorf("invalid draft id request for user %v. Error %v", userUUID, err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}
	logger.Infof("Entered controller to save preview image for user %v", "12")

	var imageSaveRequest request.PreviewImageSaveRequest
	if bindingErr := ctx.ShouldBindBodyWith(&imageSaveRequest, binding.JSON); bindingErr != nil {
		logger.Errorf("Unable to bind request body with image save request for draft %v", bindingErr)
		ctx.JSON(http.StatusBadRequest, &constants.PayloadValidationError)
		return
	}

	logger.Infof("Request body bind successful with image save request for user %v", "12")
	imageSaveRequest.DraftID = draftID
	imageSaveRequest.UserID = userUUID
	imageSaveErr := controller.service.SavePreviewImage(ctx, imageSaveRequest)
	if imageSaveErr != nil {
		logger.Errorf("Error occurred in draft service while saving image for user %v. Error %v", "12", imageSaveErr)
		constants.RespondWithGolaError(ctx, imageSaveErr)
		return
	}

	logger.Infof("writing response to draft image save request for user %v %s", "12", imageSaveRequest.DraftID)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (controller DraftController) GetAllDraft(ctx *gin.Context) {
	logger := logging.GetLogger(ctx)
	log := logger.WithField("class", "DraftController").WithField("method", "GetAllDraft")
	token, err := utils.GetIDToken(ctx)
	if err != nil {
		logger.Error("id token not found", err)
		ctx.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userUUID, _ := uuid.Parse(token.UserId)
	logger.Infof("Entered controller to get drafts request for user %v", userUUID)

	log.Infof("Entered controller to get all draft request for user %v", userUUID)
	var draftRequest models.GetAllDraftRequest
	err = ctx.ShouldBindBodyWith(&draftRequest, binding.JSON)
	if err != nil {
		log.Errorf("Unable to bind all draft request for user %v. Error %v", userUUID, err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}

	log.Infof("Request body bind successful with get all draft request for user %v", userUUID)
	draftRequest.UserID = userUUID
	drafts, draftSaveErr := controller.service.GetAllDraft(ctx, draftRequest)
	if draftSaveErr != nil {
		log.Errorf("Error occurred in draft service while saving tagline for user %v. Error %v", userUUID, draftSaveErr)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.Infof("writing response to draft all data request for user %v", userUUID)

	ctx.JSON(http.StatusOK, drafts)
}

func (controller DraftController) DeleteDraft(ctx *gin.Context) {
	logger := logging.GetLogger(ctx).WithField("class", "DraftController").WithField("method", "DeleteDraft")
	logger.Info("Entering the controller layer to delete draft")

	token, err := utils.GetIDToken(ctx)
	if err != nil {
		logger.Error("id token not found", err)
		ctx.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userUUID, _ := uuid.Parse(token.UserId)
	logger.Infof("Entered controller to get draft request for user %v", userUUID)

	draftUUID := ctx.Query("draft")
	draftID, err := uuid.Parse(draftUUID)
	if err != nil {
		logger.Errorf("invalid draft id request for user %v. Error %v", userUUID, err)
		ctx.JSON(http.StatusBadRequest, constants.PayloadValidationError)
		return
	}

	logger.Infof("Successfully bind request uri with draft delete request for draft id %v", draftID)
	deleteDraftErr := controller.service.DeleteDraft(ctx, draftID, userUUID)

	if deleteDraftErr != nil {
		logger.Errorf("error occurred while deleting draft for draft id %v. Error %v", draftID, err)
		constants.RespondWithGolaError(ctx, deleteDraftErr)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "deleted",
	})
	return
}

// GetPreviewDraft godoc
// @Tags draft
// @Summary GetPreviewDraft
// @Description get preview draft for a given draft id
// @Accept json
// @Param request body request.DraftURIRequest true "Request Body"
// @Success 200
// @Failure 400 {object} golaerror.Error
// @Failure 406 {object} golaerror.Error
// @Failure 500 {object} golaerror.Error
// @Router /api/post/v1/draft/preview-draft/:draft_id [get]
func (controller DraftController) GetPreviewDraft(ctx *gin.Context) {
	logger := logging.GetLogger(ctx).WithField("class", "DraftController").WithField("method", "GetPreviewDraft")

	token, err := utils.GetIDToken(ctx)
	if err != nil {
		logger.Error("id token not found", err)
		ctx.JSON(http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	userUUID, _ := uuid.Parse(token.UserId)
	logger.Infof("Entered controller to upsert draft request for user %v", userUUID)


	var draftURIRequest request.DraftURIRequest
	if err := ctx.ShouldBindUri(&draftURIRequest); err != nil {
		logger.Errorf("unable to bind path parameter %v", err)
		constants.RespondWithGolaError(ctx, &constants.PayloadValidationError)
		return
	}

	logger.Infof("Request body bind successful with get draft request for user %v", "12")
	draftID, _ := uuid.Parse(draftURIRequest.DraftID)
	draftData, draftGetErr := controller.service.ValidateAndGetDraft(ctx, draftID, token)
	if draftGetErr != nil {
		logger.Errorf("Error occurred in draft service while saving tagline for user %v. Error %v", "12", draftGetErr)
		constants.RespondWithGolaError(ctx, draftGetErr)
		return
	}

	logger.Infof("writing response to draft data request for user %v %s", "12", draftURIRequest.DraftID)

	ctx.JSON(http.StatusOK, draftData)
}

func NewDraftController(service service.DraftService) DraftController {
	return DraftController{
		service: service,
	}
}
