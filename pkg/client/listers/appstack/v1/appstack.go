/*
Copyright The Kubernetes Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/kenazk/genesis-operator/pkg/apis/appstack/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AppStackLister helps list AppStacks.
type AppStackLister interface {
	// List lists all AppStacks in the indexer.
	List(selector labels.Selector) (ret []*v1.AppStack, err error)
	// AppStacks returns an object that can list and get AppStacks.
	AppStacks(namespace string) AppStackNamespaceLister
	AppStackListerExpansion
}

// appStackLister implements the AppStackLister interface.
type appStackLister struct {
	indexer cache.Indexer
}

// NewAppStackLister returns a new AppStackLister.
func NewAppStackLister(indexer cache.Indexer) AppStackLister {
	return &appStackLister{indexer: indexer}
}

// List lists all AppStacks in the indexer.
func (s *appStackLister) List(selector labels.Selector) (ret []*v1.AppStack, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AppStack))
	})
	return ret, err
}

// AppStacks returns an object that can list and get AppStacks.
func (s *appStackLister) AppStacks(namespace string) AppStackNamespaceLister {
	return appStackNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppStackNamespaceLister helps list and get AppStacks.
type AppStackNamespaceLister interface {
	// List lists all AppStacks in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.AppStack, err error)
	// Get retrieves the AppStack from the indexer for a given namespace and name.
	Get(name string) (*v1.AppStack, error)
	AppStackNamespaceListerExpansion
}

// appStackNamespaceLister implements the AppStackNamespaceLister
// interface.
type appStackNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppStacks in the indexer for a given namespace.
func (s appStackNamespaceLister) List(selector labels.Selector) (ret []*v1.AppStack, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AppStack))
	})
	return ret, err
}

// Get retrieves the AppStack from the indexer for a given namespace and name.
func (s appStackNamespaceLister) Get(name string) (*v1.AppStack, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("appstack"), name)
	}
	return obj.(*v1.AppStack), nil
}
