// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"
	json "encoding/json"
	"fmt"

	machineconfigurationv1 "github.com/openshift/api/machineconfiguration/v1"
	applyconfigurationsmachineconfigurationv1 "github.com/openshift/client-go/machineconfiguration/applyconfigurations/machineconfiguration/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeContainerRuntimeConfigs implements ContainerRuntimeConfigInterface
type FakeContainerRuntimeConfigs struct {
	Fake *FakeMachineconfigurationV1
}

var containerruntimeconfigsResource = schema.GroupVersionResource{Group: "machineconfiguration.openshift.io", Version: "v1", Resource: "containerruntimeconfigs"}

var containerruntimeconfigsKind = schema.GroupVersionKind{Group: "machineconfiguration.openshift.io", Version: "v1", Kind: "ContainerRuntimeConfig"}

// Get takes name of the containerRuntimeConfig, and returns the corresponding containerRuntimeConfig object, and an error if there is any.
func (c *FakeContainerRuntimeConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *machineconfigurationv1.ContainerRuntimeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(containerruntimeconfigsResource, name), &machineconfigurationv1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationv1.ContainerRuntimeConfig), err
}

// List takes label and field selectors, and returns the list of ContainerRuntimeConfigs that match those selectors.
func (c *FakeContainerRuntimeConfigs) List(ctx context.Context, opts v1.ListOptions) (result *machineconfigurationv1.ContainerRuntimeConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(containerruntimeconfigsResource, containerruntimeconfigsKind, opts), &machineconfigurationv1.ContainerRuntimeConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &machineconfigurationv1.ContainerRuntimeConfigList{ListMeta: obj.(*machineconfigurationv1.ContainerRuntimeConfigList).ListMeta}
	for _, item := range obj.(*machineconfigurationv1.ContainerRuntimeConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested containerRuntimeConfigs.
func (c *FakeContainerRuntimeConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(containerruntimeconfigsResource, opts))
}

// Create takes the representation of a containerRuntimeConfig and creates it.  Returns the server's representation of the containerRuntimeConfig, and an error, if there is any.
func (c *FakeContainerRuntimeConfigs) Create(ctx context.Context, containerRuntimeConfig *machineconfigurationv1.ContainerRuntimeConfig, opts v1.CreateOptions) (result *machineconfigurationv1.ContainerRuntimeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(containerruntimeconfigsResource, containerRuntimeConfig), &machineconfigurationv1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationv1.ContainerRuntimeConfig), err
}

// Update takes the representation of a containerRuntimeConfig and updates it. Returns the server's representation of the containerRuntimeConfig, and an error, if there is any.
func (c *FakeContainerRuntimeConfigs) Update(ctx context.Context, containerRuntimeConfig *machineconfigurationv1.ContainerRuntimeConfig, opts v1.UpdateOptions) (result *machineconfigurationv1.ContainerRuntimeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(containerruntimeconfigsResource, containerRuntimeConfig), &machineconfigurationv1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationv1.ContainerRuntimeConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeContainerRuntimeConfigs) UpdateStatus(ctx context.Context, containerRuntimeConfig *machineconfigurationv1.ContainerRuntimeConfig, opts v1.UpdateOptions) (*machineconfigurationv1.ContainerRuntimeConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(containerruntimeconfigsResource, "status", containerRuntimeConfig), &machineconfigurationv1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationv1.ContainerRuntimeConfig), err
}

// Delete takes name of the containerRuntimeConfig and deletes it. Returns an error if one occurs.
func (c *FakeContainerRuntimeConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(containerruntimeconfigsResource, name, opts), &machineconfigurationv1.ContainerRuntimeConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeContainerRuntimeConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(containerruntimeconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &machineconfigurationv1.ContainerRuntimeConfigList{})
	return err
}

// Patch applies the patch and returns the patched containerRuntimeConfig.
func (c *FakeContainerRuntimeConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *machineconfigurationv1.ContainerRuntimeConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(containerruntimeconfigsResource, name, pt, data, subresources...), &machineconfigurationv1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationv1.ContainerRuntimeConfig), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied containerRuntimeConfig.
func (c *FakeContainerRuntimeConfigs) Apply(ctx context.Context, containerRuntimeConfig *applyconfigurationsmachineconfigurationv1.ContainerRuntimeConfigApplyConfiguration, opts v1.ApplyOptions) (result *machineconfigurationv1.ContainerRuntimeConfig, err error) {
	if containerRuntimeConfig == nil {
		return nil, fmt.Errorf("containerRuntimeConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(containerRuntimeConfig)
	if err != nil {
		return nil, err
	}
	name := containerRuntimeConfig.Name
	if name == nil {
		return nil, fmt.Errorf("containerRuntimeConfig.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(containerruntimeconfigsResource, *name, types.ApplyPatchType, data), &machineconfigurationv1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationv1.ContainerRuntimeConfig), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeContainerRuntimeConfigs) ApplyStatus(ctx context.Context, containerRuntimeConfig *applyconfigurationsmachineconfigurationv1.ContainerRuntimeConfigApplyConfiguration, opts v1.ApplyOptions) (result *machineconfigurationv1.ContainerRuntimeConfig, err error) {
	if containerRuntimeConfig == nil {
		return nil, fmt.Errorf("containerRuntimeConfig provided to Apply must not be nil")
	}
	data, err := json.Marshal(containerRuntimeConfig)
	if err != nil {
		return nil, err
	}
	name := containerRuntimeConfig.Name
	if name == nil {
		return nil, fmt.Errorf("containerRuntimeConfig.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(containerruntimeconfigsResource, *name, types.ApplyPatchType, data, "status"), &machineconfigurationv1.ContainerRuntimeConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*machineconfigurationv1.ContainerRuntimeConfig), err
}