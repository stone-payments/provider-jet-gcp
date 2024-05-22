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

type HmacKeyObservation struct {
	AccessID *string `json:"accessId,omitempty" tf:"access_id,omitempty"`

	TimeCreated *string `json:"timeCreated,omitempty" tf:"time_created,omitempty"`

	Updated *string `json:"updated,omitempty" tf:"updated,omitempty"`
}

type HmacKeyParameters struct {

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The email address of the key's associated service account.
	// +kubebuilder:validation:Required
	ServiceAccountEmail *string `json:"serviceAccountEmail" tf:"service_account_email,omitempty"`

	// The state of the key. Can be set to one of ACTIVE, INACTIVE. Default value: "ACTIVE" Possible values: ["ACTIVE", "INACTIVE"]
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

// HmacKeySpec defines the desired state of HmacKey
type HmacKeySpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HmacKeyParameters `json:"forProvider"`
}

// HmacKeyStatus defines the observed state of HmacKey.
type HmacKeyStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HmacKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// HmacKey is the Schema for the HmacKeys API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type HmacKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HmacKeySpec   `json:"spec"`
	Status            HmacKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HmacKeyList contains a list of HmacKeys
type HmacKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HmacKey `json:"items"`
}

// Repository type metadata.
var (
	HmacKey_Kind             = "HmacKey"
	HmacKey_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HmacKey_Kind}.String()
	HmacKey_KindAPIVersion   = HmacKey_Kind + "." + CRDGroupVersion.String()
	HmacKey_GroupVersionKind = CRDGroupVersion.WithKind(HmacKey_Kind)
)

func init() {
	SchemeBuilder.Register(&HmacKey{}, &HmacKeyList{})
}