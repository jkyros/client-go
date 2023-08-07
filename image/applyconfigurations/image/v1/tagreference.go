// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// TagReferenceApplyConfiguration represents an declarative configuration of the TagReference type for use
// with apply.
type TagReferenceApplyConfiguration struct {
	Name            *string                               `json:"name,omitempty"`
	Annotations     map[string]string                     `json:"annotations,omitempty"`
	From            *v1.ObjectReferenceApplyConfiguration `json:"from,omitempty"`
	Reference       *bool                                 `json:"reference,omitempty"`
	Generation      *int64                                `json:"generation,omitempty"`
	ImportPolicy    *TagImportPolicyApplyConfiguration    `json:"importPolicy,omitempty"`
	ReferencePolicy *TagReferencePolicyApplyConfiguration `json:"referencePolicy,omitempty"`
}

// TagReferenceApplyConfiguration constructs an declarative configuration of the TagReference type for use with
// apply.
func TagReference() *TagReferenceApplyConfiguration {
	return &TagReferenceApplyConfiguration{}
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *TagReferenceApplyConfiguration) WithName(value string) *TagReferenceApplyConfiguration {
	b.Name = &value
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *TagReferenceApplyConfiguration) WithAnnotations(entries map[string]string) *TagReferenceApplyConfiguration {
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithFrom sets the From field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the From field is set to the value of the last call.
func (b *TagReferenceApplyConfiguration) WithFrom(value *v1.ObjectReferenceApplyConfiguration) *TagReferenceApplyConfiguration {
	b.From = value
	return b
}

// WithReference sets the Reference field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reference field is set to the value of the last call.
func (b *TagReferenceApplyConfiguration) WithReference(value bool) *TagReferenceApplyConfiguration {
	b.Reference = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *TagReferenceApplyConfiguration) WithGeneration(value int64) *TagReferenceApplyConfiguration {
	b.Generation = &value
	return b
}

// WithImportPolicy sets the ImportPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ImportPolicy field is set to the value of the last call.
func (b *TagReferenceApplyConfiguration) WithImportPolicy(value *TagImportPolicyApplyConfiguration) *TagReferenceApplyConfiguration {
	b.ImportPolicy = value
	return b
}

// WithReferencePolicy sets the ReferencePolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ReferencePolicy field is set to the value of the last call.
func (b *TagReferenceApplyConfiguration) WithReferencePolicy(value *TagReferencePolicyApplyConfiguration) *TagReferenceApplyConfiguration {
	b.ReferencePolicy = value
	return b
}
