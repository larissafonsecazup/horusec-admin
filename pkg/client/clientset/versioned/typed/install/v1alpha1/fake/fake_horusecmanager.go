// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/tiagoangelozup/horusec-admin/pkg/api/install/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHorusecManagers implements HorusecManagerInterface
type FakeHorusecManagers struct {
	Fake *FakeInstallV1alpha1
	ns   string
}

var horusecmanagersResource = schema.GroupVersionResource{Group: "install.horusec.io", Version: "v1alpha1", Resource: "horusecmanagers"}

var horusecmanagersKind = schema.GroupVersionKind{Group: "install.horusec.io", Version: "v1alpha1", Kind: "HorusecManager"}

// Get takes name of the horusecManager, and returns the corresponding horusecManager object, and an error if there is any.
func (c *FakeHorusecManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HorusecManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(horusecmanagersResource, c.ns, name), &v1alpha1.HorusecManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HorusecManager), err
}

// List takes label and field selectors, and returns the list of HorusecManagers that match those selectors.
func (c *FakeHorusecManagers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HorusecManagerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(horusecmanagersResource, horusecmanagersKind, c.ns, opts), &v1alpha1.HorusecManagerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HorusecManagerList{ListMeta: obj.(*v1alpha1.HorusecManagerList).ListMeta}
	for _, item := range obj.(*v1alpha1.HorusecManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested horusecManagers.
func (c *FakeHorusecManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(horusecmanagersResource, c.ns, opts))

}

// Create takes the representation of a horusecManager and creates it.  Returns the server's representation of the horusecManager, and an error, if there is any.
func (c *FakeHorusecManagers) Create(ctx context.Context, horusecManager *v1alpha1.HorusecManager, opts v1.CreateOptions) (result *v1alpha1.HorusecManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(horusecmanagersResource, c.ns, horusecManager), &v1alpha1.HorusecManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HorusecManager), err
}

// Update takes the representation of a horusecManager and updates it. Returns the server's representation of the horusecManager, and an error, if there is any.
func (c *FakeHorusecManagers) Update(ctx context.Context, horusecManager *v1alpha1.HorusecManager, opts v1.UpdateOptions) (result *v1alpha1.HorusecManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(horusecmanagersResource, c.ns, horusecManager), &v1alpha1.HorusecManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HorusecManager), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHorusecManagers) UpdateStatus(ctx context.Context, horusecManager *v1alpha1.HorusecManager, opts v1.UpdateOptions) (*v1alpha1.HorusecManager, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(horusecmanagersResource, "status", c.ns, horusecManager), &v1alpha1.HorusecManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HorusecManager), err
}

// Delete takes name of the horusecManager and deletes it. Returns an error if one occurs.
func (c *FakeHorusecManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(horusecmanagersResource, c.ns, name), &v1alpha1.HorusecManager{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHorusecManagers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(horusecmanagersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HorusecManagerList{})
	return err
}

// Patch applies the patch and returns the patched horusecManager.
func (c *FakeHorusecManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HorusecManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(horusecmanagersResource, c.ns, name, pt, data, subresources...), &v1alpha1.HorusecManager{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HorusecManager), err
}