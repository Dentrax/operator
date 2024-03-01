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

// FakeVMAuths implements VMAuthInterface
type FakeVMAuths struct {
	Fake *FakeVictoriametricsV1beta1
	ns   string
}

var vmauthsResource = v1beta1.SchemeGroupVersion.WithResource("vmauths")

var vmauthsKind = v1beta1.SchemeGroupVersion.WithKind("VMAuth")

// Get takes name of the vMAuth, and returns the corresponding vMAuth object, and an error if there is any.
func (c *FakeVMAuths) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VMAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vmauthsResource, c.ns, name), &v1beta1.VMAuth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAuth), err
}

// List takes label and field selectors, and returns the list of VMAuths that match those selectors.
func (c *FakeVMAuths) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VMAuthList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vmauthsResource, vmauthsKind, c.ns, opts), &v1beta1.VMAuthList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VMAuthList{ListMeta: obj.(*v1beta1.VMAuthList).ListMeta}
	for _, item := range obj.(*v1beta1.VMAuthList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vMAuths.
func (c *FakeVMAuths) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vmauthsResource, c.ns, opts))

}

// Create takes the representation of a vMAuth and creates it.  Returns the server's representation of the vMAuth, and an error, if there is any.
func (c *FakeVMAuths) Create(ctx context.Context, vMAuth *v1beta1.VMAuth, opts v1.CreateOptions) (result *v1beta1.VMAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vmauthsResource, c.ns, vMAuth), &v1beta1.VMAuth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAuth), err
}

// Update takes the representation of a vMAuth and updates it. Returns the server's representation of the vMAuth, and an error, if there is any.
func (c *FakeVMAuths) Update(ctx context.Context, vMAuth *v1beta1.VMAuth, opts v1.UpdateOptions) (result *v1beta1.VMAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vmauthsResource, c.ns, vMAuth), &v1beta1.VMAuth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAuth), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVMAuths) UpdateStatus(ctx context.Context, vMAuth *v1beta1.VMAuth, opts v1.UpdateOptions) (*v1beta1.VMAuth, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vmauthsResource, "status", c.ns, vMAuth), &v1beta1.VMAuth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAuth), err
}

// Delete takes name of the vMAuth and deletes it. Returns an error if one occurs.
func (c *FakeVMAuths) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(vmauthsResource, c.ns, name, opts), &v1beta1.VMAuth{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVMAuths) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vmauthsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VMAuthList{})
	return err
}

// Patch applies the patch and returns the patched vMAuth.
func (c *FakeVMAuths) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vmauthsResource, c.ns, name, pt, data, subresources...), &v1beta1.VMAuth{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAuth), err
}
