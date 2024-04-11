# Do I stay or do I Go?

Simple test application using the [Go](https://go.dev) language 
and the [htmx](https://htmx.org) Javascript library to generate interactive DOM
segments.

## Demonstrates the following key K8s capabilities
* Pod and Deployment creation
* NodePort routing
* Ingress routing
* Configuration by environment variable
* health checks

## Ingress Installation
Tested with [version 3.4](https://hub.docker.com/r/nginx/nginx-ingress/tags)

## Helm Install

[Instructions](https://docs.nginx.com/nginx-ingress-controller/installation/installing-nic/installation-with-helm/)

required to get the CRDs to install ahead of the chart install
[Install CRD](https://docs.nginx.com/nginx-ingress-controller/installation/installing-nic/installation-with-manifests/#create-custom-resources)

```bash
cd kubernetes-ingress/deployments/common
kubectl apply -f ingress-class.yaml

cd kubernetes-ingress/deploy
kubectl apply -f crds.yaml
customresourcedefinition.apiextensions.k8s.io/dnsendpoints.externaldns.nginx.org created
customresourcedefinition.apiextensions.k8s.io/globalconfigurations.k8s.nginx.org created
customresourcedefinition.apiextensions.k8s.io/policies.k8s.nginx.org created
customresourcedefinition.apiextensions.k8s.io/transportservers.k8s.nginx.org created
customresourcedefinition.apiextensions.k8s.io/virtualserverroutes.k8s.nginx.org created
customresourcedefinition.apiextensions.k8s.io/virtualservers.k8s.nginx.org created
```

### first use the OCI pulled chart
```bash
helm pull oci://ghcr.io/nginxinc/charts/nginx-ingress --untar --version 1.1.2
```

# goto the pulled directory
cd nginx-ingress

# apply the cart using the default values
helm install nginx-ingress-3-4 .

# I chose the name nginx-ingress-3-4 for the helm install (could be bad for upgrades)
# also . chooses the default values.yaml (look it over first for your desired configuration)



So the default configuration listens on port 80 and 443