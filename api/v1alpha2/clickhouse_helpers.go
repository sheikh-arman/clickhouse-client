package v1alpha2

import (
	"fmt"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	meta_util "kmodules.xyz/client-go/meta"
	appcat "kmodules.xyz/custom-resources/apis/appcatalog/v1alpha1"
	"kubedb.dev/apimachinery/apis/kubedb"
	api "kubedb.dev/apimachinery/apis/kubedb/v1alpha2"
	"strings"
)

const (
	ClickHouseKeeperPort    = 9181
	ClickHouseDefaultHTTP   = 8123
	ClickHouseDefaultTLS    = 8443
	ClickHouseNativeTCP     = 9000
	ClickHouseNativeTLS     = 9440
	ClickhousePromethues    = 9363
	ClickHouseVolumeData    = "data"
	ClickHouseDataDir       = "/var/lib/clickhouse"
	ClickHouseContainerName = "clickhouse"
)

type ClickhouseApp struct {
	*ClickHouse
}

func (r *ClickHouse) AppBindingMeta() appcat.AppBindingMeta {
	return &ClickhouseApp{r}
}
func (r ClickhouseApp) Name() string {
	return r.ClickHouse.Name
}

func (r ClickhouseApp) Type() appcat.AppType {
	return appcat.AppType(fmt.Sprintf("%s/%s", kubedb.GroupName, ResourceSingularClickhouse))
}

// Owner returns owner reference to resources
func (r *ClickHouse) Owner() *meta.OwnerReference {
	return meta.NewControllerRef(r, SchemeGroupVersion.WithKind(r.ResourceKind()))
}

func (r *ClickHouse) ResourceKind() string {
	return ResourceKindClickhouse
}

func (r *ClickHouse) ServiceName() string {
	return r.OffshootName()
}

func (r *ClickHouse) OffshootName() string {
	return r.Name
}

func (r *ClickHouse) OffshootLabels() map[string]string {
	return r.offshootLabels(r.OffshootSelectors(), nil)
}

func (r *ClickHouse) offshootLabels(selector, override map[string]string) map[string]string {
	selector[meta_util.ComponentLabelKey] = api.ComponentDatabase
	return meta_util.FilterKeys(kubedb.GroupName, selector, meta_util.OverwriteKeys(nil, r.Labels, override))
}

func (r *ClickHouse) OffshootSelectors(extraSelectors ...map[string]string) map[string]string {
	selector := map[string]string{
		meta_util.NameLabelKey:      r.ResourceFQN(),
		meta_util.InstanceLabelKey:  r.Name,
		meta_util.ManagedByLabelKey: kubedb.GroupName,
	}
	return meta_util.OverwriteKeys(selector, extraSelectors...)
}

func (r *ClickHouse) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", r.ResourcePlural(), kubedb.GroupName)
}

func (r *ClickHouse) ResourcePlural() string {
	return ResourcePluralClickhouse
}

func (r *ClickHouse) GoverningServiceName() string {
	return meta_util.NameWithSuffix(r.ServiceName(), "pods")
}

func (r *ClickHouse) GetAuthSecretName() string {
	if r.Spec.AuthSecret != nil && r.Spec.AuthSecret.Name != "" {
		return r.Spec.AuthSecret.Name
	}
	return r.DefaultUserCredSecretName("admin")
}

func (r *ClickHouse) DefaultUserCredSecretName(username string) string {
	return meta_util.NameWithSuffix(r.Name, strings.ReplaceAll(fmt.Sprintf("%s-cred", username), "_", "-"))
}

func (r *ClickHouse) PVCName(alias string) string {
	return meta_util.NameWithSuffix(r.Name, alias)
}

func (r *ClickHouse) PetSetName() string {
	return r.OffshootName()
}

func (r *ClickHouse) PodLabels(extraLabels ...map[string]string) map[string]string {
	return r.offshootLabels(meta_util.OverwriteKeys(r.OffshootSelectors(), extraLabels...), r.Spec.PodTemplate.Labels)
}

func (r *ClickHouse) GetConnectionScheme() string {
	scheme := "http"
	if r.Spec.EnableSSL {
		scheme = "https"
	}
	return scheme
}
