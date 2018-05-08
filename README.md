# Kubernetes healthcheck util

This is a a quick tool for running kubernetes healthchecks against all pods in a namespace.


## To build:
```
1. go get -u -v github.com/arehmandev/kubehealth
1. Ensure you have dep - https://github.com/golang/dep
2. cd $GOPATH/src/github.com/arehmandev/kubehealth && dep ensure
3. go build
```

## To run: 
```
1. Ensure your kubeconfig is set to the right context (pointing to correct cluster)
2. ./kubehealth *namespace*
```

Note - this runs faster than simply doing a kubectl get pods -o json and parsing through with jq.