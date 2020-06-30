package sormas

import (
	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	routev1 "github.com/openshift/api/route/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func newRouteForCR(cr *sormasv1alpha1.Sormas) *routev1.Route {
	sormasLabels := map[string]string{
		"app": cr.Name,
	}
	routeMetaData := metav1.ObjectMeta{
		Name:      cr.Name + "-sormas",
		Namespace: cr.Namespace,
		Labels:    sormasLabels,
	}
	return &routev1.Route{
		ObjectMeta: routeMetaData,
		Spec: routev1.RouteSpec{
			Port: &routev1.RoutePort{
				TargetPort: intstr.FromInt(6080),
			},
			TLS: &routev1.TLSConfig{
				Termination: routev1.TLSTerminationEdge,
				InsecureEdgeTerminationPolicy: routev1.InsecureEdgeTerminationPolicyRedirect,
			},
			Host: cr.Spec.Server.ServerURL,
			To: routev1.RouteTargetReference{
				Kind: "Service",
				Name: cr.Name + "-sormas",
			},
		},
	}
}