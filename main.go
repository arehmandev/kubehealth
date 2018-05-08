package main

import (
	"fmt"
	"log"
)

var (
	healthCheckStatus  bool
	healthCheckRetries = 5
	healthCheckBanner  = `
================================================
	PERFORMING HEALTHCHECK`
)

func main() {

	namespace := checkParams()
	clientset := configureClient()

	fmt.Println(healthCheckBanner)

	podList := getPodList(namespace, clientset)

	// Retry loop
	for index := 0; index < healthCheckRetries; index++ {
		try := index + 1

		// Check health of pods for all cases
		fmt.Printf("Healthcheck try %v of %v\n", try, healthCheckRetries)
		healthCheckStatus = runPodHealthcheck(podList)
		if !healthCheckStatus {
			fmt.Printf("Healtcheck failed on try %v- retrying\n", try)
			continue
		} else {
			break
		}
	}

	if !healthCheckStatus {
		log.Fatalln("Healtcheck failed")
	} else {
		fmt.Println("Healthcheck succeeded")
	}
}
