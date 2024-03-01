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

// FakeVMRules implements VMRuleInterface
type FakeVMRules struct {
	Fake *FakeVictoriametricsV1beta1
	ns   string
}

var vmrulesResource = v1beta1.SchemeGroupVersion.WithResource("vmrules")

var vmrulesKind = v1beta1.SchemeGroupVersion.WithKind("VMRule")

// Get takes name of the vMRule, and returns the corresponding vMRule object, and an error if there is any.
func (c *FakeVMRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VMRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vmrulesResource, c.ns, name), &v1beta1.VMRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMRule), err
}

// List takes label and field selectors, and returns the list of VMRules that match those selectors.
func (c *FakeVMRules) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VMRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vmrulesResource, vmrulesKind, c.ns, opts), &v1beta1.VMRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VMRuleList{ListMeta: obj.(*v1beta1.VMRuleList).ListMeta}
	for _, item := range obj.(*v1beta1.VMRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vMRules.
func (c *FakeVMRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vmrulesResource, c.ns, opts))

}

// Create takes the representation of a vMRule and creates it.  Returns the server's representation of the vMRule, and an error, if there is any.
func (c *FakeVMRules) Create(ctx context.Context, vMRule *v1beta1.VMRule, opts v1.CreateOptions) (result *v1beta1.VMRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vmrulesResource, c.ns, vMRule), &v1beta1.VMRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMRule), err
}

// Update takes the representation of a vMRule and updates it. Returns the server's representation of the vMRule, and an error, if there is any.
func (c *FakeVMRules) Update(ctx context.Context, vMRule *v1beta1.VMRule, opts v1.UpdateOptions) (result *v1beta1.VMRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vmrulesResource, c.ns, vMRule), &v1beta1.VMRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVMRules) UpdateStatus(ctx context.Context, vMRule *v1beta1.VMRule, opts v1.UpdateOptions) (*v1beta1.VMRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vmrulesResource, "status", c.ns, vMRule), &v1beta1.VMRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMRule), err
}

// Delete takes name of the vMRule and deletes it. Returns an error if one occurs.
func (c *FakeVMRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(vmrulesResource, c.ns, name, opts), &v1beta1.VMRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVMRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vmrulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VMRuleList{})
	return err
}

// Patch applies the patch and returns the patched vMRule.
func (c *FakeVMRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vmrulesResource, c.ns, name, pt, data, subresources...), &v1beta1.VMRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMRule), err
}
