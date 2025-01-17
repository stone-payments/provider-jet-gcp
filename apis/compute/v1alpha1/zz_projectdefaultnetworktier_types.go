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

type ProjectDefaultNetworkTierObservation struct {
}

type ProjectDefaultNetworkTierParameters struct {

	// The default network tier to be configured for the project. This field can take the following values: PREMIUM or STANDARD.
	// +kubebuilder:validation:Required
	NetworkTier *string `json:"networkTier" tf:"network_tier,omitempty"`

	// The ID of the project in which the resource belongs. If it is not provided, the provider project is used.
	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

// ProjectDefaultNetworkTierSpec defines the desired state of ProjectDefaultNetworkTier
type ProjectDefaultNetworkTierSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ProjectDefaultNetworkTierParameters `json:"forProvider"`
}

// ProjectDefaultNetworkTierStatus defines the observed state of ProjectDefaultNetworkTier.
type ProjectDefaultNetworkTierStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ProjectDefaultNetworkTierObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDefaultNetworkTier is the Schema for the ProjectDefaultNetworkTiers API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ProjectDefaultNetworkTier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectDefaultNetworkTierSpec   `json:"spec"`
	Status            ProjectDefaultNetworkTierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ProjectDefaultNetworkTierList contains a list of ProjectDefaultNetworkTiers
type ProjectDefaultNetworkTierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectDefaultNetworkTier `json:"items"`
}

// Repository type metadata.
var (
	ProjectDefaultNetworkTier_Kind             = "ProjectDefaultNetworkTier"
	ProjectDefaultNetworkTier_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ProjectDefaultNetworkTier_Kind}.String()
	ProjectDefaultNetworkTier_KindAPIVersion   = ProjectDefaultNetworkTier_Kind + "." + CRDGroupVersion.String()
	ProjectDefaultNetworkTier_GroupVersionKind = CRDGroupVersion.WithKind(ProjectDefaultNetworkTier_Kind)
)

func init() {
	SchemeBuilder.Register(&ProjectDefaultNetworkTier{}, &ProjectDefaultNetworkTierList{})
}
