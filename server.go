package main

import (
	"fmt"
	"net/http"

	"github.com/proishan11/mux-rest/repository"

	"github.com/proishan11/mux-rest/service"

	"github.com/proishan11/mux-rest/controller"

	"github.com/proishan11/mux-rest/router"
)

var (
	postRepository repository.PostRepo       = repository.NewFirestoreRepo()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(response http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(response, "Server up and running")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
