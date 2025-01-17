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

type ApplicationUrlDispatchRulesObservation struct {
}

type ApplicationUrlDispatchRulesParameters struct {

	// Rules to match an HTTP request and dispatch that request to a service.
	// +kubebuilder:validation:Required
	DispatchRules []DispatchRulesParameters `json:"dispatchRules" tf:"dispatch_rules,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`
}

type DispatchRulesObservation struct {
}

type DispatchRulesParameters struct {

	// Domain name to match against. The wildcard "*" is supported if specified before a period: "*.".
	// Defaults to matching all domains: "*".
	// +kubebuilder:validation:Optional
	Domain *string `json:"domain,omitempty" tf:"domain,omitempty"`

	// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	// +kubebuilder:validation:Required
	Path *string `json:"path" tf:"path,omitempty"`

	// Pathname within the host. Must start with a "/". A single "*" can be included at the end of the path.
	// The sum of the lengths of the domain and path may not exceed 100 characters.
	// +kubebuilder:validation:Required
	Service *string `json:"service" tf:"service,omitempty"`
}

// ApplicationUrlDispatchRulesSpec defines the desired state of ApplicationUrlDispatchRules
type ApplicationUrlDispatchRulesSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationUrlDispatchRulesParameters `json:"forProvider"`
}

// ApplicationUrlDispatchRulesStatus defines the observed state of ApplicationUrlDispatchRules.
type ApplicationUrlDispatchRulesStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationUrlDispatchRulesObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationUrlDispatchRules is the Schema for the ApplicationUrlDispatchRuless API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type ApplicationUrlDispatchRules struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationUrlDispatchRulesSpec   `json:"spec"`
	Status            ApplicationUrlDispatchRulesStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationUrlDispatchRulesList contains a list of ApplicationUrlDispatchRuless
type ApplicationUrlDispatchRulesList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ApplicationUrlDispatchRules `json:"items"`
}

// Repository type metadata.
var (
	ApplicationUrlDispatchRules_Kind             = "ApplicationUrlDispatchRules"
	ApplicationUrlDispatchRules_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ApplicationUrlDispatchRules_Kind}.String()
	ApplicationUrlDispatchRules_KindAPIVersion   = ApplicationUrlDispatchRules_Kind + "." + CRDGroupVersion.String()
	ApplicationUrlDispatchRules_GroupVersionKind = CRDGroupVersion.WithKind(ApplicationUrlDispatchRules_Kind)
)

func init() {
	SchemeBuilder.Register(&ApplicationUrlDispatchRules{}, &ApplicationUrlDispatchRulesList{})
}
