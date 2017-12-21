/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubernetes-practice/k8s-custom-resource/pkg/apis/podwatchers.nahid.try.com/v1alpha1"
	scheme "kubernetes-practice/k8s-custom-resource/pkg/client/clientset/versioned/scheme"
)

// PodWatchersGetter has a method to return a PodWatcherInterface.
// A group's client should implement this interface.
type PodWatchersGetter interface {
	PodWatchers(namespace string) PodWatcherInterface
}

// PodWatcherInterface has methods to work with PodWatcher resources.
type PodWatcherInterface interface {
	Create(*v1alpha1.PodWatcher) (*v1alpha1.PodWatcher, error)
	Update(*v1alpha1.PodWatcher) (*v1alpha1.PodWatcher, error)
	UpdateStatus(*v1alpha1.PodWatcher) (*v1alpha1.PodWatcher, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.PodWatcher, error)
	List(opts v1.ListOptions) (*v1alpha1.PodWatcherList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodWatcher, err error)
	PodWatcherExpansion
}

// podWatchers implements PodWatcherInterface
type podWatchers struct {
	client rest.Interface
	ns     string
}

// newPodWatchers returns a PodWatchers
func newPodWatchers(c *PodwatchersV1alpha1Client, namespace string) *podWatchers {
	return &podWatchers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the podWatcher, and returns the corresponding podWatcher object, and an error if there is any.
func (c *podWatchers) Get(name string, options v1.GetOptions) (result *v1alpha1.PodWatcher, err error) {
	result = &v1alpha1.PodWatcher{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podwatchers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PodWatchers that match those selectors.
func (c *podWatchers) List(opts v1.ListOptions) (result *v1alpha1.PodWatcherList, err error) {
	result = &v1alpha1.PodWatcherList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("podwatchers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested podWatchers.
func (c *podWatchers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("podwatchers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a podWatcher and creates it.  Returns the server's representation of the podWatcher, and an error, if there is any.
func (c *podWatchers) Create(podWatcher *v1alpha1.PodWatcher) (result *v1alpha1.PodWatcher, err error) {
	result = &v1alpha1.PodWatcher{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("podwatchers").
		Body(podWatcher).
		Do().
		Into(result)
	return
}

// Update takes the representation of a podWatcher and updates it. Returns the server's representation of the podWatcher, and an error, if there is any.
func (c *podWatchers) Update(podWatcher *v1alpha1.PodWatcher) (result *v1alpha1.PodWatcher, err error) {
	result = &v1alpha1.PodWatcher{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("podwatchers").
		Name(podWatcher.Name).
		Body(podWatcher).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *podWatchers) UpdateStatus(podWatcher *v1alpha1.PodWatcher) (result *v1alpha1.PodWatcher, err error) {
	result = &v1alpha1.PodWatcher{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("podwatchers").
		Name(podWatcher.Name).
		SubResource("status").
		Body(podWatcher).
		Do().
		Into(result)
	return
}

// Delete takes name of the podWatcher and deletes it. Returns an error if one occurs.
func (c *podWatchers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podwatchers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *podWatchers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("podwatchers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched podWatcher.
func (c *podWatchers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodWatcher, err error) {
	result = &v1alpha1.PodWatcher{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("podwatchers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
