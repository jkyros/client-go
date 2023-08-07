// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	authorizationv1 "github.com/openshift/api/authorization/v1"
	internal "github.com/openshift/client-go/authorization/applyconfigurations/internal"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	managedfields "k8s.io/apimachinery/pkg/util/managedfields"
	corev1 "k8s.io/client-go/applyconfigurations/core/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// RoleBindingApplyConfiguration represents an declarative configuration of the RoleBinding type for use
// with apply.
type RoleBindingApplyConfiguration struct {
	v1.TypeMetaApplyConfiguration    `json:",inline"`
	*v1.ObjectMetaApplyConfiguration `json:"metadata,omitempty"`
	UserNames                        *authorizationv1.OptionalNames             `json:"userNames,omitempty"`
	GroupNames                       *authorizationv1.OptionalNames             `json:"groupNames,omitempty"`
	Subjects                         []corev1.ObjectReferenceApplyConfiguration `json:"subjects,omitempty"`
	RoleRef                          *corev1.ObjectReferenceApplyConfiguration  `json:"roleRef,omitempty"`
}

// RoleBinding constructs an declarative configuration of the RoleBinding type for use with
// apply.
func RoleBinding(name, namespace string) *RoleBindingApplyConfiguration {
	b := &RoleBindingApplyConfiguration{}
	b.WithName(name)
	b.WithNamespace(namespace)
	b.WithKind("RoleBinding")
	b.WithAPIVersion("authorization.openshift.io/v1")
	return b
}

// ExtractRoleBinding extracts the applied configuration owned by fieldManager from
// roleBinding. If no managedFields are found in roleBinding for fieldManager, a
// RoleBindingApplyConfiguration is returned with only the Name, Namespace (if applicable),
// APIVersion and Kind populated. It is possible that no managed fields were found for because other
// field managers have taken ownership of all the fields previously owned by fieldManager, or because
// the fieldManager never owned fields any fields.
// roleBinding must be a unmodified RoleBinding API object that was retrieved from the Kubernetes API.
// ExtractRoleBinding provides a way to perform a extract/modify-in-place/apply workflow.
// Note that an extracted apply configuration will contain fewer fields than what the fieldManager previously
// applied if another fieldManager has updated or force applied any of the previously applied fields.
// Experimental!
func ExtractRoleBinding(roleBinding *authorizationv1.RoleBinding, fieldManager string) (*RoleBindingApplyConfiguration, error) {
	return extractRoleBinding(roleBinding, fieldManager, "")
}

// ExtractRoleBindingStatus is the same as ExtractRoleBinding except
// that it extracts the status subresource applied configuration.
// Experimental!
func ExtractRoleBindingStatus(roleBinding *authorizationv1.RoleBinding, fieldManager string) (*RoleBindingApplyConfiguration, error) {
	return extractRoleBinding(roleBinding, fieldManager, "status")
}

func extractRoleBinding(roleBinding *authorizationv1.RoleBinding, fieldManager string, subresource string) (*RoleBindingApplyConfiguration, error) {
	b := &RoleBindingApplyConfiguration{}
	err := managedfields.ExtractInto(roleBinding, internal.Parser().Type("com.github.openshift.api.authorization.v1.RoleBinding"), fieldManager, b, subresource)
	if err != nil {
		return nil, err
	}
	b.WithName(roleBinding.Name)
	b.WithNamespace(roleBinding.Namespace)

	b.WithKind("RoleBinding")
	b.WithAPIVersion("authorization.openshift.io/v1")
	return b, nil
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithKind(value string) *RoleBindingApplyConfiguration {
	b.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithAPIVersion(value string) *RoleBindingApplyConfiguration {
	b.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithName(value string) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Name = &value
	return b
}

// WithGenerateName sets the GenerateName field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GenerateName field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithGenerateName(value string) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.GenerateName = &value
	return b
}

// WithNamespace sets the Namespace field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Namespace field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithNamespace(value string) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Namespace = &value
	return b
}

// WithUID sets the UID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UID field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithUID(value types.UID) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.UID = &value
	return b
}

// WithResourceVersion sets the ResourceVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ResourceVersion field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithResourceVersion(value string) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.ResourceVersion = &value
	return b
}

// WithGeneration sets the Generation field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Generation field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithGeneration(value int64) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.Generation = &value
	return b
}

// WithCreationTimestamp sets the CreationTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CreationTimestamp field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithCreationTimestamp(value metav1.Time) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.CreationTimestamp = &value
	return b
}

// WithDeletionTimestamp sets the DeletionTimestamp field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionTimestamp field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithDeletionTimestamp(value metav1.Time) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionTimestamp = &value
	return b
}

// WithDeletionGracePeriodSeconds sets the DeletionGracePeriodSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the DeletionGracePeriodSeconds field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithDeletionGracePeriodSeconds(value int64) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	b.DeletionGracePeriodSeconds = &value
	return b
}

// WithLabels puts the entries into the Labels field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Labels field,
// overwriting an existing map entries in Labels field with the same key.
func (b *RoleBindingApplyConfiguration) WithLabels(entries map[string]string) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Labels == nil && len(entries) > 0 {
		b.Labels = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Labels[k] = v
	}
	return b
}

// WithAnnotations puts the entries into the Annotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Annotations field,
// overwriting an existing map entries in Annotations field with the same key.
func (b *RoleBindingApplyConfiguration) WithAnnotations(entries map[string]string) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	if b.Annotations == nil && len(entries) > 0 {
		b.Annotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.Annotations[k] = v
	}
	return b
}

// WithOwnerReferences adds the given value to the OwnerReferences field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the OwnerReferences field.
func (b *RoleBindingApplyConfiguration) WithOwnerReferences(values ...*v1.OwnerReferenceApplyConfiguration) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithOwnerReferences")
		}
		b.OwnerReferences = append(b.OwnerReferences, *values[i])
	}
	return b
}

// WithFinalizers adds the given value to the Finalizers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Finalizers field.
func (b *RoleBindingApplyConfiguration) WithFinalizers(values ...string) *RoleBindingApplyConfiguration {
	b.ensureObjectMetaApplyConfigurationExists()
	for i := range values {
		b.Finalizers = append(b.Finalizers, values[i])
	}
	return b
}

func (b *RoleBindingApplyConfiguration) ensureObjectMetaApplyConfigurationExists() {
	if b.ObjectMetaApplyConfiguration == nil {
		b.ObjectMetaApplyConfiguration = &v1.ObjectMetaApplyConfiguration{}
	}
}

// WithUserNames sets the UserNames field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the UserNames field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithUserNames(value authorizationv1.OptionalNames) *RoleBindingApplyConfiguration {
	b.UserNames = &value
	return b
}

// WithGroupNames sets the GroupNames field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GroupNames field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithGroupNames(value authorizationv1.OptionalNames) *RoleBindingApplyConfiguration {
	b.GroupNames = &value
	return b
}

// WithSubjects adds the given value to the Subjects field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Subjects field.
func (b *RoleBindingApplyConfiguration) WithSubjects(values ...*corev1.ObjectReferenceApplyConfiguration) *RoleBindingApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithSubjects")
		}
		b.Subjects = append(b.Subjects, *values[i])
	}
	return b
}

// WithRoleRef sets the RoleRef field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RoleRef field is set to the value of the last call.
func (b *RoleBindingApplyConfiguration) WithRoleRef(value *corev1.ObjectReferenceApplyConfiguration) *RoleBindingApplyConfiguration {
	b.RoleRef = value
	return b
}
