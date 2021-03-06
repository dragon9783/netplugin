/*
Copyright 2016 The Kubernetes Authors.

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

package v1beta1

import (
	api "github.com/contiv/client-go/pkg/api"
	v1 "github.com/contiv/client-go/pkg/api/v1"
	v1beta1 "github.com/contiv/client-go/pkg/apis/extensions/v1beta1"
	meta_v1 "github.com/contiv/client-go/pkg/apis/meta/v1"
	watch "github.com/contiv/client-go/pkg/watch"
	rest "github.com/contiv/client-go/rest"
)

// NetworkPoliciesGetter has a method to return a NetworkPolicyInterface.
// A group's client should implement this interface.
type NetworkPoliciesGetter interface {
	NetworkPolicies(namespace string) NetworkPolicyInterface
}

// NetworkPolicyInterface has methods to work with NetworkPolicy resources.
type NetworkPolicyInterface interface {
	Create(*v1beta1.NetworkPolicy) (*v1beta1.NetworkPolicy, error)
	Update(*v1beta1.NetworkPolicy) (*v1beta1.NetworkPolicy, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1beta1.NetworkPolicy, error)
	List(opts v1.ListOptions) (*v1beta1.NetworkPolicyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *v1beta1.NetworkPolicy, err error)
	NetworkPolicyExpansion
}

// networkPolicies implements NetworkPolicyInterface
type networkPolicies struct {
	client rest.Interface
	ns     string
}

// newNetworkPolicies returns a NetworkPolicies
func newNetworkPolicies(c *ExtensionsV1beta1Client, namespace string) *networkPolicies {
	return &networkPolicies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a networkPolicy and creates it.  Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *networkPolicies) Create(networkPolicy *v1beta1.NetworkPolicy) (result *v1beta1.NetworkPolicy, err error) {
	result = &v1beta1.NetworkPolicy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkpolicies").
		Body(networkPolicy).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkPolicy and updates it. Returns the server's representation of the networkPolicy, and an error, if there is any.
func (c *networkPolicies) Update(networkPolicy *v1beta1.NetworkPolicy) (result *v1beta1.NetworkPolicy, err error) {
	result = &v1beta1.NetworkPolicy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(networkPolicy.Name).
		Body(networkPolicy).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkPolicy and deletes it. Returns an error if one occurs.
func (c *networkPolicies) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkPolicies) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&listOptions, api.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the networkPolicy, and returns the corresponding networkPolicy object, and an error if there is any.
func (c *networkPolicies) Get(name string, options meta_v1.GetOptions) (result *v1beta1.NetworkPolicy, err error) {
	result = &v1beta1.NetworkPolicy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkpolicies").
		Name(name).
		VersionedParams(&options, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkPolicies that match those selectors.
func (c *networkPolicies) List(opts v1.ListOptions) (result *v1beta1.NetworkPolicyList, err error) {
	result = &v1beta1.NetworkPolicyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkPolicies.
func (c *networkPolicies) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.client.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("networkpolicies").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched networkPolicy.
func (c *networkPolicies) Patch(name string, pt api.PatchType, data []byte, subresources ...string) (result *v1beta1.NetworkPolicy, err error) {
	result = &v1beta1.NetworkPolicy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkpolicies").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
