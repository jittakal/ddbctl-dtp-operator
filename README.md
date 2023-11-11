# ddbctl-dtp-operator

DynamoDB Delete Table Partition Kubernetes Operator (Golang) using Kubebuilder framework


```bash
$ cd ~\workspace\git.ws
$ git clone https://github.com/jittakal/ddbctl-dtp-operator
$ cd ddbctl-dtp-operator
```

Initialize the Go Moudle

```bash
$ go mod init github.com/jittakal/ddbctl-dtp-operator
```

Kubebuilder init

```bash
$ kubebuilder init --domain operators.jittakal.io --repo github.com/jittakal/ddbctl-dtp-operator
```

Custom Resource Defination - API

```bash
$ kubebuilder create api --group ddbctl --version v1alpha1 --kind DeleteTablePartitionDataJob
```

