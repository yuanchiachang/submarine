/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/apache/submarine/submarine-cloud-v2/pkg/apis/submarine/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SubmarineLister helps list Submarines.
// All objects returned here must be treated as read-only.
type SubmarineLister interface {
	// List lists all Submarines in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Submarine, err error)
	// Submarines returns an object that can list and get Submarines.
	Submarines(namespace string) SubmarineNamespaceLister
	SubmarineListerExpansion
}

// submarineLister implements the SubmarineLister interface.
type submarineLister struct {
	indexer cache.Indexer
}

// NewSubmarineLister returns a new SubmarineLister.
func NewSubmarineLister(indexer cache.Indexer) SubmarineLister {
	return &submarineLister{indexer: indexer}
}

// List lists all Submarines in the indexer.
func (s *submarineLister) List(selector labels.Selector) (ret []*v1alpha1.Submarine, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Submarine))
	})
	return ret, err
}

// Submarines returns an object that can list and get Submarines.
func (s *submarineLister) Submarines(namespace string) SubmarineNamespaceLister {
	return submarineNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SubmarineNamespaceLister helps list and get Submarines.
// All objects returned here must be treated as read-only.
type SubmarineNamespaceLister interface {
	// List lists all Submarines in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Submarine, err error)
	// Get retrieves the Submarine from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Submarine, error)
	SubmarineNamespaceListerExpansion
}

// submarineNamespaceLister implements the SubmarineNamespaceLister
// interface.
type submarineNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Submarines in the indexer for a given namespace.
func (s submarineNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Submarine, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Submarine))
	})
	return ret, err
}

// Get retrieves the Submarine from the indexer for a given namespace and name.
func (s submarineNamespaceLister) Get(name string) (*v1alpha1.Submarine, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("submarine"), name)
	}
	return obj.(*v1alpha1.Submarine), nil
}
