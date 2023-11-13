## Steps for reference

- Kubebuilder Installation

Ensure that Kubebuilder and its related dependencies are installed on your local machine. You can follow the Kubebuilder installation guide provided in the official documentation to set up the necessary tooling for building Kubernetes operators.

```bash
$ brew install kubebulder

$ brew install kustomize
```

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