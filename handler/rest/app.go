package rest

import (
	"assignment-2/database"
	"assignment-2/docs"
	"assignment-2/repository/item_repository/item_pg"
	"assignment-2/repository/orderrepository/orderpg"
	"assignment-2/service"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"

	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

const port = ":8080"

func StartApp() {

	database.InitializeDB()

	db := database.GetDB()

	route := gin.Default()

	docs.SwaggerInfo.Title = "Belajar DDD"
	docs.SwaggerInfo.Description = "Ini adalah API dengan pattern DDD"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	// docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http"}

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	testRoute(route)
	orderRoute(route, db)

	fmt.Println("Server running on PORT =>", port)
	route.Run(port)

}

func testRoute(route *gin.Engine) {

	route.GET("/ping")
}

func orderRoute(route *gin.Engine, db *sqlx.DB) {

	orderRepository := orderpg.NewOrderPG(db)
	itemRepository := item_pg.NewItemPG(db)
	orderService := service.NewOrderService(orderRepository, itemRepository)
	orderHandler := NewOrderHandler(orderService)

	// result, err := orderRepository.CreateOrder(&dto.OrderRequest{
	// 	OrderedAt:    "2022-10-15 00:00:00",
	// 	CustomerName: "fauzi",
	// })

	// fmt.Println(result)
	// fmt.Println(err)

	orderRoute := route.Group("/orders")
	{
		orderRoute.GET("/", orderHandler.FindAll)
		orderRoute.POST("/", orderHandler.CreateOrder)
		orderRoute.PUT("/:orderId", orderHandler.UpdateOrder)
		orderRoute.DELETE("/:orderId", orderHandler.Delete)
		orderRoute.GET("/:orderId", orderHandler.FindById)
	}
}
