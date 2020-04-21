/*
Copyright 2020 The Crossplane Authors.

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

package unstructured

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Wrapper returns the underlying *unstructured.Unstructured.
type Wrapper interface {
	GetUnstructured() *unstructured.Unstructured
}

// ListWrapper allows the *unstructured.UnstructuredList to be accessed.
type ListWrapper interface {
	GetUnstructuredList() *unstructured.UnstructuredList
}

// NewClient returns a client.Client that will operate on the
// underlying *unstructured.Unstructured if the object satisfies Wrapper
// or ListWrapper interfaces. It relies on *unstructured.Unstructured
// instead of simpler map[string]interface{} to avoid unnecessary copying.
func NewClient(c client.Client) client.Client {
	return &wrapperClient{
		kube: c,
	}
}

type wrapperClient struct {
	kube client.Client
}

// Get retrieves an obj for the given object key from the Kubernetes Cluster.
// obj must be a struct pointer so that obj can be updated with the response
// returned by the Server.
func (c *wrapperClient) Get(ctx context.Context, key client.ObjectKey, obj runtime.Object) error {
	if u, ok := obj.(Wrapper); ok {
		return c.kube.Get(ctx, key, u.GetUnstructured())
	}
	return c.kube.Get(ctx, key, obj)
}

// List retrieves list of objects for a given namespace and list options. On a
// successful call, Items field in the list will be populated with the
// result returned from the server.
func (c *wrapperClient) List(ctx context.Context, list runtime.Object, opts ...client.ListOption) error {
	if u, ok := list.(ListWrapper); ok {
		return c.kube.List(ctx, u.GetUnstructuredList(), opts...)
	}
	return c.kube.List(ctx, list, opts...)
}

// Create saves the object obj in the Kubernetes cluster.
func (c *wrapperClient) Create(ctx context.Context, obj runtime.Object, opts ...client.CreateOption) error {
	if u, ok := obj.(Wrapper); ok {
		return c.kube.Create(ctx, u.GetUnstructured(), opts...)
	}
	return c.kube.Create(ctx, obj, opts...)
}

// Delete deletes the given obj from Kubernetes cluster.
func (c *wrapperClient) Delete(ctx context.Context, obj runtime.Object, opts ...client.DeleteOption) error {
	if u, ok := obj.(Wrapper); ok {
		return c.kube.Delete(ctx, u.GetUnstructured(), opts...)
	}
	return c.kube.Delete(ctx, obj, opts...)
}

// Update updates the given obj in the Kubernetes cluster. obj must be a
// struct pointer so that obj can be updated with the content returned by the Server.
func (c *wrapperClient) Update(ctx context.Context, obj runtime.Object, opts ...client.UpdateOption) error {
	if u, ok := obj.(Wrapper); ok {
		return c.kube.Update(ctx, u.GetUnstructured(), opts...)
	}
	return c.kube.Update(ctx, obj, opts...)
}

// Patch patches the given obj in the Kubernetes cluster. obj must be a
// struct pointer so that obj can be updated with the content returned by the Server.
func (c *wrapperClient) Patch(ctx context.Context, obj runtime.Object, patch client.Patch, opts ...client.PatchOption) error {
	if u, ok := obj.(Wrapper); ok {
		return c.kube.Patch(ctx, u.GetUnstructured(), patch, opts...)
	}
	return c.kube.Patch(ctx, obj, patch, opts...)
}

// DeleteAllOf deletes all objects of the given type matching the given options.
func (c *wrapperClient) DeleteAllOf(ctx context.Context, obj runtime.Object, opts ...client.DeleteAllOfOption) error {
	if u, ok := obj.(Wrapper); ok {
		return c.kube.DeleteAllOf(ctx, u.GetUnstructured(), opts...)
	}
	return c.kube.DeleteAllOf(ctx, obj, opts...)
}

func (c *wrapperClient) Status() client.StatusWriter {
	return &wrapperStatusClient{
		kube: c.kube.Status(),
	}
}

type wrapperStatusClient struct {
	kube client.StatusWriter
}

// Update updates the fields corresponding to the status subresource for the
// given obj. obj must be a struct pointer so that obj can be updated
// with the content returned by the Server.
func (c *wrapperStatusClient) Update(ctx context.Context, obj runtime.Object, opts ...client.UpdateOption) error {
	if u, ok := obj.(Wrapper); ok {
		return c.kube.Update(ctx, u.GetUnstructured(), opts...)
	}
	return c.kube.Update(ctx, obj, opts...)
}

// Patch patches the given object's subresource. obj must be a struct
// pointer so that obj can be updated with the content returned by the
// Server.
func (c *wrapperStatusClient) Patch(ctx context.Context, obj runtime.Object, patch client.Patch, opts ...client.PatchOption) error {
	if u, ok := obj.(Wrapper); ok {
		return c.kube.Patch(ctx, u.GetUnstructured(), patch, opts...)
	}
	return c.kube.Patch(ctx, obj, patch, opts...)
}
