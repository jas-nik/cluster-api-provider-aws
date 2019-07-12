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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	deprecatedv1alpha1 "sigs.k8s.io/cluster-api/pkg/apis/deprecated/v1alpha1"
	clientset "sigs.k8s.io/cluster-api/pkg/client/clientset_generated/clientset"
	internalinterfaces "sigs.k8s.io/cluster-api/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "sigs.k8s.io/cluster-api/pkg/client/listers_generated/deprecated/v1alpha1"
)

// MachineSetInformer provides access to a shared informer and lister for
// MachineSets.
type MachineSetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MachineSetLister
}

type machineSetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMachineSetInformer constructs a new informer for MachineSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineSetInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMachineSetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMachineSetInformer constructs a new informer for MachineSet type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMachineSetInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterDeprecatedV1alpha1().MachineSets(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterDeprecatedV1alpha1().MachineSets(namespace).Watch(options)
			},
		},
		&deprecatedv1alpha1.MachineSet{},
		resyncPeriod,
		indexers,
	)
}

func (f *machineSetInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMachineSetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *machineSetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&deprecatedv1alpha1.MachineSet{}, f.defaultInformer)
}

func (f *machineSetInformer) Lister() v1alpha1.MachineSetLister {
	return v1alpha1.NewMachineSetLister(f.Informer().GetIndexer())
}
