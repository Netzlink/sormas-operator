
package sormas

import(
	"k8s.io/kubernetes/pkg/apis/core"
	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func newServiceForCR(cr *sormasv1alpha1.Sormas) *core.Service {
	sormasLabels := map[string]string{
		"app": cr.Name,
	}
	serviceMetaData := metav1.ObjectMeta{
		Name:      cr.Name + "-sormas",
		Namespace: cr.Namespace,
		Labels:    sormasLabels,
	}
	return &core.Service{
		ObjectMeta: serviceMetaData,
		Spec: core.ServiceSpec{
			Type: core.ServiceTypeClusterIP,
			Ports: []core.ServicePort{
				core.ServicePort{
					Name: "web",
					Protocol: core.ProtocolTCP,
					Port: 6080,
					TargetPort: intstr.FromInt(6080),
				},
			},
			Selector: sormasLabels,
		},
	}
}