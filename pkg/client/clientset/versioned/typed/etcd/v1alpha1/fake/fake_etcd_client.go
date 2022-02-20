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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
	v1alpha1 "tkestack.io/kstone-etcd-operator/pkg/client/clientset/versioned/typed/etcd/v1alpha1"
)

type FakeEtcdV1alpha1 struct {
	*testing.Fake
}

func (c *FakeEtcdV1alpha1) EtcdClusters(namespace string) v1alpha1.EtcdClusterInterface {
	return &FakeEtcdClusters{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeEtcdV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
