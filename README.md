# Kubernetes (custom-sample) Operator - ddbctl-dtp-operator

## Overview

The `ddbctl-dtp-operator` (DynamoDB Controller - Delete Table Partition) – a Kubernetes Operator written in Golang using the Kubebuilder framework. This operator simplifies the process of managing DynamoDB table partitions within a Kubernetes environment.

## Features

**Automated Partition Deletion:**

Efficiently delete partitions in DynamoDB tables through the declarative configuration provided by Kubernetes Custom Resources.

**Built with Kubebuilder:**

Leveraging the Kubebuilder framework ensures adherence to best practices for building scalable and maintainable Kubernetes operators.

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

**Clone ddbctl-dtp-operator git repository:**

```bash
$ cd ~\workspace\git.ws
$ git clone https://github.com/jittakal/ddbctl-dtp-operator
$ cd ddbctl-dtp-operator
```
**Deploy DeleteTablePartitionDataJob CRD:**

```bash
make deploy
```

```bash
$ #quickly verify the deployments

$ kubectl get crd # new entry for deletetablepartitiondatajobs.ddbctl.operators.jittakal.io

$ kubectl get namespaces # new entry for ddbctl-dtp-operator-system

$ kubectl get pods -n ddbctl-dtp-operator-system

$ kubctl logs -f <<pod-name-from-above-command>> -n ddbctl-dtp-operator-system
```

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

- Customer Resource for DeleteTablePartitionDataJob CRD

```bash
$ kubectl create -f delete_partition_data_tablename_job.yaml

$ # Sample
$ kubectl create -f config/samples/ddbctl_v1alpha1_deletepartitiondatajob.yaml
```

- View Delete Table Partition Data Job - Summary

```bash
$ kubectl get jobs # default - namespace

$ kubectl get pods

$ kubectl logs <<podname-of-delete-table-partition-data-job>> 

$ # verify log table name and number of records delete 
$ # delete summary report on job completion
```

## Open Issues

- Limit the operator's ability to create Jobs exclusively within the "default" namespace, and adjust the RBAC permissions of the controller manager accordingly.

## ToDo

- Automatically remove successfully completed jobs after 5 or 30 minutes.
- Automatically delete failed jobs older than the last 5 instances.

## Reference

- [Docker image - operator controller manager](https://hub.docker.com/repository/docker/jittakal/ddbctl-dtp-operator/general)
- [Docker image - job](https://hub.docker.com/repository/docker/jittakal/go-dynamodb-partition-delete/general)


## License

Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.