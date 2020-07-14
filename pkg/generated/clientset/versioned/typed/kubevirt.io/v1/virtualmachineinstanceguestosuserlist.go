/*
Copyright 2020 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	scheme "github.com/cnrancher/rancher-vm/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "kubevirt.io/client-go/api/v1"
)

// VirtualMachineInstanceGuestOSUserListsGetter has a method to return a VirtualMachineInstanceGuestOSUserListInterface.
// A group's client should implement this interface.
type VirtualMachineInstanceGuestOSUserListsGetter interface {
	VirtualMachineInstanceGuestOSUserLists(namespace string) VirtualMachineInstanceGuestOSUserListInterface
}

// VirtualMachineInstanceGuestOSUserListInterface has methods to work with VirtualMachineInstanceGuestOSUserList resources.
type VirtualMachineInstanceGuestOSUserListInterface interface {
	Create(ctx context.Context, virtualMachineInstanceGuestOSUserList *v1.VirtualMachineInstanceGuestOSUserList, opts metav1.CreateOptions) (*v1.VirtualMachineInstanceGuestOSUserList, error)
	Update(ctx context.Context, virtualMachineInstanceGuestOSUserList *v1.VirtualMachineInstanceGuestOSUserList, opts metav1.UpdateOptions) (*v1.VirtualMachineInstanceGuestOSUserList, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.VirtualMachineInstanceGuestOSUserList, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.VirtualMachineInstanceGuestOSUserListList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstanceGuestOSUserList, err error)
	VirtualMachineInstanceGuestOSUserListExpansion
}

// virtualMachineInstanceGuestOSUserLists implements VirtualMachineInstanceGuestOSUserListInterface
type virtualMachineInstanceGuestOSUserLists struct {
	client rest.Interface
	ns     string
}

// newVirtualMachineInstanceGuestOSUserLists returns a VirtualMachineInstanceGuestOSUserLists
func newVirtualMachineInstanceGuestOSUserLists(c *KubevirtV1Client, namespace string) *virtualMachineInstanceGuestOSUserLists {
	return &virtualMachineInstanceGuestOSUserLists{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualMachineInstanceGuestOSUserList, and returns the corresponding virtualMachineInstanceGuestOSUserList object, and an error if there is any.
func (c *virtualMachineInstanceGuestOSUserLists) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.VirtualMachineInstanceGuestOSUserList, err error) {
	result = &v1.VirtualMachineInstanceGuestOSUserList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstanceguestosuserlists").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualMachineInstanceGuestOSUserLists that match those selectors.
func (c *virtualMachineInstanceGuestOSUserLists) List(ctx context.Context, opts metav1.ListOptions) (result *v1.VirtualMachineInstanceGuestOSUserListList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.VirtualMachineInstanceGuestOSUserListList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstanceguestosuserlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualMachineInstanceGuestOSUserLists.
func (c *virtualMachineInstanceGuestOSUserLists) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualmachineinstanceguestosuserlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a virtualMachineInstanceGuestOSUserList and creates it.  Returns the server's representation of the virtualMachineInstanceGuestOSUserList, and an error, if there is any.
func (c *virtualMachineInstanceGuestOSUserLists) Create(ctx context.Context, virtualMachineInstanceGuestOSUserList *v1.VirtualMachineInstanceGuestOSUserList, opts metav1.CreateOptions) (result *v1.VirtualMachineInstanceGuestOSUserList, err error) {
	result = &v1.VirtualMachineInstanceGuestOSUserList{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualmachineinstanceguestosuserlists").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstanceGuestOSUserList).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a virtualMachineInstanceGuestOSUserList and updates it. Returns the server's representation of the virtualMachineInstanceGuestOSUserList, and an error, if there is any.
func (c *virtualMachineInstanceGuestOSUserLists) Update(ctx context.Context, virtualMachineInstanceGuestOSUserList *v1.VirtualMachineInstanceGuestOSUserList, opts metav1.UpdateOptions) (result *v1.VirtualMachineInstanceGuestOSUserList, err error) {
	result = &v1.VirtualMachineInstanceGuestOSUserList{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualmachineinstanceguestosuserlists").
		Name(virtualMachineInstanceGuestOSUserList.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(virtualMachineInstanceGuestOSUserList).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the virtualMachineInstanceGuestOSUserList and deletes it. Returns an error if one occurs.
func (c *virtualMachineInstanceGuestOSUserLists) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineinstanceguestosuserlists").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualMachineInstanceGuestOSUserLists) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualmachineinstanceguestosuserlists").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched virtualMachineInstanceGuestOSUserList.
func (c *virtualMachineInstanceGuestOSUserLists) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.VirtualMachineInstanceGuestOSUserList, err error) {
	result = &v1.VirtualMachineInstanceGuestOSUserList{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualmachineinstanceguestosuserlists").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}