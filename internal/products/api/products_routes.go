package api

func MapProductsRoutes(controller *ProductsController) {
	v1 := controller.echo.Group("/api/v1")
	products := v1.Group("/products")

	products.POST("", controller.createProduct())
	products.GET("/:id", controller.getProductByID())
}
