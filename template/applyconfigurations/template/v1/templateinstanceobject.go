// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// TemplateInstanceObjectApplyConfiguration represents an declarative configuration of the TemplateInstanceObject type for use
// with apply.
type TemplateInstanceObjectApplyConfiguration struct {
	Ref *v1.ObjectReferenceApplyConfiguration `json:"ref,omitempty"`
}

// TemplateInstanceObjectApplyConfiguration constructs an declarative configuration of the TemplateInstanceObject type for use with
// apply.
func TemplateInstanceObject() *TemplateInstanceObjectApplyConfiguration {
	return &TemplateInstanceObjectApplyConfiguration{}
}

// WithRef sets the Ref field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Ref field is set to the value of the last call.
func (b *TemplateInstanceObjectApplyConfiguration) WithRef(value *v1.ObjectReferenceApplyConfiguration) *TemplateInstanceObjectApplyConfiguration {
	b.Ref = value
	return b
}
