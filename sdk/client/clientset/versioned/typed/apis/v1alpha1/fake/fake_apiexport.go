/*
Copyright The KCP Authors.

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
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apis/v1alpha1"
	apisv1alpha1 "github.com/kcp-dev/kcp/sdk/client/applyconfiguration/apis/v1alpha1"
)

// FakeAPIExports implements APIExportInterface
type FakeAPIExports struct {
	Fake *FakeApisV1alpha1
}

var apiexportsResource = schema.GroupVersionResource{Group: "apis.kcp.io", Version: "v1alpha1", Resource: "apiexports"}

var apiexportsKind = schema.GroupVersionKind{Group: "apis.kcp.io", Version: "v1alpha1", Kind: "APIExport"}

// Get takes name of the aPIExport, and returns the corresponding aPIExport object, and an error if there is any.
func (c *FakeAPIExports) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.APIExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(apiexportsResource, name), &v1alpha1.APIExport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExport), err
}

// List takes label and field selectors, and returns the list of APIExports that match those selectors.
func (c *FakeAPIExports) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.APIExportList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(apiexportsResource, apiexportsKind, opts), &v1alpha1.APIExportList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.APIExportList{ListMeta: obj.(*v1alpha1.APIExportList).ListMeta}
	for _, item := range obj.(*v1alpha1.APIExportList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aPIExports.
func (c *FakeAPIExports) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(apiexportsResource, opts))
}

// Create takes the representation of a aPIExport and creates it.  Returns the server's representation of the aPIExport, and an error, if there is any.
func (c *FakeAPIExports) Create(ctx context.Context, aPIExport *v1alpha1.APIExport, opts v1.CreateOptions) (result *v1alpha1.APIExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(apiexportsResource, aPIExport), &v1alpha1.APIExport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExport), err
}

// Update takes the representation of a aPIExport and updates it. Returns the server's representation of the aPIExport, and an error, if there is any.
func (c *FakeAPIExports) Update(ctx context.Context, aPIExport *v1alpha1.APIExport, opts v1.UpdateOptions) (result *v1alpha1.APIExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(apiexportsResource, aPIExport), &v1alpha1.APIExport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExport), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAPIExports) UpdateStatus(ctx context.Context, aPIExport *v1alpha1.APIExport, opts v1.UpdateOptions) (*v1alpha1.APIExport, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(apiexportsResource, "status", aPIExport), &v1alpha1.APIExport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExport), err
}

// Delete takes name of the aPIExport and deletes it. Returns an error if one occurs.
func (c *FakeAPIExports) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(apiexportsResource, name, opts), &v1alpha1.APIExport{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAPIExports) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(apiexportsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.APIExportList{})
	return err
}

// Patch applies the patch and returns the patched aPIExport.
func (c *FakeAPIExports) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.APIExport, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apiexportsResource, name, pt, data, subresources...), &v1alpha1.APIExport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExport), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied aPIExport.
func (c *FakeAPIExports) Apply(ctx context.Context, aPIExport *apisv1alpha1.APIExportApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.APIExport, err error) {
	if aPIExport == nil {
		return nil, fmt.Errorf("aPIExport provided to Apply must not be nil")
	}
	data, err := json.Marshal(aPIExport)
	if err != nil {
		return nil, err
	}
	name := aPIExport.Name
	if name == nil {
		return nil, fmt.Errorf("aPIExport.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apiexportsResource, *name, types.ApplyPatchType, data), &v1alpha1.APIExport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExport), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeAPIExports) ApplyStatus(ctx context.Context, aPIExport *apisv1alpha1.APIExportApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.APIExport, err error) {
	if aPIExport == nil {
		return nil, fmt.Errorf("aPIExport provided to Apply must not be nil")
	}
	data, err := json.Marshal(aPIExport)
	if err != nil {
		return nil, err
	}
	name := aPIExport.Name
	if name == nil {
		return nil, fmt.Errorf("aPIExport.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(apiexportsResource, *name, types.ApplyPatchType, data, "status"), &v1alpha1.APIExport{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.APIExport), err
}
