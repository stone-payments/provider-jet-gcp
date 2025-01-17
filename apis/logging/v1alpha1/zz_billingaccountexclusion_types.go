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

type BillingAccountExclusionObservation struct {
}

type BillingAccountExclusionParameters struct {

	// +kubebuilder:validation:Required
	BillingAccount *string `json:"billingAccount" tf:"billing_account,omitempty"`

	// A human-readable description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether this exclusion rule should be disabled or not. This defaults to false.
	// +kubebuilder:validation:Optional
	Disabled *bool `json:"disabled,omitempty" tf:"disabled,omitempty"`

	// The filter to apply when excluding logs. Only log entries that match the filter are excluded.
	// +kubebuilder:validation:Required
	Filter *string `json:"filter" tf:"filter,omitempty"`

	// The name of the logging exclusion.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

// BillingAccountExclusionSpec defines the desired state of BillingAccountExclusion
type BillingAccountExclusionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     BillingAccountExclusionParameters `json:"forProvider"`
}

// BillingAccountExclusionStatus defines the observed state of BillingAccountExclusion.
type BillingAccountExclusionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        BillingAccountExclusionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// BillingAccountExclusion is the Schema for the BillingAccountExclusions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type BillingAccountExclusion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BillingAccountExclusionSpec   `json:"spec"`
	Status            BillingAccountExclusionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BillingAccountExclusionList contains a list of BillingAccountExclusions
type BillingAccountExclusionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BillingAccountExclusion `json:"items"`
}

// Repository type metadata.
var (
	BillingAccountExclusion_Kind             = "BillingAccountExclusion"
	BillingAccountExclusion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: BillingAccountExclusion_Kind}.String()
	BillingAccountExclusion_KindAPIVersion   = BillingAccountExclusion_Kind + "." + CRDGroupVersion.String()
	BillingAccountExclusion_GroupVersionKind = CRDGroupVersion.WithKind(BillingAccountExclusion_Kind)
)

func init() {
	SchemeBuilder.Register(&BillingAccountExclusion{}, &BillingAccountExclusionList{})
}
