/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AiDatasetObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	UpdateTime *string `json:"updateTime,omitempty" tf:"update_time,omitempty"`
}

type AiDatasetParameters struct {

	// The user-defined name of the Dataset. The name can be up to 128 characters long and can be consist of any UTF-8 characters.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// Customer-managed encryption key spec for a Dataset. If set, this Dataset and all sub-resources of this Dataset will be secured by this key.
	// +kubebuilder:validation:Optional
	EncryptionSpec []EncryptionSpecParameters `json:"encryptionSpec,omitempty" tf:"encryption_spec,omitempty"`

	// A set of key/value label pairs to assign to this Workflow.
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// Points to a YAML file stored on Google Cloud Storage describing additional information about the Dataset. The schema is defined as an OpenAPI 3.0.2 Schema Object. The schema files that can be used here are found in gs://google-cloud-aiplatform/schema/dataset/metadata/.
	// +kubebuilder:validation:Required
	MetadataSchemaURI *string `json:"metadataSchemaUri" tf:"metadata_schema_uri,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The region of the dataset. eg us-central1
	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`
}

type EncryptionSpecObservation struct {
}

type EncryptionSpecParameters struct {

	// Required. The Cloud KMS resource identifier of the customer managed encryption key used to protect a resource.
	// Has the form: projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key. The key needs to be in the same region as where the resource is created.
	// +kubebuilder:validation:Optional
	KmsKeyName *string `json:"kmsKeyName,omitempty" tf:"kms_key_name,omitempty"`
}

// AiDatasetSpec defines the desired state of AiDataset
type AiDatasetSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AiDatasetParameters `json:"forProvider"`
}

// AiDatasetStatus defines the observed state of AiDataset.
type AiDatasetStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AiDatasetObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// AiDataset is the Schema for the AiDatasets API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type AiDataset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AiDatasetSpec   `json:"spec"`
	Status            AiDatasetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AiDatasetList contains a list of AiDatasets
type AiDatasetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AiDataset `json:"items"`
}

// Repository type metadata.
var (
	AiDataset_Kind             = "AiDataset"
	AiDataset_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: AiDataset_Kind}.String()
	AiDataset_KindAPIVersion   = AiDataset_Kind + "." + CRDGroupVersion.String()
	AiDataset_GroupVersionKind = CRDGroupVersion.WithKind(AiDataset_Kind)
)

func init() {
	SchemeBuilder.Register(&AiDataset{}, &AiDatasetList{})
}
