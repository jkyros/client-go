// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/core/v1"
)

// BuildOutputApplyConfiguration represents an declarative configuration of the BuildOutput type for use
// with apply.
type BuildOutputApplyConfiguration struct {
	To          *v1.ObjectReferenceApplyConfiguration `json:"to,omitempty"`
	PushSecret  *corev1.LocalObjectReference          `json:"pushSecret,omitempty"`
	ImageLabels []ImageLabelApplyConfiguration        `json:"imageLabels,omitempty"`
}

// BuildOutputApplyConfiguration constructs an declarative configuration of the BuildOutput type for use with
// apply.
func BuildOutput() *BuildOutputApplyConfiguration {
	return &BuildOutputApplyConfiguration{}
}

// WithTo sets the To field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the To field is set to the value of the last call.
func (b *BuildOutputApplyConfiguration) WithTo(value *v1.ObjectReferenceApplyConfiguration) *BuildOutputApplyConfiguration {
	b.To = value
	return b
}

// WithPushSecret sets the PushSecret field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PushSecret field is set to the value of the last call.
func (b *BuildOutputApplyConfiguration) WithPushSecret(value corev1.LocalObjectReference) *BuildOutputApplyConfiguration {
	b.PushSecret = &value
	return b
}

// WithImageLabels adds the given value to the ImageLabels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the ImageLabels field.
func (b *BuildOutputApplyConfiguration) WithImageLabels(values ...*ImageLabelApplyConfiguration) *BuildOutputApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithImageLabels")
		}
		b.ImageLabels = append(b.ImageLabels, *values[i])
	}
	return b
}
