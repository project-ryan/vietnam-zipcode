package api

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quanghia24/vietnam-zipcode/config"
	"github.com/quanghia24/vietnam-zipcode/service"
)

type Server struct {
	router  *gin.Engine
	service service.ZipcodeService
}

func NewServer(service service.ZipcodeService, config *config.Config) *Server {
	server := &Server{
		service: service,
	}

	router := gin.Default()

	// Add recovery middleware
	router.Use(gin.Recovery())

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Add cache headers middleware with configurable max-age
	router.Use(func(c *gin.Context) {
		c.Header("Cache-Control", "public, max-age="+strconv.Itoa(config.CacheMaxAge))
		c.Next()
	})

	// Add request logging middleware
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.GET("/", server.getLocation)
	router.GET("/health", server.healthCheck)

	server.router = router
	return server
}

func (s *Server) Start(address string) error {
	log.Printf("Starting server on %s", address)
	return s.router.Run(address)
}

func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"timestamp": time.Now().UTC(),
	})
}

func errorResponse(err string, code string, detail string) gin.H {
	return gin.H{
		"Error":   err,
		"Code":    code,
		"Details": detail,
	}
}
