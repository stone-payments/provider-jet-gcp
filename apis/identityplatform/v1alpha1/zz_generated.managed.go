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
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this DefaultSupportedIdpConfig.
func (mg *DefaultSupportedIdpConfig) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this DefaultSupportedIdpConfig.
func (mg *DefaultSupportedIdpConfig) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this DefaultSupportedIdpConfig.
func (mg *DefaultSupportedIdpConfig) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this DefaultSupportedIdpConfig.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *DefaultSupportedIdpConfig) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this DefaultSupportedIdpConfig.
func (mg *DefaultSupportedIdpConfig) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this DefaultSupportedIdpConfig.
func (mg *DefaultSupportedIdpConfig) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this DefaultSupportedIdpConfig.
func (mg *DefaultSupportedIdpConfig) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this DefaultSupportedIdpConfig.
func (mg *DefaultSupportedIdpConfig) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this DefaultSupportedIdpConfig.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *DefaultSupportedIdpConfig) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this DefaultSupportedIdpConfig.
func (mg *DefaultSupportedIdpConfig) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this InboundSamlConfig.
func (mg *InboundSamlConfig) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this InboundSamlConfig.
func (mg *InboundSamlConfig) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this InboundSamlConfig.
func (mg *InboundSamlConfig) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this InboundSamlConfig.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *InboundSamlConfig) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this InboundSamlConfig.
func (mg *InboundSamlConfig) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this InboundSamlConfig.
func (mg *InboundSamlConfig) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this InboundSamlConfig.
func (mg *InboundSamlConfig) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this InboundSamlConfig.
func (mg *InboundSamlConfig) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this InboundSamlConfig.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *InboundSamlConfig) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this InboundSamlConfig.
func (mg *InboundSamlConfig) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this OauthIdpConfig.
func (mg *OauthIdpConfig) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this OauthIdpConfig.
func (mg *OauthIdpConfig) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this OauthIdpConfig.
func (mg *OauthIdpConfig) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this OauthIdpConfig.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *OauthIdpConfig) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this OauthIdpConfig.
func (mg *OauthIdpConfig) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this OauthIdpConfig.
func (mg *OauthIdpConfig) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this OauthIdpConfig.
func (mg *OauthIdpConfig) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this OauthIdpConfig.
func (mg *OauthIdpConfig) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this OauthIdpConfig.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *OauthIdpConfig) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this OauthIdpConfig.
func (mg *OauthIdpConfig) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this Tenant.
func (mg *Tenant) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this Tenant.
func (mg *Tenant) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this Tenant.
func (mg *Tenant) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this Tenant.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *Tenant) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this Tenant.
func (mg *Tenant) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this Tenant.
func (mg *Tenant) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this Tenant.
func (mg *Tenant) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this Tenant.
func (mg *Tenant) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this Tenant.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *Tenant) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this Tenant.
func (mg *Tenant) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this TenantDefaultSupportedIdpConfig.
func (mg *TenantDefaultSupportedIdpConfig) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TenantDefaultSupportedIdpConfig.
func (mg *TenantDefaultSupportedIdpConfig) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this TenantDefaultSupportedIdpConfig.
func (mg *TenantDefaultSupportedIdpConfig) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this TenantDefaultSupportedIdpConfig.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *TenantDefaultSupportedIdpConfig) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this TenantDefaultSupportedIdpConfig.
func (mg *TenantDefaultSupportedIdpConfig) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this TenantDefaultSupportedIdpConfig.
func (mg *TenantDefaultSupportedIdpConfig) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TenantDefaultSupportedIdpConfig.
func (mg *TenantDefaultSupportedIdpConfig) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this TenantDefaultSupportedIdpConfig.
func (mg *TenantDefaultSupportedIdpConfig) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this TenantDefaultSupportedIdpConfig.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *TenantDefaultSupportedIdpConfig) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this TenantDefaultSupportedIdpConfig.
func (mg *TenantDefaultSupportedIdpConfig) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this TenantInboundSamlConfig.
func (mg *TenantInboundSamlConfig) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TenantInboundSamlConfig.
func (mg *TenantInboundSamlConfig) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this TenantInboundSamlConfig.
func (mg *TenantInboundSamlConfig) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this TenantInboundSamlConfig.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *TenantInboundSamlConfig) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this TenantInboundSamlConfig.
func (mg *TenantInboundSamlConfig) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this TenantInboundSamlConfig.
func (mg *TenantInboundSamlConfig) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TenantInboundSamlConfig.
func (mg *TenantInboundSamlConfig) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this TenantInboundSamlConfig.
func (mg *TenantInboundSamlConfig) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this TenantInboundSamlConfig.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *TenantInboundSamlConfig) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this TenantInboundSamlConfig.
func (mg *TenantInboundSamlConfig) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}

// GetCondition of this TenantOauthIdpConfig.
func (mg *TenantOauthIdpConfig) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this TenantOauthIdpConfig.
func (mg *TenantOauthIdpConfig) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetProviderConfigReference of this TenantOauthIdpConfig.
func (mg *TenantOauthIdpConfig) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this TenantOauthIdpConfig.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *TenantOauthIdpConfig) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetWriteConnectionSecretToReference of this TenantOauthIdpConfig.
func (mg *TenantOauthIdpConfig) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this TenantOauthIdpConfig.
func (mg *TenantOauthIdpConfig) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this TenantOauthIdpConfig.
func (mg *TenantOauthIdpConfig) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetProviderConfigReference of this TenantOauthIdpConfig.
func (mg *TenantOauthIdpConfig) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this TenantOauthIdpConfig.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *TenantOauthIdpConfig) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetWriteConnectionSecretToReference of this TenantOauthIdpConfig.
func (mg *TenantOauthIdpConfig) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
