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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "kubernetes-practice/k8s-custom-resource/pkg/apis/podwatchers.nahid.try.com/v1alpha1"
)

// FakePodWatchers implements PodWatcherInterface
type FakePodWatchers struct {
	Fake *FakePodwatchersV1alpha1
	ns   string
}

var podwatchersResource = schema.GroupVersionResource{Group: "podwatchers.nahid.try.com", Version: "v1alpha1", Resource: "podwatchers"}

var podwatchersKind = schema.GroupVersionKind{Group: "podwatchers.nahid.try.com", Version: "v1alpha1", Kind: "PodWatcher"}

// Get takes name of the podWatcher, and returns the corresponding podWatcher object, and an error if there is any.
func (c *FakePodWatchers) Get(name string, options v1.GetOptions) (result *v1alpha1.PodWatcher, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podwatchersResource, c.ns, name), &v1alpha1.PodWatcher{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodWatcher), err
}

// List takes label and field selectors, and returns the list of PodWatchers that match those selectors.
func (c *FakePodWatchers) List(opts v1.ListOptions) (result *v1alpha1.PodWatcherList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podwatchersResource, podwatchersKind, c.ns, opts), &v1alpha1.PodWatcherList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodWatcherList{}
	for _, item := range obj.(*v1alpha1.PodWatcherList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podWatchers.
func (c *FakePodWatchers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podwatchersResource, c.ns, opts))

}

// Create takes the representation of a podWatcher and creates it.  Returns the server's representation of the podWatcher, and an error, if there is any.
func (c *FakePodWatchers) Create(podWatcher *v1alpha1.PodWatcher) (result *v1alpha1.PodWatcher, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podwatchersResource, c.ns, podWatcher), &v1alpha1.PodWatcher{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodWatcher), err
}

// Update takes the representation of a podWatcher and updates it. Returns the server's representation of the podWatcher, and an error, if there is any.
func (c *FakePodWatchers) Update(podWatcher *v1alpha1.PodWatcher) (result *v1alpha1.PodWatcher, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podwatchersResource, c.ns, podWatcher), &v1alpha1.PodWatcher{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodWatcher), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodWatchers) UpdateStatus(podWatcher *v1alpha1.PodWatcher) (*v1alpha1.PodWatcher, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podwatchersResource, "status", c.ns, podWatcher), &v1alpha1.PodWatcher{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodWatcher), err
}

// Delete takes name of the podWatcher and deletes it. Returns an error if one occurs.
func (c *FakePodWatchers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podwatchersResource, c.ns, name), &v1alpha1.PodWatcher{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodWatchers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podwatchersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PodWatcherList{})
	return err
}

// Patch applies the patch and returns the patched podWatcher.
func (c *FakePodWatchers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodWatcher, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podwatchersResource, c.ns, name, data, subresources...), &v1alpha1.PodWatcher{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodWatcher), err
}
