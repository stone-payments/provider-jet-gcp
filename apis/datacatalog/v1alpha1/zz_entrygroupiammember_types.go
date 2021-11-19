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

type EntryGroupIamMemberConditionObservation struct {
}

type EntryGroupIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type EntryGroupIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type EntryGroupIamMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []EntryGroupIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	EntryGroup *string `json:"entryGroup" tf:"entry_group,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Optional
	Region *string `json:"region,omitempty" tf:"region,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`
}

// EntryGroupIamMemberSpec defines the desired state of EntryGroupIamMember
type EntryGroupIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     EntryGroupIamMemberParameters `json:"forProvider"`
}

// EntryGroupIamMemberStatus defines the observed state of EntryGroupIamMember.
type EntryGroupIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        EntryGroupIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// EntryGroupIamMember is the Schema for the EntryGroupIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type EntryGroupIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EntryGroupIamMemberSpec   `json:"spec"`
	Status            EntryGroupIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EntryGroupIamMemberList contains a list of EntryGroupIamMembers
type EntryGroupIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EntryGroupIamMember `json:"items"`
}

// Repository type metadata.
var (
	EntryGroupIamMember_Kind             = "EntryGroupIamMember"
	EntryGroupIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: EntryGroupIamMember_Kind}.String()
	EntryGroupIamMember_KindAPIVersion   = EntryGroupIamMember_Kind + "." + CRDGroupVersion.String()
	EntryGroupIamMember_GroupVersionKind = CRDGroupVersion.WithKind(EntryGroupIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&EntryGroupIamMember{}, &EntryGroupIamMemberList{})
}
