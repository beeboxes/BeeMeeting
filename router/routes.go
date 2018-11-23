package router

import (
	"github.com/gin-gonic/gin"
	"github.com/manyminds/api2go"
)

func init() {
	r := gin.Default()
	api := api2go.NewAPIWithRouting(
		"api",
		api2go.NewStaticResolver("/"),
		Gin(r),
	)

	_ := api
}
