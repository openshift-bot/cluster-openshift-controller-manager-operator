// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KubeStorageVersionMigratorLister helps list KubeStorageVersionMigrators.
type KubeStorageVersionMigratorLister interface {
	// List lists all KubeStorageVersionMigrators in the indexer.
	List(selector labels.Selector) (ret []*v1.KubeStorageVersionMigrator, err error)
	// Get retrieves the KubeStorageVersionMigrator from the index for a given name.
	Get(name string) (*v1.KubeStorageVersionMigrator, error)
	KubeStorageVersionMigratorListerExpansion
}

// kubeStorageVersionMigratorLister implements the KubeStorageVersionMigratorLister interface.
type kubeStorageVersionMigratorLister struct {
	indexer cache.Indexer
}

// NewKubeStorageVersionMigratorLister returns a new KubeStorageVersionMigratorLister.
func NewKubeStorageVersionMigratorLister(indexer cache.Indexer) KubeStorageVersionMigratorLister {
	return &kubeStorageVersionMigratorLister{indexer: indexer}
}

// List lists all KubeStorageVersionMigrators in the indexer.
func (s *kubeStorageVersionMigratorLister) List(selector labels.Selector) (ret []*v1.KubeStorageVersionMigrator, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.KubeStorageVersionMigrator))
	})
	return ret, err
}

// Get retrieves the KubeStorageVersionMigrator from the index for a given name.
func (s *kubeStorageVersionMigratorLister) Get(name string) (*v1.KubeStorageVersionMigrator, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("kubestorageversionmigrator"), name)
	}
	return obj.(*v1.KubeStorageVersionMigrator), nil
}
