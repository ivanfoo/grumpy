package main

import (
	"flag"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	listenAddr = flag.String("listen", ":8080", "listen address")
	niceWords  = []string{
		"moron",
		"idiot",
		"pussy",
		"asshole",
		"wanker",
		"looser",
		"pillock",
		"muppet",
		"twat",
		"dickhead",
		"cunt",
	}
)

func sayNiceOneFrom(list []string) string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Int() % len(list)

	return list[n]
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	flag.Parse()

	router := gin.New()
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Get out of here, %s!", sayNiceOneFrom(niceWords))
	})

	router.GET("/hi/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "You are a fucking %s, %s\n", sayNiceOneFrom(niceWords), name)
	})

	router.Run(*listenAddr)
}
