/*
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
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// DeleteTablePartitionDataJobSpec defines the desired state of DeleteTablePartitionDataJob
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

// DeleteTablePartitionDataJobStatus defines the observed state of DeleteTablePartitionDataJob
type DeleteTablePartitionDataJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// DeleteTablePartitionDataJob is the Schema for the deletetablepartitiondatajobs API
type DeleteTablePartitionDataJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeleteTablePartitionDataJobSpec   `json:"spec,omitempty"`
	Status DeleteTablePartitionDataJobStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// DeleteTablePartitionDataJobList contains a list of DeleteTablePartitionDataJob
type DeleteTablePartitionDataJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeleteTablePartitionDataJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DeleteTablePartitionDataJob{}, &DeleteTablePartitionDataJobList{})
}
