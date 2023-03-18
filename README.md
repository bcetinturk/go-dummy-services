# Linkerd Service Mesh Demo
This repository contains two web applications written in Golang, deployed on Kubernetes, and using Linkerd service mesh. The purpose of this repository is to experiment and discover features of Linkerd service mesh.

## Prerequisites
- A running Kubernetes cluster
- Kustomize installed on your machine

## Installation
To install the web applications, run the following commands:

```
kubectl apply -k kustomize/overlays/servicea
kubectl apply -k kustomize/overlays/serviceb
```
This will deploy both applications to your Kubernetes cluster.

## Usage
To access the web applications, use the following commands:

```
kubectl port-forward svc/servicea 8080:80
kubectl port-forward svc/serviceb 8081:80
```
Then, open your web browser and go to http://localhost:8080 or http://localhost:8081 to access the applications.

## Credits
Author: Barış Çetintürk