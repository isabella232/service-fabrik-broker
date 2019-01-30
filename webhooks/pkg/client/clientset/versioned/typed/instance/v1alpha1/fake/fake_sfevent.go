/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/cloudfoundry-incubator/service-fabrik-broker/webhooks/pkg/apis/instance/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSfevents implements SfeventInterface
type FakeSfevents struct {
	Fake *FakeSamplecontrollerV1alpha1
	ns   string
}

var sfeventsResource = schema.GroupVersionResource{Group: "samplecontroller.k8s.io", Version: "v1alpha1", Resource: "sfevents"}

var sfeventsKind = schema.GroupVersionKind{Group: "samplecontroller.k8s.io", Version: "v1alpha1", Kind: "Sfevent"}

// Get takes name of the sfevent, and returns the corresponding sfevent object, and an error if there is any.
func (c *FakeSfevents) Get(name string, options v1.GetOptions) (result *v1alpha1.Sfevent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sfeventsResource, c.ns, name), &v1alpha1.Sfevent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sfevent), err
}

// List takes label and field selectors, and returns the list of Sfevents that match those selectors.
func (c *FakeSfevents) List(opts v1.ListOptions) (result *v1alpha1.SfeventList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sfeventsResource, sfeventsKind, c.ns, opts), &v1alpha1.SfeventList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SfeventList{ListMeta: obj.(*v1alpha1.SfeventList).ListMeta}
	for _, item := range obj.(*v1alpha1.SfeventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sfevents.
func (c *FakeSfevents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sfeventsResource, c.ns, opts))

}

// Create takes the representation of a sfevent and creates it.  Returns the server's representation of the sfevent, and an error, if there is any.
func (c *FakeSfevents) Create(sfevent *v1alpha1.Sfevent) (result *v1alpha1.Sfevent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sfeventsResource, c.ns, sfevent), &v1alpha1.Sfevent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sfevent), err
}

// Update takes the representation of a sfevent and updates it. Returns the server's representation of the sfevent, and an error, if there is any.
func (c *FakeSfevents) Update(sfevent *v1alpha1.Sfevent) (result *v1alpha1.Sfevent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sfeventsResource, c.ns, sfevent), &v1alpha1.Sfevent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sfevent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSfevents) UpdateStatus(sfevent *v1alpha1.Sfevent) (*v1alpha1.Sfevent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sfeventsResource, "status", c.ns, sfevent), &v1alpha1.Sfevent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sfevent), err
}

// Delete takes name of the sfevent and deletes it. Returns an error if one occurs.
func (c *FakeSfevents) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sfeventsResource, c.ns, name), &v1alpha1.Sfevent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSfevents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sfeventsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SfeventList{})
	return err
}

// Patch applies the patch and returns the patched sfevent.
func (c *FakeSfevents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Sfevent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sfeventsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Sfevent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Sfevent), err
}
