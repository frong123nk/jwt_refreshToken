package main

import (
	"fmt"
	"jwt/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}
	router := gin.Default()
	routers.SetupRouter(router)
	router.Run(":88")
}
