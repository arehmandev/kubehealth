package main

import (
	"fmt"
	"log"
	"time"
)

var (
	healthCheckStatus  bool
	healthCheckRetries = 5
	healthCheckSleep   = 10
	healthCheckBanner  = `
==============================================================
	PERFORMING HEALTHCHECK`
)

func main() {

	namespace, labels := checkParams()
	clientset := configureClient()
	sleepDuration := time.Duration(int64(healthCheckSleep)) * time.Second

	fmt.Println(healthCheckBanner)

	// Retry loop
	for index := 0; index < healthCheckRetries; index++ {
		try := index + 1

		podList := getPodList(namespace, labels, clientset)

		// Check health of pods for all cases
		fmt.Printf("Healthcheck try %v of %v\n", try, healthCheckRetries)
		healthCheckStatus = runPodHealthcheck(podList)
		if !healthCheckStatus {
			fmt.Printf("Healthcheck failed on try %v- retrying after sleep (%v seconds)\n", try, healthCheckSleep)
			time.Sleep(sleepDuration)
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
