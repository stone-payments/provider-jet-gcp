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

type CxVersionNluSettingsObservation struct {
}

type CxVersionNluSettingsParameters struct {

	// To filter out false positive results and still get variety in matched natural language inputs for your agent, you can tune the machine learning classification threshold. If the returned score value is less than the threshold value, then a no-match event will be triggered.
	// The score values range from 0.0 (completely uncertain) to 1.0 (completely certain). If set to 0.0, the default of 0.3 is used.
	// +kubebuilder:validation:Optional
	ClassificationThreshold *float64 `json:"classificationThreshold,omitempty" tf:"classification_threshold,omitempty"`

	// Indicates NLU model training mode.
	// * MODEL_TRAINING_MODE_AUTOMATIC: NLU model training is automatically triggered when a flow gets modified. User can also manually trigger model training in this mode.
	// * MODEL_TRAINING_MODE_MANUAL: User needs to manually trigger NLU model training. Best for large flows whose models take long time to train. Possible values: ["MODEL_TRAINING_MODE_AUTOMATIC", "MODEL_TRAINING_MODE_MANUAL"]
	// +kubebuilder:validation:Optional
	ModelTrainingMode *string `json:"modelTrainingMode,omitempty" tf:"model_training_mode,omitempty"`

	// Indicates the type of NLU model.
	// * MODEL_TYPE_STANDARD: Use standard NLU model.
	// * MODEL_TYPE_ADVANCED: Use advanced NLU model. Possible values: ["MODEL_TYPE_STANDARD", "MODEL_TYPE_ADVANCED"]
	// +kubebuilder:validation:Optional
	ModelType *string `json:"modelType,omitempty" tf:"model_type,omitempty"`
}

type CxVersionObservation struct {
	CreateTime *string `json:"createTime,omitempty" tf:"create_time,omitempty"`

	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	NluSettings []CxVersionNluSettingsObservation `json:"nluSettings,omitempty" tf:"nlu_settings,omitempty"`

	State *string `json:"state,omitempty" tf:"state,omitempty"`
}

type CxVersionParameters struct {

	// The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The human-readable name of the version. Limit of 64 characters.
	// +kubebuilder:validation:Required
	DisplayName *string `json:"displayName" tf:"display_name,omitempty"`

	// The Flow to create an Version for.
	// Format: projects/<Project ID>/locations/<Location ID>/agents/<Agent ID>/flows/<Flow ID>.
	// +kubebuilder:validation:Optional
	Parent *string `json:"parent,omitempty" tf:"parent,omitempty"`
}

// CxVersionSpec defines the desired state of CxVersion
type CxVersionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CxVersionParameters `json:"forProvider"`
}

// CxVersionStatus defines the observed state of CxVersion.
type CxVersionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CxVersionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// CxVersion is the Schema for the CxVersions API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gcpjet}
type CxVersion struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CxVersionSpec   `json:"spec"`
	Status            CxVersionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CxVersionList contains a list of CxVersions
type CxVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CxVersion `json:"items"`
}

// Repository type metadata.
var (
	CxVersion_Kind             = "CxVersion"
	CxVersion_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: CxVersion_Kind}.String()
	CxVersion_KindAPIVersion   = CxVersion_Kind + "." + CRDGroupVersion.String()
	CxVersion_GroupVersionKind = CRDGroupVersion.WithKind(CxVersion_Kind)
)

func init() {
	SchemeBuilder.Register(&CxVersion{}, &CxVersionList{})
}
