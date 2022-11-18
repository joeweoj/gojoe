# gojoe server

## Minikube

### setup 

1. `minikube start`
1. `eval $(minikube -p minikube docker-env)`
1. `minikube image load joeserver:0.1`
1. `minikube tunnel` (new terminal)

### deployment

1. `k apply -f deployment.yaml`
1. `k expose deployment joeserver-deployment --type=LoadBalancer --port=8090 --target-port=8090`
1. `k describe svc joeserver-deployment` (note IP)
1. Navigate to `<IP>:8090`