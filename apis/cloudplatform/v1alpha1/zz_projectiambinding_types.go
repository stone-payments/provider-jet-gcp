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

type ProjectIAMBindingConditionObservation struct {
}

type ProjectIAMBindingConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type ProjectIAMBindingObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type ProjectIAMBindingParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ProjectIAMBindingConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Members []*string `json:"members" tf:"members,omitempty"`

	// +kubebuilder:validation:Required
	Project *string `json:"project" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// ProjectIAMBindingSpec defines the desired state of ProjectIAMBinding
type ProjectIAMBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectIAMBindingParameters `json:"forProvider"`
}

// ProjectIAMBindingStatus defines the observed state of ProjectIAMBinding.
type ProjectIAMBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectIAMBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMBinding is the Schema for the ProjectIAMBindings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ProjectIAMBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectIAMBindingSpec   `json:"spec"`
	Status            ProjectIAMBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectIAMBindingList contains a list of ProjectIAMBindings
type ProjectIAMBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectIAMBinding `json:"items"`
}

// Repository type metadata.
var (
	ProjectIAMBinding_Kind             = "ProjectIAMBinding"
	ProjectIAMBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectIAMBinding_Kind}.String()
	ProjectIAMBinding_KindAPIVersion   = ProjectIAMBinding_Kind + "." + CRDGroupVersion.String()
	ProjectIAMBinding_GroupVersionKind = CRDGroupVersion.WithKind(ProjectIAMBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectIAMBinding{}, &ProjectIAMBindingList{})
}
