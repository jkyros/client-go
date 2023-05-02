// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/openshift/client-go/insights/clientset/versioned/typed/insights/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeInsightsV1alpha1 struct {
	*testing.Fake
}

func (c *FakeInsightsV1alpha1) DataGathers() v1alpha1.DataGatherInterface {
	return &FakeDataGathers{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeInsightsV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}