package sormas

import (
	"k8s.io/kubernetes/pkg/apis/core"
	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newSecretForCR(cr *sormasv1alpha1.Sormas) *core.Secret {
	sormasLabels := map[string]string{
		"app": cr.Name,
	}
	secretMetaData := metav1.ObjectMeta{
		Name:      cr.Name + "-sormas",
		Namespace: cr.Namespace,
		Labels:    sormasLabels,
	}
	return &core.Secret{
		ObjectMeta: secretMetaData,
		Data: map[string][]byte{
			"password": []byte(cr.Spec.Database.Password),
		},
	}
}