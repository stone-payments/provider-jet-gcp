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

type ServiceIamMemberConditionObservation struct {
}

type ServiceIamMemberConditionParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Required
	Expression *string `json:"expression" tf:"expression,omitempty"`

	// +kubebuilder:validation:Required
	Title *string `json:"title" tf:"title,omitempty"`
}

type ServiceIamMemberObservation struct {
	Etag *string `json:"etag,omitempty" tf:"etag,omitempty"`
}

type ServiceIamMemberParameters struct {

	// +kubebuilder:validation:Optional
	Condition []ServiceIamMemberConditionParameters `json:"condition,omitempty" tf:"condition,omitempty"`

	// +kubebuilder:validation:Required
	Member *string `json:"member" tf:"member,omitempty"`

	// +kubebuilder:validation:Required
	Role *string `json:"role" tf:"role,omitempty"`

	// +kubebuilder:validation:Required
	ServiceName *string `json:"serviceName" tf:"service_name,omitempty"`
}

// ServiceIamMemberSpec defines the desired state of ServiceIamMember
type ServiceIamMemberSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceIamMemberParameters `json:"forProvider"`
}

// ServiceIamMemberStatus defines the observed state of ServiceIamMember.
type ServiceIamMemberStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceIamMemberObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceIamMember is the Schema for the ServiceIamMembers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ServiceIamMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceIamMemberSpec   `json:"spec"`
	Status            ServiceIamMemberStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceIamMemberList contains a list of ServiceIamMembers
type ServiceIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServiceIamMember `json:"items"`
}

// Repository type metadata.
var (
	ServiceIamMember_Kind             = "ServiceIamMember"
	ServiceIamMember_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServiceIamMember_Kind}.String()
	ServiceIamMember_KindAPIVersion   = ServiceIamMember_Kind + "." + CRDGroupVersion.String()
	ServiceIamMember_GroupVersionKind = CRDGroupVersion.WithKind(ServiceIamMember_Kind)
)

func init() {
	SchemeBuilder.Register(&ServiceIamMember{}, &ServiceIamMemberList{})
}
