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

type FhirStoreIamBindingConditionObservation struct {
}

type FhirStoreIamBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type FhirStoreIamBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type FhirStoreIamBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []FhirStoreIamBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	FhirStoreID *string `json:"fhirStoreId" tf:"fhir_store_id,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// FhirStoreIamBindingSpec defines the desired state of FhirStoreIamBinding
type FhirStoreIamBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FhirStoreIamBindingParameters `json:"forProvider"`
}

// FhirStoreIamBindingStatus defines the observed state of FhirStoreIamBinding.
type FhirStoreIamBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FhirStoreIamBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FhirStoreIamBinding is the Schema for the FhirStoreIamBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type FhirStoreIamBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FhirStoreIamBindingSpec   `json:"spec"`
	Status            FhirStoreIamBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FhirStoreIamBindingList contains a list of FhirStoreIamBindings
type FhirStoreIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FhirStoreIamBinding `json:"items"`
}

// Repository type metadata.
var (
	FhirStoreIamBinding_Kind             = "FhirStoreIamBinding"
	FhirStoreIamBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FhirStoreIamBinding_Kind}.String()
	FhirStoreIamBinding_KindAPIVersion   = FhirStoreIamBinding_Kind + "." + CRDGroupVersion.String()
	FhirStoreIamBinding_GroupVersionKind = CRDGroupVersion.WithKind(FhirStoreIamBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&FhirStoreIamBinding{}, &FhirStoreIamBindingList{})
}
