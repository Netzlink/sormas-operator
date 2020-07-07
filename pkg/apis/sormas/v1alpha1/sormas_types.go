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
		Host      string `json:"host"`     
		// database user
		User      string `json:"user"`   
		// database password
		Password  string `json:"password"` 
		// database name
		Name      string `json:"name"`
		// audit database name
		AuditName string `json:"audit"`
		// pvc size
		Size      int64  `json:"size"`
	} `json:"database"`
	// Server-config
	Server struct {
		// Baseimage for deployment
		Image      string `json:"image"`
		// url for the payara server
		ServerURL  string `json:"url"`  
		// sormas domain name
		DomainName string `json:"domain"` 
		// maximum jvm heap memory
		JvmMax     string `json:"jvmMax"`  
		// sormas version
		Version    string `json:"version"` 
		// development mode
		DevMode    string `json:"dev"`    
		// sormas server replicas
		Replicas   int32  `json:"replica"`  
		// custom mode for test
		Custom     bool   `json:"custom"`   
		// sormas admin password (not working rn)
		Password   string `json:"password"` 
	} `json:"server"`
	// Mail server config
	Mail struct {
		// Mail host 
		MailHost   string `json:"host"`         
		// Mail from
		MailFrom   string `json:"from"`        
		// Sender address
		SenderAddr string `json:"senderAddress"` 
		// Sender name
		SenderName string `json:"senderName"`  
	} `json:"mail"`
	// Sormas config
	Config struct {
		// Localization
		Locale struct {
			// Latitude
			Latitude  string `json:"latitude"` 
			// Longitude
			Longitude string `json:"longitude"`
			// Linux locale
			Locale    string `json:"locale"`   
			// OpenStreetmap zoom
			MapZoom   string `json:"mapZoom"`  
			// Timezone
			Timezone  string `json:"timezone"` 
			// GeoUUID
			GeoUUID   string `json:"geoUUID"`  
		} `json:"local"`
		// Prefix in database
		Epidprefix string `json:"epidprefix"` 
		// Seperator to use in CSV export
		Seperator  string `json:"seperator"`
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
