// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/tiagoangelozup/horusec-admin/pkg/api/install/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHorusecManagerSpecs implements HorusecManagerSpecInterface
type FakeHorusecManagerSpecs struct {
	Fake *FakeInstallV1alpha1
	ns   string
}

var horusecmanagerspecsResource = schema.GroupVersionResource{Group: "install.horusec.io", Version: "v1alpha1", Resource: "horusecmanagerspecs"}

var horusecmanagerspecsKind = schema.GroupVersionKind{Group: "install.horusec.io", Version: "v1alpha1", Kind: "HorusecManagerSpec"}

// Get takes name of the horusecManagerSpec, and returns the corresponding horusecManagerSpec object, and an error if there is any.
func (c *FakeHorusecManagerSpecs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HorusecManagerSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(horusecmanagerspecsResource, c.ns, name), &v1alpha1.HorusecManagerSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HorusecManagerSpec), err
}

// List takes label and field selectors, and returns the list of HorusecManagerSpecs that match those selectors.
func (c *FakeHorusecManagerSpecs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HorusecManagerSpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(horusecmanagerspecsResource, horusecmanagerspecsKind, c.ns, opts), &v1alpha1.HorusecManagerSpecList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HorusecManagerSpecList), err
}

// Watch returns a watch.Interface that watches the requested horusecManagerSpecs.
func (c *FakeHorusecManagerSpecs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(horusecmanagerspecsResource, c.ns, opts))

}

// Create takes the representation of a horusecManagerSpec and creates it.  Returns the server's representation of the horusecManagerSpec, and an error, if there is any.
func (c *FakeHorusecManagerSpecs) Create(ctx context.Context, horusecManagerSpec *v1alpha1.HorusecManagerSpec, opts v1.CreateOptions) (result *v1alpha1.HorusecManagerSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(horusecmanagerspecsResource, c.ns, horusecManagerSpec), &v1alpha1.HorusecManagerSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HorusecManagerSpec), err
}

// Update takes the representation of a horusecManagerSpec and updates it. Returns the server's representation of the horusecManagerSpec, and an error, if there is any.
func (c *FakeHorusecManagerSpecs) Update(ctx context.Context, horusecManagerSpec *v1alpha1.HorusecManagerSpec, opts v1.UpdateOptions) (result *v1alpha1.HorusecManagerSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(horusecmanagerspecsResource, c.ns, horusecManagerSpec), &v1alpha1.HorusecManagerSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HorusecManagerSpec), err
}

// Delete takes name of the horusecManagerSpec and deletes it. Returns an error if one occurs.
func (c *FakeHorusecManagerSpecs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(horusecmanagerspecsResource, c.ns, name), &v1alpha1.HorusecManagerSpec{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHorusecManagerSpecs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(horusecmanagerspecsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HorusecManagerSpecList{})
	return err
}

// Patch applies the patch and returns the patched horusecManagerSpec.
func (c *FakeHorusecManagerSpecs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HorusecManagerSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(horusecmanagerspecsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HorusecManagerSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HorusecManagerSpec), err
}