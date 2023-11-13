# Kubernetes (custom-sample) Operator - ddbctl-dtp-operator

## Overview

The `ddbctl-dtp-operator` (DynamoDB Controller - Delete Table Partition) – a Kubernetes Operator written in Golang using the Kubebuilder framework. This operator simplifies the process of managing DynamoDB table partitions within a Kubernetes environment.

## Features

**Automated Partition Deletion:**

Efficiently delete partitions in DynamoDB tables through the declarative configuration provided by Kubernetes Custom Resources.

**Built with Kubebuilder:**

Leveraging the Kubebuilder framework ensures adherence to best practices for building scalable and maintainable Kubernetes operators.

## How it Works

The `ddbctl-dtp-operator` extends the functionality of Kubernetes by introducing a custom resource, DeleteTablePartitionDataJob. This resource allows you to express your intent to delete a specific partition in a DynamoDB table directly through Kubernetes manifests.

```yaml
apiVersion: ddbctl.operators.jittakal.io/v1alpha1
kind: DeleteTablePartitionDataJob
metadata:
  name: delete-table-partition-data-job
spec:
  tableName: my-dynamodb-table
  partitionValue: partition-key-value
  endpointURL: http://dynamodb.local:8000
  awsRegion: us-east-1
```

## Pre-requisites

- Kubernetes local cluster using microk8s up and running.
- DynamoDB local kubernetes deployment [steps](https://medium.com/@jittakal/running-dynamodb-local-within-microk8s-a-step-by-step-guide-with-sample-code-38aac0aea803) followed.
- Kubebuilder and related pre-requisites installed.

## Pre-requisites

Before you start working with the ddbctl-dtp-operator, make sure you have the following prerequisites in place:

**Kubernetes Local Cluster:**

Ensure that you have a local Kubernetes cluster set up using microk8s. If you haven't set up microk8s, you can follow the official documentation or use the microk8s up command to initialize a local cluster.

**DynamoDB Local Kubernetes Deployment:**

Follow the [steps](https://medium.com/@jittakal/running-dynamodb-local-within-microk8s-a-step-by-step-guide-with-sample-code-38aac0aea803) outlined in the DynamoDB local Kubernetes deployment guide to deploy DynamoDB locally within your microk8s cluster. This local DynamoDB instance will serve as the database for your ddbctl-dtp-operator.

**Kubebuilder Installation:**

Ensure that Kubebuilder and its related dependencies are installed on your local machine. You can follow the Kubebuilder installation guide provided in the official documentation to set up the necessary tooling for building Kubernetes operators.

```bash
$ brew install kubebulder

$ brew install kustomize
```


## Steps for reference

- Clone empty repository

```bash
$ cd ~\workspace\git.ws
$ git clone https://github.com/jittakal/ddbctl-dtp-operator
$ cd ddbctl-dtp-operator
```

- Initialize the Go Moudle

```bash
$ go mod init github.com/jittakal/ddbctl-dtp-operator
```

- Kubebuilder init

```bash
$ kubebuilder init --domain operators.jittakal.io --repo github.com/jittakal/ddbctl-dtp-operator
```

- Custom Resource Defination - API

```bash
$ kubebuilder create api --group ddbctl --version v1alpha1 --kind DeleteTablePartitionDataJob
```

- Modify DeleteTablePartitionDataJob Spec

```go
type DeleteTablePartitionDataJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// DynamoDB Table Name
	// +kubebuilder:validation:Required
	TableName string `json:"tableName"`

	// Partition Value
	// +kubebuilder:validation:Required
	PartitionValue string `json:"partitionValue"`

	// Endpoint URL - Optional
	// +kubebuilder:validation:Optional
	EndpointURL string `json:"endpointURL,omitempty"`

	// AWS Region
	// +kubebuilder:default := "us-east-1"
	// +kubebuilder:validation:Required
	AWSRegion string `json:"awsRegion"`
}
```

- Implement Controller Reconcile Function

- Build the Controller 

```bash
$ make build # manifest generate fmt vet
```

- Build and publish docker controller manager images

```bash
$ make docker-build docker-push
```

- Deploy the CRD

```bash
$ make deploy
```

- Verify the deployment

```bash
$ kubectl get crd # new entry for deletetablepartitiondatajobs.ddbctl.operators.jittakal.io

$ kubectl get namespaces # new entry for ddbctl-dtp-operator-system

$ kubectl get pods -n ddbctl-dtp-operator-system

$ kubctl logs -f <<pod-name-from-above-command>> -n ddbctl-dtp-operator-system
```

- Install Sample Customer Resource

```bash
$ #Open in new terminal
$ kubectl create -f config/samples/ddbctl_v1alpha1_deletepartitiondatajob.yaml

$ kubectl get jobs # default - namespace

$ kubectl get pods

$ kubectl logs <<podname-of-delete-table-partition-data-job>> # verify log table name and number of records delete / delete summary report on job completion
```


## Open Issues

- Limit the operator's ability to create Jobs exclusively within the "default" namespace, and adjust the RBAC permissions of the controller manager accordingly.


## Reference

- [Docker image - operator controller manager](https://hub.docker.com/repository/docker/jittakal/ddbctl-dtp-operator/general)
- [Docker image - job](https://hub.docker.com/repository/docker/jittakal/go-dynamodb-partition-delete/general)
