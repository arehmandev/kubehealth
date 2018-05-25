package main

import (
	"fmt"

	"k8s.io/api/core/v1"
)

func runPodHealthcheck(podList *v1.PodList) (healthCheckStatus bool) {

	// defaults to true
	healthCheckStatus = true

	for _, pod := range podList.Items {

		if pod.Status.Phase != "Running" {
			if pod.Status.Phase != "Succeeded" {
				healthCheckStatus = false
				fmt.Printf("Reason: [ERROR] POD IS STILL SCHEDULING - NAME:%v INFO: %v\n", pod.Name, pod.Status.Phase)
			}
		}

		// totalContainers := len(pod.Spec.Containers)
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if !containerStatus.Ready {

				containerName := containerStatus.Name

				// Check if containers lost in a running state [RARE]
				if containerStatus.State.Running != nil {
					fmt.Printf("[ERROR] Pod: %s Container: %s | Reason: [BAD RUNNING STATE] | Status: Container appears to be running with false ready state: %v\n", pod.Name, containerName, containerStatus.State.Running)
					healthCheckStatus = false
				}

				// Check if containers in waiting state - means pod belongs to deployment / statefulset
				if containerStatus.State.Waiting != nil {
					fmt.Printf("[ERROR] Pod: %s Container: %s | Reason: [WAITING] | Status: %s\n", pod.Name, containerName, containerStatus.State.Waiting.Reason)
					healthCheckStatus = false
				}

				// Check if container terminated and exit code isn't 0
				if containerStatus.State.Terminated != nil {
					if containerStatus.State.Terminated.ExitCode != 0 {
						fmt.Printf("[ERROR] Pod: %s Container: %s | Reason: [TERMINATED BADLY] | Status: %s | Exit code: %v\n", pod.Name, containerName, containerStatus.State.Terminated.Reason, containerStatus.State.Terminated.ExitCode)
						healthCheckStatus = false
					}
				}

			}
		}
	}

	return healthCheckStatus
}
