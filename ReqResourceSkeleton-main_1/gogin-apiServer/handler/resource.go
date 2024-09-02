package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"

	ys "main/ystruct"
)

func ResourceHandler(ctx *gin.Context) {

	recvData, err := ctx.GetRawData()
	if err != nil {
		ctx.YAML(400, gin.H{
			"message": fmt.Sprintf("%+v", err),
		})
		return
	}

	time.Sleep(time.Millisecond * 20)

	recvStr := fmt.Sprintf("%s", recvData)
	log.Println("----------[RecvMsg]----------\n" + recvStr)

	var reqBody ys.ReqResource
	yaml.Unmarshal([]byte(recvData), &reqBody)

	size := len(reqBody.Request.Containers)
	containers := make([]ys.Container, size)

	for idx, val := range reqBody.Request.Containers {

		if val.Name == "" {
			// DAG와 같은 컨테이너가 아닌 경우
			continue
		}

		containers[idx].Name = val.Name
		containers[idx].Cluster = "123"
		containers[idx].Node = val.Name + "-KETI"
	}

	// sample response
	ctx.YAML(http.StatusOK, ys.RespResource{
		Response: ys.Response{
			ID:               "eddieKim",
			Date:             "2024-01-01 14:25:59",
			Container:        containers,
			PriorityClass:    "TEST-Class",
			Priority:         "100000",
			PreemptionPolicy: "Never",
		},
	})
}

func FinalHandler(ctx *gin.Context) {

	recvData, err := ctx.GetRawData()
	if err != nil {
		ctx.YAML(400, gin.H{
			"message": fmt.Sprintf("%+v", err),
		})
		return
	}

	time.Sleep(time.Millisecond * 20)

	recvStr := fmt.Sprintf("%s", recvData)
	log.Println("----------[RecvMsg]----------\n" + recvStr)

	ctx.YAML(http.StatusOK, gin.H{
		"status": "test",
		"value":  "OK",
		"msg":    "",
	})
}
