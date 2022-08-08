package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.StaticFile("/", "index.html")

	router.GET("/tale_of_two_cities", func(ctx *gin.Context) {
		// Simple File Serving
		//ctx.File("a_tale_of_two_cities.txt")

		// Data way of sharing
		f, err := os.Open("a_tale_of_two_cities.txt")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		defer f.Close()
		// IO Reader way
		/* data, err := io.ReadAll(f)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.Data(http.StatusOK, "test/plain", data) */
		///Stream way of handling
		ctx.Stream(streamer(f))
	})

	router.GET("/great_expectations", func(ctx *gin.Context) {
		f, err := os.Open("great_expectations.txt")
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		defer f.Close()
		fi, err := f.Stat()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.DataFromReader(http.StatusOK, fi.Size(), "text/plain", f, map[string]string{
			"Content-Disposition": "attachment;filename=great_expectations.txt",
		})
	})

	log.Fatal(router.Run(":3000"))

}

func streamer(r io.Reader) func(io.Writer) bool {
	return func(step io.Writer) bool {
		for {
			buf := make([]byte, 4*2^10)
			if _, err := r.Read(buf); err == nil {
				_, err := step.Write(buf)
				return err == nil
			} else {
				return false
			}
		}
	}
}
