/*


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
	"context"

	v1beta1 "github.com/VictoriaMetrics/operator/api/victoriametrics/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVMSingles implements VMSingleInterface
type FakeVMSingles struct {
	Fake *FakeVictoriametricsV1beta1
	ns   string
}

var vmsinglesResource = v1beta1.SchemeGroupVersion.WithResource("vmsingles")

var vmsinglesKind = v1beta1.SchemeGroupVersion.WithKind("VMSingle")

// Get takes name of the vMSingle, and returns the corresponding vMSingle object, and an error if there is any.
func (c *FakeVMSingles) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VMSingle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vmsinglesResource, c.ns, name), &v1beta1.VMSingle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMSingle), err
}

// List takes label and field selectors, and returns the list of VMSingles that match those selectors.
func (c *FakeVMSingles) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VMSingleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vmsinglesResource, vmsinglesKind, c.ns, opts), &v1beta1.VMSingleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VMSingleList{ListMeta: obj.(*v1beta1.VMSingleList).ListMeta}
	for _, item := range obj.(*v1beta1.VMSingleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vMSingles.
func (c *FakeVMSingles) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vmsinglesResource, c.ns, opts))

}

// Create takes the representation of a vMSingle and creates it.  Returns the server's representation of the vMSingle, and an error, if there is any.
func (c *FakeVMSingles) Create(ctx context.Context, vMSingle *v1beta1.VMSingle, opts v1.CreateOptions) (result *v1beta1.VMSingle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vmsinglesResource, c.ns, vMSingle), &v1beta1.VMSingle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMSingle), err
}

// Update takes the representation of a vMSingle and updates it. Returns the server's representation of the vMSingle, and an error, if there is any.
func (c *FakeVMSingles) Update(ctx context.Context, vMSingle *v1beta1.VMSingle, opts v1.UpdateOptions) (result *v1beta1.VMSingle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vmsinglesResource, c.ns, vMSingle), &v1beta1.VMSingle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMSingle), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVMSingles) UpdateStatus(ctx context.Context, vMSingle *v1beta1.VMSingle, opts v1.UpdateOptions) (*v1beta1.VMSingle, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vmsinglesResource, "status", c.ns, vMSingle), &v1beta1.VMSingle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMSingle), err
}

// Delete takes name of the vMSingle and deletes it. Returns an error if one occurs.
func (c *FakeVMSingles) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(vmsinglesResource, c.ns, name, opts), &v1beta1.VMSingle{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVMSingles) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vmsinglesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VMSingleList{})
	return err
}

// Patch applies the patch and returns the patched vMSingle.
func (c *FakeVMSingles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMSingle, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vmsinglesResource, c.ns, name, pt, data, subresources...), &v1beta1.VMSingle{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMSingle), err
}
