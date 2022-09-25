package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

var (
	HOST = os.Getenv("HOST")
)

func RNG() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000)
}
func handleHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"randomNumber": RNG(),
	})
}

func main() {
	router := gin.Default()

	//allowing cors for localhost interport communication for dev, big nono for deployment
	corsConf := cors.DefaultConfig()
	corsConf.AllowAllOrigins = true
	router.Use(cors.New(corsConf))

	router.GET("/", handleHome)

	router.Run(HOST)
}
