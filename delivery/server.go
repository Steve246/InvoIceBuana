package delivery

import (
	"invoiceBuana/config"
	"invoiceBuana/delivery/controller"
	"invoiceBuana/manager"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type appServer struct {
	usecaseManager manager.UsecaseManager
	// router         *gin.RouterGroup
	routerDev *gin.RouterGroup
	engine    *gin.Engine
	host      string
}

func Server() *appServer {
	router := gin.Default()

	appConfig := config.NewConfig()

	infra := manager.NewInfra(appConfig)

	repoManager := manager.NewRepositoryManager(infra)

	usecaseManager := manager.NewUsecaseManager(repoManager)

	host := appConfig.Url
	// Add CORS middleware
	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowAllOrigins:  true,
		AllowHeaders:     []string{"Origin", "Date", "Content-Length", "Content-Type", "Content-Disposition", "Accept", "X-Requested-With", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Request-Method", "Access-Control-Request-Headers", "Authorization", "token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	return &appServer{
		usecaseManager: usecaseManager,
		engine:         router,
		// router:         router.Group("", nil),

		routerDev: router.Group("invoice/"),
		host:      host,
	}
}

func (a *appServer) initControllers() {
	// buat daftarin controller ada disini
	// setiap controller, isinya harus ada isian dari usecaseManager

	controller.NewItemController(a.routerDev, a.usecaseManager.ItemUsecase())

	controller.NewCustomerController(a.routerDev, a.usecaseManager.CustomerUsecase())

}

func (a *appServer) Run() {
	a.initControllers()
	err := a.engine.Run(a.host)
	if err != nil {
		panic(err)
	}
}
