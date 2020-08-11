package sormas

import (
	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	networking "k8s.io/kubernetes/pkg/apis/networking"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func newIngressForCR(cr *sormasv1alpha1.Sormas) *routev1.Route {
	sormasLabels := map[string]string{
		"app": cr.Name,
	}
	ingressMetaData := metav1.ObjectMeta{
		Name:      cr.Name + "-sormas",
		Namespace: cr.Namespace,
		Labels:    sormasLabels,
	}
	return &networking.Ingress{
		ObjectMeta: ingressMetaData,
		Spec: networking.IngressSpec{
			Backend: &networking.IngressBackend{
				ServiceName: cr.Name + "-sormas",
				ServicePort: web,
			},
			Rules: []networking.IngressRule{
				networking.IngressRule{
					Host: cr.Spec.Server.ServerURL,
				},
			},
		},
	}
}