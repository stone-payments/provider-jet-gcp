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

type SubscriptionIamMemberConditionObservation struct {
}

type SubscriptionIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type SubscriptionIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type SubscriptionIamMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []SubscriptionIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	Subscription *string `json:"subscription" tf:"subscription,omitempty"`
}

// SubscriptionIamMemberSpec defines the desired state of SubscriptionIamMember
type SubscriptionIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionIamMemberParameters `json:"forProvider"`
}

// SubscriptionIamMemberStatus defines the observed state of SubscriptionIamMember.
type SubscriptionIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionIamMember is the Schema for the SubscriptionIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type SubscriptionIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SubscriptionIamMemberSpec   `json:"spec"`
	Status            SubscriptionIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionIamMemberList contains a list of SubscriptionIamMembers
type SubscriptionIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SubscriptionIamMember `json:"items"`
}

// Repository type metadata.
var (
	SubscriptionIamMember_Kind             = "SubscriptionIamMember"
	SubscriptionIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: SubscriptionIamMember_Kind}.String()
	SubscriptionIamMember_KindAPIVersion   = SubscriptionIamMember_Kind + "." + CRDGroupVersion.String()
	SubscriptionIamMember_GroupVersionKind = CRDGroupVersion.WithKind(SubscriptionIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&SubscriptionIamMember{}, &SubscriptionIamMemberList{})
}
