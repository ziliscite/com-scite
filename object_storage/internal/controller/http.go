package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ziliscite/com-scite/object_storage/internal/repository"
	"net/http"
	"os"
)

type HttpServer struct {
	st repository.Read
}

func NewHttpServer(st repository.Read) *HttpServer {
	return &HttpServer{
		st: st,
	}
}

func (h *HttpServer) Serve(ctx *gin.Context) {
	signedUrl := ctx.Param("signedUrl")

	// Decrypt the signed URL
	filePath, err := h.st.Get(signedUrl)
	if err != nil {
		ctx.String(http.StatusForbidden, "Invalid or expired URL")
		return
	}

	// Check if the file exists
	if _, err = os.Stat(filePath); os.IsNotExist(err) {
		ctx.String(http.StatusNotFound, fmt.Sprintf("File not found: %s", filePath))
		return
	}

	// Serve the file
	ctx.File(filePath)
}

func (h *HttpServer) Run(port int) error {
	r := gin.Default()

	r.GET("/file/:signedUrl", h.Serve)

	return r.Run(fmt.Sprintf(":%d", port))
}
