# Do I stay or do I Go?

Simple test application using the [Go](https://go.dev) language 
and the [htmx](https://htmx.org) Javascript library to generate interactive DOM
segments.
Expects a local K8s cluster using [Orbstack](https://orbstack.dev)
## Demonstrates the following key K8s capabilities
* Pod and Deployment creation
* NodePort routing
* Ingress routing
* Configuration by environment variable
* health checks

# Demo Steps
Stay or Go list of items application.
Going from development through to a deployed application
## Local Development
- [ ] Start the local container for Redis via docker-compose
- [ ] Launch the Application from the IDE with configuration
- [ ] Verify 'Hello' endpoint using localtest.http
- [ ] Show the application running on http://localhost:8080
## Pod Deployment
_note_ Redis is already deployed as a running service
- [ ] Docker build of the container
- [ ] Deploy as a pod scripts/k8s/pod-deployment/stayorgo-pod.yaml
- [ ] Local port-forward to port 8080
- [ ] Show the application executing (connected to Redis in K8s)
## Service with Nodeport
- [ ] Apply scripts/k8s/service-nodeport/stayorgo-service-nodeport.yaml
- [ ] Show the application running on http://localhost:30007
## Service with Ingress allowed with Nginx Ingress Controller
- [ ] Apply scripts/k8s/service-ingress/stayorgo-service-ingress.yaml
- [ ] Show the application running on http://stayorgo.k8s.orb.local

# Ingress Installation (preliminary to the demo)

## Helm Install
Tested with [version 3.4](https://hub.docker.com/r/nginx/nginx-ingress/tags)

Install the Ingress controller and CRDs directly using the 
default values and the yaml.
[Instructions](https://docs.nginx.com/nginx-ingress-controller/installation/installing-nic/installation-with-helm/)

Get the CRDs to install ahead of the Helm chart installation
[Install CRD](https://docs.nginx.com/nginx-ingress-controller/installation/installing-nic/installation-with-manifests/#create-custom-resources)

```bash
cd kubernetes-ingress/deployments/common
kubectl apply -f ingress-class.yaml

cd kubernetes-ingress/deploy
kubectl apply -f crds.yaml
# results in 
customresourcedefinition.apiextensions.k8s.io/dnsendpoints.externaldns.nginx.org created
customresourcedefinition.apiextensions.k8s.io/globalconfigurations.k8s.nginx.org created
customresourcedefinition.apiextensions.k8s.io/policies.k8s.nginx.org created
customresourcedefinition.apiextensions.k8s.io/transportservers.k8s.nginx.org created
customresourcedefinition.apiextensions.k8s.io/virtualserverroutes.k8s.nginx.org created
customresourcedefinition.apiextensions.k8s.io/virtualservers.k8s.nginx.org created
```

## Helm install first uses the OCI pulled chart
The version is the Chart version and not the application version
```bash
helm pull oci://ghcr.io/nginxinc/charts/nginx-ingress --untar --version 1.1.2
```
### goto the pulled directory
cd nginx-ingress

### apply the controller using the default values
helm install nginx-ingress-3-4 .

I chose the name nginx-ingress-3-4 for the helm install
also . chooses the default values.yaml (look it over first for your desired configuration)

The default configuration listens on port 80 and 443 for routing