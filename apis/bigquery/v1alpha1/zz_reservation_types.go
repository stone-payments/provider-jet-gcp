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

type ReservationObservation struct {
}

type ReservationParameters struct {

	// If false, any query using this reservation will use idle slots from other reservations within
	// the same admin project. If true, a query using this reservation will execute with the slot
	// capacity specified above at most.
	// +kubebuilder:validation:Optional
	IgnoreIdleSlots *bool `json:"ignoreIdleSlots,omitempty" tf:"ignore_idle_slots,omitempty"`

	// The geographic location where the transfer config should reside.
	// Examples: US, EU, asia-northeast1. The default value is US.
	// +kubebuilder:validation:Optional
	Location *string `json:"location,omitempty" tf:"location,omitempty"`

	// The name of the reservation. This field must only contain alphanumeric characters or dash.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Project *string `json:"project,omitempty" tf:"project,omitempty"`

	// Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the
	// unit of parallelism. Queries using this reservation might use more slots during runtime if ignoreIdleSlots is set to false.
	// +kubebuilder:validation:Required
	SlotCapacity *int64 `json:"slotCapacity" tf:"slot_capacity,omitempty"`
}

// ReservationSpec defines the desired state of Reservation
type ReservationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ReservationParameters `json:"forProvider"`
}

// ReservationStatus defines the observed state of Reservation.
type ReservationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ReservationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Reservation is the Schema for the Reservations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type Reservation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ReservationSpec   `json:"spec"`
	Status            ReservationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ReservationList contains a list of Reservations
type ReservationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Reservation `json:"items"`
}

// Repository type metadata.
var (
	Reservation_Kind             = "Reservation"
	Reservation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Reservation_Kind}.String()
	Reservation_KindAPIVersion   = Reservation_Kind + "." + CRDGroupVersion.String()
	Reservation_GroupVersionKind = CRDGroupVersion.WithKind(Reservation_Kind)
)

func init() {
	SchemeBuilder.Register(&Reservation{}, &ReservationList{})
}
