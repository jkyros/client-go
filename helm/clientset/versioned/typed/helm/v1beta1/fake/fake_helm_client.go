// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.com/openshift/client-go/helm/clientset/versioned/typed/helm/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeHelmV1beta1 struct {
	*testing.Fake
}

func (c *FakeHelmV1beta1) HelmChartRepositories() v1beta1.HelmChartRepositoryInterface {
	return &FakeHelmChartRepositories{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeHelmV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}