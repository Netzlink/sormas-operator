package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/code-generator/pkg/namer"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// SormasSpec defines the desired state of Sormas
type SormasSpec struct {
	//TODO: Types?
	Database struct {
		Host 		string	`json:"host"`
		User 		string	`json:"user"`
		SecretName 	string	`json:"secretName"`
		Name		string	`json:"name"`
		AuditName	string	`json:"auditName"`
	} `json:"database"`
	Server struct {
		ServerUrl	string	`json:"url"`
		DomainName	string	`json:"domain"`
		JvmMax		string	`json:"jvmMax"`
		Version		string 	`json:"version"`
		DevMode		string 	`json:"devMode"` //TODO: Type
	} `json:"server"`
	Mail struct {
		MailHost	string 	`json:"host"`
		MailFrom	string 	`json:"from"`
		SenderAddr	string 	`json:"senderAddress"`
		SenderName	string 	`json:"senderName"`
	} `json:"mail"`
	Config struct {
		Locale struct {
			Latitude	string 	`json:"latitude"`
			Longitude	string 	`json:"longitude"`
			Locale 		string 	`json:"locale"`
			MapZoom		string 	`json:"mapZoom"`
			Timezone	string 	`json:"timezone"` // as TZ
			GeoUUID		string 	`json:"geoUUID"`
		} `json:"local"`
		Epidprefix	string 	`json:"epidprefix"`
		Seperator	string 	`json:"seperator"`
	} `json:"config"`
}

// SormasStatus defines the observed state of Sormas
type SormasStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Sormas is the Schema for the sormas API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=sormas,scope=Namespaced
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
