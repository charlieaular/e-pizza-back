package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"e-pizza-backend/src/database"
	"e-pizza-backend/src/routes"
)

func main() {

	mysqlIns := database.MysqlConn

	db := mysqlIns.Init()

	r := gin.Default()

	r.Use(cors.Default())

	routes.RegisteCategoryRoutes(r, db)
	routes.RegisteProductRoutes(r, db)
	routes.RegisteCartRoutes(r, db)

	r.Run(":9000")
}
