package main

import (
	"fmt"
	"net/http"
	"time"
)

var IP_ADDRESS = "10.0.2.161"
var PORT_NUMBER = "9001"
var BASE_URL = "http://" + IP_ADDRESS + ":" + PORT_NUMBER
var RESOURCE_PATH = "/resource"
var FINAL_YAML_PATH = "/final"

func main() {

	restURL := BASE_URL + RESOURCE_PATH

	fileName := "ml-comparison-pipeline-1.0-before"
	// fileName := "sample"

	oriPath := "./input/" + fileName + ".yaml"
	resultPath := "./output/" + fileName + "-after.yaml"

	// first. request information by node allocator
	ackBody := ReqResourceAllocInfo(restURL, oriPath)

	// second. store the request info into ketidatabase/validity
	CreateRequestValidityInMySQL(fileName)

	// wait for 10 seconds
	fmt.Println("Waiting for 10 seconds...")
	time.Sleep(time.Second * 10)

	// thirt. made final workload yaml by response data
	finalYaml := MadeFinalWorkloadYAML(ackBody, oriPath, resultPath)

	// fourth. make workload yaml file
	MakeYamlFile(finalYaml, resultPath)

	// fifth. send deploied yaml data to node allocator (kware).
	ack, _ := SEND_REST_DATA(BASE_URL+FINAL_YAML_PATH, finalYaml)
	if ack.StatusCode == http.StatusOK {
		fmt.Println("=== Done ===")
	} else {
		fmt.Printf("[FinalYAML] Request failed with status: %s\n", ack.Status)
	}

	// sixth. save final yaml file to database
	StoreInMySQL(fileName, finalYaml)

	// fin. update the validity of yaml.
	UpdateResponseInMySQL(fileName)
}
