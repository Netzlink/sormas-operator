package sormas

import (
	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

func newSecretForCR(cr *sormasv1alpha1.Sormas) *v1.Secret {
	sormasLabels := map[string]string{
		"app": cr.Name,
	}
	secretMetaData := metav1.ObjectMeta{
		Name:      cr.Name + "-sormas",
		Namespace: cr.Namespace,
		Labels:    sormasLabels,
	}
	return &v1.Secret{
		ObjectMeta: secretMetaData,
		Data: map[string][]byte{
			"password": []byte(cr.Spec.Database.Password),
		},
	}
}