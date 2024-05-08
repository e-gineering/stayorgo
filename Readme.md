# Do I stay or do I Go?

Simple demo application as an example of an external state store deployed capability.
Built using the [Go](https://go.dev) language with [GoLand](https://www.jetbrains.com/go/)
and the [htmx](https://htmx.org) Javascript library to generate interactive DOM
segments.
The application requires an available [Redis](https://redis.io/docs/latest/operate/oss_and_stack/) key value store
configured through environment variables.

Demonstrated with a local K8s cluster using [Orbstack](https://orbstack.dev)
and [k9s](https://k9scli.io) for visualization

## Demonstrates the following key K8s capabilities
* Kubernetes Pod and Deployment
* NodePort routing
* Ingress routing
* Configuration through environment variables
* Deployment Health checks

# Demo
Go from local development through to a deployed application

## Local Development
- [ ] Start the local container for Redis via docker-compose
- [ ] Launch the Application from the IDE with configuration
- [ ] Verify 'Hello' endpoint using localtest.http
- [ ] Show the application running on http://localhost:8080
## Pod Deployment
_note_ Redis is already deployed as a running service
- [ ] Docker build of the container ```docker build -t stayorgo:0.1 .```
- [ ] Deploy as a pod scripts/k8s/pod-deployment/stayorgo-pod.yaml
- [ ] Local port-forward to port 8080 ```kubectl port-forward <pod> 8080:8080```
- [ ] Show the application executing (connected to Redis in K8s)
- [ ] Show shutdown without restart
## Service with Nodeport
- [ ] Apply scripts/k8s/service-nodeport/stayorgo-service-nodeport.yaml
- [ ] Show the application running on http://localhost:30007
## Service with Ingress allowed with Nginx Ingress Controller
- [ ] Apply scripts/k8s/service-ingress/stayorgo-service-ingress.yaml
- [ ] Show the application running on http://stayorgo.k8s.orb.local
- [ ] Show scaling the application
- [ ] Show deleting a replica
- [ ] Show restarting a deployment

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

# Reference Links
Why readiness and healthiness tests are important
[Hello IT, have you tried turning it off and on again?](https://www.youtube.com/watch?v=t2F1rFmyQmY)

Spending more on Kubernetes due to overprovisioning
[InfoQ Article](https://www.infoq.com/news/2024/03/cncf-finops-kubernetes-overspend/?utm_campaign=infoq_content&utm_source=infoq&utm_medium=feed&utm_term=DevOps)