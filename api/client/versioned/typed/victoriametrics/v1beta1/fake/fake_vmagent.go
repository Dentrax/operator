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

// FakeVMAgents implements VMAgentInterface
type FakeVMAgents struct {
	Fake *FakeVictoriametricsV1beta1
	ns   string
}

var vmagentsResource = v1beta1.SchemeGroupVersion.WithResource("vmagents")

var vmagentsKind = v1beta1.SchemeGroupVersion.WithKind("VMAgent")

// Get takes name of the vMAgent, and returns the corresponding vMAgent object, and an error if there is any.
func (c *FakeVMAgents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.VMAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vmagentsResource, c.ns, name), &v1beta1.VMAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAgent), err
}

// List takes label and field selectors, and returns the list of VMAgents that match those selectors.
func (c *FakeVMAgents) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.VMAgentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vmagentsResource, vmagentsKind, c.ns, opts), &v1beta1.VMAgentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.VMAgentList{ListMeta: obj.(*v1beta1.VMAgentList).ListMeta}
	for _, item := range obj.(*v1beta1.VMAgentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vMAgents.
func (c *FakeVMAgents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vmagentsResource, c.ns, opts))

}

// Create takes the representation of a vMAgent and creates it.  Returns the server's representation of the vMAgent, and an error, if there is any.
func (c *FakeVMAgents) Create(ctx context.Context, vMAgent *v1beta1.VMAgent, opts v1.CreateOptions) (result *v1beta1.VMAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vmagentsResource, c.ns, vMAgent), &v1beta1.VMAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAgent), err
}

// Update takes the representation of a vMAgent and updates it. Returns the server's representation of the vMAgent, and an error, if there is any.
func (c *FakeVMAgents) Update(ctx context.Context, vMAgent *v1beta1.VMAgent, opts v1.UpdateOptions) (result *v1beta1.VMAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vmagentsResource, c.ns, vMAgent), &v1beta1.VMAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAgent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVMAgents) UpdateStatus(ctx context.Context, vMAgent *v1beta1.VMAgent, opts v1.UpdateOptions) (*v1beta1.VMAgent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vmagentsResource, "status", c.ns, vMAgent), &v1beta1.VMAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAgent), err
}

// Delete takes name of the vMAgent and deletes it. Returns an error if one occurs.
func (c *FakeVMAgents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(vmagentsResource, c.ns, name, opts), &v1beta1.VMAgent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVMAgents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vmagentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.VMAgentList{})
	return err
}

// Patch applies the patch and returns the patched vMAgent.
func (c *FakeVMAgents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VMAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vmagentsResource, c.ns, name, pt, data, subresources...), &v1beta1.VMAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.VMAgent), err
}
