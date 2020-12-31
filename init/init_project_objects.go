package init

import (
	"github.com/neo4j/neo4j-go-driver/neo4j"
	"post-api/configuration"
	"post-api/controller"
	"post-api/repository"
	"post-api/service"
	"post-api/utils"
)

var (
	draftController     controller.DraftController
	interestsController controller.InterestsController
	postController      controller.PostController
)

func Objects(db neo4j.Session, configData *configuration.ConfigData) {
	draftRepository := repository.NewDraftRepository(db)
	draftService := service.NewDraftService(draftRepository)
	draftController = controller.NewDraftController(draftService)
	interestsRepository := repository.NewInterestRepository(db)
	interestsService := service.NewInterestsService(interestsRepository)
	interestsController = controller.NewInterestsController(interestsService)
	postRepository := repository.NewPostsRepository(db)
	postValidator := utils.NewPostValidator(configData)
	postService := service.NewPostService(postRepository, draftRepository, postValidator)
	postController = controller.NewPostController(postService)
}
