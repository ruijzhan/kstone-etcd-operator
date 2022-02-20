/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2023 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "tkestack.io/kstone-etcd-operator/pkg/apis/etcd/v1alpha1"
)

// EtcdClusterLister helps list EtcdClusters.
// All objects returned here must be treated as read-only.
type EtcdClusterLister interface {
	// List lists all EtcdClusters in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EtcdCluster, err error)
	// EtcdClusters returns an object that can list and get EtcdClusters.
	EtcdClusters(namespace string) EtcdClusterNamespaceLister
	EtcdClusterListerExpansion
}

// etcdClusterLister implements the EtcdClusterLister interface.
type etcdClusterLister struct {
	indexer cache.Indexer
}

// NewEtcdClusterLister returns a new EtcdClusterLister.
func NewEtcdClusterLister(indexer cache.Indexer) EtcdClusterLister {
	return &etcdClusterLister{indexer: indexer}
}

// List lists all EtcdClusters in the indexer.
func (s *etcdClusterLister) List(selector labels.Selector) (ret []*v1alpha1.EtcdCluster, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EtcdCluster))
	})
	return ret, err
}

// EtcdClusters returns an object that can list and get EtcdClusters.
func (s *etcdClusterLister) EtcdClusters(namespace string) EtcdClusterNamespaceLister {
	return etcdClusterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EtcdClusterNamespaceLister helps list and get EtcdClusters.
// All objects returned here must be treated as read-only.
type EtcdClusterNamespaceLister interface {
	// List lists all EtcdClusters in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EtcdCluster, err error)
	// Get retrieves the EtcdCluster from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EtcdCluster, error)
	EtcdClusterNamespaceListerExpansion
}

// etcdClusterNamespaceLister implements the EtcdClusterNamespaceLister
// interface.
type etcdClusterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EtcdClusters in the indexer for a given namespace.
func (s etcdClusterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.EtcdCluster, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.EtcdCluster))
	})
	return ret, err
}

// Get retrieves the EtcdCluster from the indexer for a given namespace and name.
func (s etcdClusterNamespaceLister) Get(name string) (*v1alpha1.EtcdCluster, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("etcdcluster"), name)
	}
	return obj.(*v1alpha1.EtcdCluster), nil
}
