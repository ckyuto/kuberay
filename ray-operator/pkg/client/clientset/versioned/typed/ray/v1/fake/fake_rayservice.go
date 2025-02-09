// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/ray-project/kuberay/ray-operator/apis/ray/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRayServices implements RayServiceInterface
type FakeRayServices struct {
	Fake *FakeRayV1
	ns   string
}

var rayservicesResource = v1.SchemeGroupVersion.WithResource("rayservices")

var rayservicesKind = v1.SchemeGroupVersion.WithKind("RayService")

// Get takes name of the rayService, and returns the corresponding rayService object, and an error if there is any.
func (c *FakeRayServices) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RayService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rayservicesResource, c.ns, name), &v1.RayService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayService), err
}

// List takes label and field selectors, and returns the list of RayServices that match those selectors.
func (c *FakeRayServices) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RayServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rayservicesResource, rayservicesKind, c.ns, opts), &v1.RayServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RayServiceList{ListMeta: obj.(*v1.RayServiceList).ListMeta}
	for _, item := range obj.(*v1.RayServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rayServices.
func (c *FakeRayServices) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rayservicesResource, c.ns, opts))

}

// Create takes the representation of a rayService and creates it.  Returns the server's representation of the rayService, and an error, if there is any.
func (c *FakeRayServices) Create(ctx context.Context, rayService *v1.RayService, opts metav1.CreateOptions) (result *v1.RayService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rayservicesResource, c.ns, rayService), &v1.RayService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayService), err
}

// Update takes the representation of a rayService and updates it. Returns the server's representation of the rayService, and an error, if there is any.
func (c *FakeRayServices) Update(ctx context.Context, rayService *v1.RayService, opts metav1.UpdateOptions) (result *v1.RayService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rayservicesResource, c.ns, rayService), &v1.RayService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRayServices) UpdateStatus(ctx context.Context, rayService *v1.RayService, opts metav1.UpdateOptions) (*v1.RayService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rayservicesResource, "status", c.ns, rayService), &v1.RayService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayService), err
}

// Delete takes name of the rayService and deletes it. Returns an error if one occurs.
func (c *FakeRayServices) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rayservicesResource, c.ns, name, opts), &v1.RayService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRayServices) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rayservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RayServiceList{})
	return err
}

// Patch applies the patch and returns the patched rayService.
func (c *FakeRayServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RayService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rayservicesResource, c.ns, name, pt, data, subresources...), &v1.RayService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RayService), err
}
