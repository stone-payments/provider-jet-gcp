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

type DiskIamMemberConditionObservation struct {
}

type DiskIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type DiskIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type DiskIamMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []DiskIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Zone *string `json:"zone,omitempty" tf:"zone,omitempty"`
}

// DiskIamMemberSpec defines the desired state of DiskIamMember
type DiskIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DiskIamMemberParameters `json:"forProvider"`
}

// DiskIamMemberStatus defines the observed state of DiskIamMember.
type DiskIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DiskIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DiskIamMember is the Schema for the DiskIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type DiskIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskIamMemberSpec   `json:"spec"`
	Status            DiskIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DiskIamMemberList contains a list of DiskIamMembers
type DiskIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskIamMember `json:"items"`
}

// Repository type metadata.
var (
	DiskIamMember_Kind             = "DiskIamMember"
	DiskIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DiskIamMember_Kind}.String()
	DiskIamMember_KindAPIVersion   = DiskIamMember_Kind + "." + CRDGroupVersion.String()
	DiskIamMember_GroupVersionKind = CRDGroupVersion.WithKind(DiskIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&DiskIamMember{}, &DiskIamMemberList{})
}