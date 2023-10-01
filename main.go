package main

import (
	"btpn/router"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	engine := gin.Default()
	router.NewRouter(engine)
	log.Fatalf("error : %+v", engine.Run(":8080"))
}
