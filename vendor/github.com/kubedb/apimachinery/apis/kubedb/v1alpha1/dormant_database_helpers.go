package v1alpha1

import (
	crdutils "github.com/appscode/kutil/apiextensions/v1beta1"
	meta_util "github.com/appscode/kutil/meta"
	apiextensions "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

func (d DormantDatabase) OffshootSelectors() map[string]string {
	selector := map[string]string{
		LabelDatabaseName: d.Name,
	}
	switch {
	case d.Spec.Origin.Spec.Etcd != nil:
		selector[LabelDatabaseKind] = ResourceKindEtcd
	case d.Spec.Origin.Spec.Elasticsearch != nil:
		selector[LabelDatabaseKind] = ResourceKindElasticsearch
	case d.Spec.Origin.Spec.Memcached != nil:
		selector[LabelDatabaseKind] = ResourceKindMemcached
	case d.Spec.Origin.Spec.MongoDB != nil:
		selector[LabelDatabaseKind] = ResourceKindMongoDB
	case d.Spec.Origin.Spec.MySQL != nil:
		selector[LabelDatabaseKind] = ResourceKindMySQL
	case d.Spec.Origin.Spec.Postgres != nil:
		selector[LabelDatabaseKind] = ResourceKindPostgres
	case d.Spec.Origin.Spec.Redis != nil:
		selector[LabelDatabaseKind] = ResourceKindRedis
	}
	return selector
}

func (d DormantDatabase) OffshootLabels() map[string]string {
	return meta_util.FilterKeys(GenericKey, d.OffshootSelectors(), d.Spec.Origin.Labels)
}

func (d DormantDatabase) OffshootName() string {
	return d.Name
}

func (d DormantDatabase) ResourceShortCode() string {
	return ResourceCodeDormantDatabase
}

func (d DormantDatabase) ResourceKind() string {
	return ResourceKindDormantDatabase
}

func (d DormantDatabase) ResourceSingular() string {
	return ResourceSingularDormantDatabase
}

func (d DormantDatabase) ResourcePlural() string {
	return ResourcePluralDormantDatabase
}

func (d DormantDatabase) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crdutils.NewCustomResourceDefinition(crdutils.Config{
		Group:         SchemeGroupVersion.Group,
		Plural:        ResourcePluralDormantDatabase,
		Singular:      ResourceSingularDormantDatabase,
		Kind:          ResourceKindDormantDatabase,
		ShortNames:    []string{ResourceCodeDormantDatabase},
		Categories:    []string{"datastore", "kubedb", "appscode", "all"},
		ResourceScope: string(apiextensions.NamespaceScoped),
		Versions: []apiextensions.CustomResourceDefinitionVersion{
			{
				Name:    SchemeGroupVersion.Version,
				Served:  true,
				Storage: true,
			},
		},
		Labels: crdutils.Labels{
			LabelsMap: map[string]string{"app": "kubedb"},
		},
		SpecDefinitionName:      "github.com/kubedb/apimachinery/apis/kubedb/v1alpha1.DormantDatabase",
		EnableValidation:        false,
		GetOpenAPIDefinitions:   GetOpenAPIDefinitions,
		EnableStatusSubresource: EnableStatusSubresource,
		AdditionalPrinterColumns: []apiextensions.CustomResourceColumnDefinition{
			{
				Name:     "Status",
				Type:     "string",
				JSONPath: ".status.phase",
			},
			{
				Name:     "Age",
				Type:     "date",
				JSONPath: ".metadata.creationTimestamp",
			},
		},
	}, setNameSchema)
}

func (d *DormantDatabase) Migrate() {
	if d == nil {
		return
	}
	d.Spec.Origin.Spec.Elasticsearch.Migrate()
	d.Spec.Origin.Spec.Postgres.Migrate()
	d.Spec.Origin.Spec.MySQL.Migrate()
	d.Spec.Origin.Spec.MongoDB.Migrate()
	d.Spec.Origin.Spec.Redis.Migrate()
	d.Spec.Origin.Spec.Memcached.Migrate()
	d.Spec.Origin.Spec.Etcd.Migrate()
}
