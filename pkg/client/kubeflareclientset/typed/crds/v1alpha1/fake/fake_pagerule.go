/*
Copyright 2020 Replicated, Inc.

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

	v1alpha1 "github.com/replicatedhq/kubeflare/pkg/apis/crds/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePageRules implements PageRuleInterface
type FakePageRules struct {
	Fake *FakeCrdsV1alpha1
	ns   string
}

var pagerulesResource = schema.GroupVersionResource{Group: "crds.kubeflare.io", Version: "v1alpha1", Resource: "pagerules"}

var pagerulesKind = schema.GroupVersionKind{Group: "crds.kubeflare.io", Version: "v1alpha1", Kind: "PageRule"}

// Get takes name of the pageRule, and returns the corresponding pageRule object, and an error if there is any.
func (c *FakePageRules) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PageRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(pagerulesResource, c.ns, name), &v1alpha1.PageRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PageRule), err
}

// List takes label and field selectors, and returns the list of PageRules that match those selectors.
func (c *FakePageRules) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PageRuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(pagerulesResource, pagerulesKind, c.ns, opts), &v1alpha1.PageRuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PageRuleList{ListMeta: obj.(*v1alpha1.PageRuleList).ListMeta}
	for _, item := range obj.(*v1alpha1.PageRuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested pageRules.
func (c *FakePageRules) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(pagerulesResource, c.ns, opts))

}

// Create takes the representation of a pageRule and creates it.  Returns the server's representation of the pageRule, and an error, if there is any.
func (c *FakePageRules) Create(ctx context.Context, pageRule *v1alpha1.PageRule, opts v1.CreateOptions) (result *v1alpha1.PageRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(pagerulesResource, c.ns, pageRule), &v1alpha1.PageRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PageRule), err
}

// Update takes the representation of a pageRule and updates it. Returns the server's representation of the pageRule, and an error, if there is any.
func (c *FakePageRules) Update(ctx context.Context, pageRule *v1alpha1.PageRule, opts v1.UpdateOptions) (result *v1alpha1.PageRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(pagerulesResource, c.ns, pageRule), &v1alpha1.PageRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PageRule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePageRules) UpdateStatus(ctx context.Context, pageRule *v1alpha1.PageRule, opts v1.UpdateOptions) (*v1alpha1.PageRule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(pagerulesResource, "status", c.ns, pageRule), &v1alpha1.PageRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PageRule), err
}

// Delete takes name of the pageRule and deletes it. Returns an error if one occurs.
func (c *FakePageRules) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(pagerulesResource, c.ns, name), &v1alpha1.PageRule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePageRules) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(pagerulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.PageRuleList{})
	return err
}

// Patch applies the patch and returns the patched pageRule.
func (c *FakePageRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PageRule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(pagerulesResource, c.ns, name, pt, data, subresources...), &v1alpha1.PageRule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PageRule), err
}
