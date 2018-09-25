package main

import (
	"go-worker/factory"

	"fmt"

	"net/http"
	//"os"

	"github.com/gin-gonic/gin"
)

func payload(c *gin.Context) {
	if c.Request.Method != "POST" {
		//fmt.Fprint(c.Writer, http.StatusMethodNotAllowed)

		p := c.Request.Host
		fmt.Println("test", p)

		go func(p string) {
			factory.JobQueue <- factory.Job{Payload: p}
		}(p)
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}

func main() {
	//fmt.Println(os.Getenv("GO_MAX_WORKERS"))
	r := gin.Default()
	r.Any("/", payload)
	dispatch := factory.NewDispatcher(factory.MAX_WORKERS)
	dispatch.Run()
	r.Run()
}
