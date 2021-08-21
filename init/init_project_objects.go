package init

import (
	"github.com/jmoiron/sqlx"
	"post-api/configuration"
	"post-api/helper"
	"post-api/story/controller"
	"post-api/story/repository"
	"post-api/story/service"
	"post-api/story/utils"
)

var (
	draftController     controller.DraftController
	interestsController controller.InterestsController
	postController      controller.PostController
)

func Objects(db *sqlx.DB, configData *configuration.ConfigData) {
	manager := helper.NewTransactionManager(db)
	draftRepository := repository.NewDraftRepository(db)
	draftService := service.NewDraftService(draftRepository)
	draftController = controller.NewDraftController(draftService)
	interestsRepository := repository.NewInterestRepository(db)
	interestsService := service.NewInterestsService(interestsRepository)
	interestsController = controller.NewInterestsController(interestsService)
	postRepository := repository.NewPostsRepository(db)
	postValidator := utils.NewPostValidator(configData)
	previewPostRepository := repository.NewAbstractPostRepository(db)
	postService := service.NewPostService(postRepository, draftRepository, postValidator, previewPostRepository, interestsRepository, manager)
	postController = controller.NewPostController(postService)
}
