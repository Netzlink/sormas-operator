package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	// "k8s.io/code-generator/pkg/namer"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SormasSpec defines the desired state of Sormas
type SormasSpec struct {
	// Databaseconfig
	Database struct {
		// Baseimage for database deployment
		Image     string `json:"image"`
		// Host of external database
		Host      string `json:"host"`     // Done nil or ""  ?
		// database user
		User      string `json:"user"`     // Done
		// database password
		Password  string `json:"password"` // base64 // Done
		// database name
		Name      string `json:"name"`     // Done
		// audit database name
		AuditName string `json:"audit"`    // Done
		// pvc size
		Size      int64  `json:"size"`
	} `json:"database"`
	// Server-config
	Server struct {
		// Baseimage for deployment
		Image      string `json:"image"`    // Done
		// url for the payara server
		ServerURL  string `json:"url"`      // Done
		// sormas domain name
		DomainName string `json:"domain"`   // Done
		// maximum jvm heap memory
		JvmMax     string `json:"jvmMax"`   // Done
		// sormas version
		Version    string `json:"version"`  // Done
		// development mode
		DevMode    string `json:"dev"`      // Done
		// sormas server replicas
		Replicas   int32  `json:"replica"`  // Done
		// custom mode for test
		Custom     bool   `json:"custom"`   // Done Deploy! TODO pvc
		// sormas admin password (not working rn)
		Password   string `json:"password"` // base64 TODO init??
	} `json:"server"`
	// Mail server config
	Mail struct {
		// Mail host 
		MailHost   string `json:"host"`          // Done
		// Mail from
		MailFrom   string `json:"from"`          // Done
		// Sender address
		SenderAddr string `json:"senderAddress"` // Done
		// Sender name
		SenderName string `json:"senderName"`    // Done
	} `json:"mail"`
	// Sormas config
	Config struct {
		// Localization
		Locale struct {
			// Latitude
			Latitude  string `json:"latitude"`  // Done
			// Longitude
			Longitude string `json:"longitude"` // Done
			// Linux locale
			Locale    string `json:"locale"`    // Done
			// OpenStreetmap zoom
			MapZoom   string `json:"mapZoom"`   // Done
			// Timezone
			Timezone  string `json:"timezone"`  // Done
			// GeoUUID
			GeoUUID   string `json:"geoUUID"`   // Done
		} `json:"local"`
		// Prefix in database
		Epidprefix string `json:"epidprefix"` 	// Done
		// Seperator to use in CSV export
		Seperator  string `json:"seperator"`	// Done
		// Ticket for password generation
		Ticket     string `json:"ticket"`
	} `json:"config"`
}

// SormasStatus defines the observed state of Sormas
type SormasStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	Status string `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Sormas is the Schema for the sormas API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=sormas,scope=Namespaced
// +kubebuilder:storageversion
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
