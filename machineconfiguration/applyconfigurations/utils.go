// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfigurations

import (
	v1 "github.com/openshift/api/machineconfiguration/v1"
	machineconfigurationv1 "github.com/openshift/client-go/machineconfiguration/applyconfigurations/machineconfiguration/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=machineconfiguration.openshift.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("ContainerRuntimeConfig"):
		return &machineconfigurationv1.ContainerRuntimeConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ContainerRuntimeConfigCondition"):
		return &machineconfigurationv1.ContainerRuntimeConfigConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ContainerRuntimeConfigSpec"):
		return &machineconfigurationv1.ContainerRuntimeConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ContainerRuntimeConfigStatus"):
		return &machineconfigurationv1.ContainerRuntimeConfigStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ContainerRuntimeConfiguration"):
		return &machineconfigurationv1.ContainerRuntimeConfigurationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ControllerConfig"):
		return &machineconfigurationv1.ControllerConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ControllerConfigSpec"):
		return &machineconfigurationv1.ControllerConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ControllerConfigStatus"):
		return &machineconfigurationv1.ControllerConfigStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("ControllerConfigStatusCondition"):
		return &machineconfigurationv1.ControllerConfigStatusConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeletConfig"):
		return &machineconfigurationv1.KubeletConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeletConfigCondition"):
		return &machineconfigurationv1.KubeletConfigConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeletConfigSpec"):
		return &machineconfigurationv1.KubeletConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("KubeletConfigStatus"):
		return &machineconfigurationv1.KubeletConfigStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MachineConfig"):
		return &machineconfigurationv1.MachineConfigApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MachineConfigPool"):
		return &machineconfigurationv1.MachineConfigPoolApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MachineConfigPoolCondition"):
		return &machineconfigurationv1.MachineConfigPoolConditionApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MachineConfigPoolSpec"):
		return &machineconfigurationv1.MachineConfigPoolSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MachineConfigPoolStatus"):
		return &machineconfigurationv1.MachineConfigPoolStatusApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MachineConfigPoolStatusConfiguration"):
		return &machineconfigurationv1.MachineConfigPoolStatusConfigurationApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("MachineConfigSpec"):
		return &machineconfigurationv1.MachineConfigSpecApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("NetworkInfo"):
		return &machineconfigurationv1.NetworkInfoApplyConfiguration{}

	}
	return nil
}
