package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type StatusCode string

const (
	StatusOK       StatusCode = "OK"
	StatusNotFound StatusCode = "NOT FOUND"
	StatusDeleted  StatusCode = "DELETED"
)

// QueryResponse - tagged as a sub-struct for swagger purposes
type QueryResponse struct {
	// The status
	// Required: true
	Status StatusCode `json:"status"`
	// The detailed message
	// Required: false
	Message string `json:"message,omitempty"`
}

/*
swagger:response queryResponse

Not actually returned, but the swagger docs generate strangely without it!
Used to wrap things for documentation purposes.
*/
type MetaResponse struct {
	// The query response
	// in: body
	Body QueryResponse
}

/*
swagger:route GET / handler getHandler

Produces:
- application/json

Schemes: http

Responses:
  default: queryResponse
*/
func GetHandler(c *gin.Context) {
	// The actual response goes here
	resp := QueryResponse{
		Status:  StatusOK,
		Message: "10-4 good buddy",
	}
	c.JSON(http.StatusOK, resp)
}

func main() {
	r := gin.Default()

	log.Info("starting handler")
	r.GET("/", GetHandler)

	r.Run()
}
