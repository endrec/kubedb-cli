package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	ResourceCodeMemcachedVersion     = "mcversion"
	ResourceKindMemcachedVersion     = "MemcachedVersion"
	ResourceSingularMemcachedVersion = "memcachedversion"
	ResourcePluralMemcachedVersion   = "memcachedversions"
)

// +genclient
// +genclient:nonNamespaced
// +genclient:skipVerbs=updateStatus
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MemcachedVersion defines a Memcached database version.
type MemcachedVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MemcachedVersionSpec `json:"spec,omitempty"`
}

// MemcachedVersionSpec is the spec for memcached version
type MemcachedVersionSpec struct {
	// Version
	Version string `json:"version"`
	// Database Image
	DB MemcachedVersionDatabase `json:"db"`
	// Exporter Image
	Exporter MemcachedVersionExporter `json:"exporter"`
}

// MemcachedVersionDatabase is the Memcached Database image
type MemcachedVersionDatabase struct {
	Image string `json:"image"`
}

// MemcachedVersionExporter is the image for the Memcached exporter
type MemcachedVersionExporter struct {
	Image string `json:"image"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// MemcachedVersionList is a list of MemcachedVersions
type MemcachedVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MemcachedVersion CRD objects
	Items []MemcachedVersion `json:"items,omitempty"`
}
