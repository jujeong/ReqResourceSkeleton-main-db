package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func EchoTestHandler(ctx *gin.Context) {

	recvData, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(400, gin.H{
			"message": fmt.Sprintf("%+v", err),
		})
		return
	}

	time.Sleep(time.Millisecond * 20)

	recvStr := fmt.Sprintf("%s", recvData)
	log.Printf("[RecvMsg]: " + recvStr)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "test",
		"value":  "OK",
		"msg":    "",
	})
}
