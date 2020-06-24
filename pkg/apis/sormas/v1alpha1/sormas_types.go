package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/code-generator/pkg/namer"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SormasSpec defines the desired state of Sormas
type SormasSpec struct {
	Database struct {
		Image     string `json:"image"`
		Host      string `json:"host"`     // Done nil or ""  ?
		User      string `json:"user"`     // Done
		Password  string `json:"password"` // base64 // Done
		Name      string `json:"name"`     // Done
		AuditName string `json:"audit"`    // Done
		Size      string `json:"size"`
	} `json:"database"`
	Server struct {
		Image      string `json:"image"`    // Done
		ServerURL  string `json:"url"`      // Done
		DomainName string `json:"domain"`   // Done
		JvmMax     string `json:"jvmMax"`   // Done
		Version    string `json:"version"`  // Done
		DevMode    string `json:"dev"`      // Done
		Replicas   int32  `json:"replica"`  // Done
		Custom     bool   `json:"custom"`   // Done Deploy! TODO pvc
		Password   string `json:"password"` // base64 TODO init??
	} `json:"server"`
	Mail struct {
		MailHost   string `json:"host"`          // Done
		MailFrom   string `json:"from"`          // Done
		SenderAddr string `json:"senderAddress"` // Done
		SenderName string `json:"senderName"`    // Done
	} `json:"mail"`
	Config struct {
		Locale struct {
			Latitude  string `json:"latitude"`  // Done
			Longitude string `json:"longitude"` // Done
			Locale    string `json:"locale"`    // Done
			MapZoom   string `json:"mapZoom"`   // Done
			Timezone  string `json:"timezone"`  // Done
			GeoUUID   string `json:"geoUUID"`   // Done
		} `json:"local"`
		Epidprefix string `json:"epidprefix"` 	// Done
		Seperator  string `json:"seperator"`	// Done
		Ticket     string `json:"ticket"`
	} `json:"config"`
}

// SormasStatus defines the observed state of Sormas
type SormasStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Sormas is the Schema for the sormas API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=sormas,scope=Namespaced
// +kubebuilder:storageversion=
type Sormas struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SormasSpec   `json:"spec,omitempty"`
	Status SormasStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SormasList contains a list of Sormas
type SormasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Sormas `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Sormas{}, &SormasList{})
}
