package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"oneclick/server/controller"
	"oneclick/server/middleware"
)

func Post(c *gin.Context) {
	res, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Println(string(res))
}

func CollectRoute() {
	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(middleware.RecoveryMiddleware())
	//postController := processor.NewIPostController()
	//r.POST("/ArchivedPingsInfo/post", controller.ArchivedPingsInfo)
	//r.POST("/ScopedBindingBuilder/post", controller.ScopedBindingBuilder)
	//r.POST("/BundleUnbinding/post", controller.BundleUnbinding)
	//r.POST("/CheckStartup/post", controller.CheckStartup)
	//r.POST("/CheckKey/post", controller.CheckKey)
	//r.POST("/", Post)

	postController := controller.NewIPostController()
	r.POST("/ArchivedPingsInfo/post", postController.Info)
	r.POST("/ScopedBindingBuilder/post", postController.Update)
	r.POST("/BundleUnbinding/post", postController.Delete)
	r.POST("/CheckStartup/post", controller.CheckStartup)
	r.POST("/CheckKey/post", postController.Show)
	//r.POST("/", Post)

	r.Run(":8082")
}
