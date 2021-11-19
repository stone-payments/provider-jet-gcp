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

type ConditionsObservation struct {
	Message *string `json:"message,omitempty" tf:"message,omitempty"`

	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type ConditionsParameters struct {
}

type DomainMappingObservation struct {
	Status []StatusObservation `json:"status,omitempty" tf:"status,omitempty"`
}

type DomainMappingParameters struct {

	// The location of the cloud run instance. eg us-central1
	// +kubebuilder:validation:Required
	Location *string `json:"location" tf:"location,omitempty"`

	// Metadata associated with this DomainMapping.
	// +kubebuilder:validation:Required
	Metadata []MetadataParameters `json:"metadata" tf:"metadata,omitempty"`

	// Name should be a [verified](https://support.google.com/webmasters/answer/9008080) domain
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// The spec for this DomainMapping.
	// +kubebuilder:validation:Required
	Spec []SpecParameters `json:"spec" tf:"spec,omitempty"`
}

type MetadataObservation struct {
	Generation *int64 `json:"generation,omitempty" tf:"generation,omitempty"`

	ResourceVersion *string `json:"resourceVersion,omitempty" tf:"resource_version,omitempty"`

	SelfLink *string `json:"selfLink,omitempty" tf:"self_link,omitempty"`

	UID *string `json:"uid,omitempty" tf:"uid,omitempty"`
}

type MetadataParameters struct {

	// Annotations is a key value map stored with a resource that
	// may be set by external tools to store and retrieve arbitrary metadata. More
	// info: http://kubernetes.io/docs/user-guide/annotations
	//
	// **Note**: The Cloud Run API may add additional annotations that were not provided in your config.
	// If terraform plan shows a diff where a server-side annotation is added, you can add it to your config
	// or apply the lifecycle.ignore_changes rule to the metadata.0.annotations field.
	// +kubebuilder:validation:Optional
	Annotations map[string]*string `json:"annotations,omitempty" tf:"annotations,omitempty"`

	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and routes.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +kubebuilder:validation:Optional
	Labels map[string]*string `json:"labels,omitempty" tf:"labels,omitempty"`

	// In Cloud Run the namespace must be equal to either the
	// project ID or project number.
	// +kubebuilder:validation:Required
	Namespace *string `json:"namespace" tf:"namespace,omitempty"`
}

type ResourceRecordsObservation struct {
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	Rrdata *string `json:"rrdata,omitempty" tf:"rrdata,omitempty"`
}

type ResourceRecordsParameters struct {

	// Resource record type. Example: 'AAAA'. Possible values: ["A", "AAAA", "CNAME"]
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type SpecObservation struct {
}

type SpecParameters struct {

	// The mode of the certificate. Default value: "AUTOMATIC" Possible values: ["NONE", "AUTOMATIC"]
	// +kubebuilder:validation:Optional
	CertificateMode *string `json:"certificateMode,omitempty" tf:"certificate_mode,omitempty"`

	// If set, the mapping will override any mapping set before this spec was set.
	// It is recommended that the user leaves this empty to receive an error
	// warning about a potential conflict and only set it once the respective UI
	// has given such a warning.
	// +kubebuilder:validation:Optional
	ForceOverride *bool `json:"forceOverride,omitempty" tf:"force_override,omitempty"`

	// The name of the Cloud Run Service that this DomainMapping applies to.
	// The route must exist.
	// +kubebuilder:validation:Required
	RouteName *string `json:"routeName" tf:"route_name,omitempty"`
}

type StatusObservation struct {
	Conditions []ConditionsObservation `json:"conditions,omitempty" tf:"conditions,omitempty"`

	MappedRouteName *string `json:"mappedRouteName,omitempty" tf:"mapped_route_name,omitempty"`

	ObservedGeneration *int64 `json:"observedGeneration,omitempty" tf:"observed_generation,omitempty"`
}

type StatusParameters struct {

	// The resource records required to configure this domain mapping. These
	// records must be added to the domain's DNS configuration in order to
	// serve the application via this domain mapping.
	// +kubebuilder:validation:Optional
	ResourceRecords []ResourceRecordsParameters `json:"resourceRecords,omitempty" tf:"resource_records,omitempty"`
}

// DomainMappingSpec defines the desired state of DomainMapping
type DomainMappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DomainMappingParameters `json:"forProvider"`
}

// DomainMappingStatus defines the observed state of DomainMapping.
type DomainMappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DomainMappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DomainMapping is the Schema for the DomainMappings API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type DomainMapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainMappingSpec   `json:"spec"`
	Status            DomainMappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DomainMappingList contains a list of DomainMappings
type DomainMappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainMapping `json:"items"`
}

// Repository type metadata.
var (
	DomainMapping_Kind             = "DomainMapping"
	DomainMapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DomainMapping_Kind}.String()
	DomainMapping_KindAPIVersion   = DomainMapping_Kind + "." + CRDGroupVersion.String()
	DomainMapping_GroupVersionKind = CRDGroupVersion.WithKind(DomainMapping_Kind)
)

func init() {
	SchemeBuilder.Register(&DomainMapping{}, &DomainMappingList{})
}
