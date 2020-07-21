package sormas

import (
	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/core/v1"
)

func newPVCForCR(cr *sormasv1alpha1.Sormas) *v1.PersistentVolumeClaim {
	sormasLabels := map[string]string{
		"app": cr.Name,
	}
	secretMetaData := metav1.ObjectMeta{
		Name:      cr.Name + "-sormas",
		Namespace: cr.Namespace,
		Labels:    sormasLabels,
	}
	return &v1.PersistentVolumeClaim{
		ObjectMeta: secretMetaData,
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: []v1.PersistentVolumeAccessMode{
				v1.PersistentVolumeAccessMode("RWO"),
			},
			Resources: v1.ResourceRequirements{
				Requests: v1.ResourceList(
					map[v1.ResourceName]resource.Quantity{
						"storage": *resource.NewQuantity(1*1000000000, resource.DecimalSI),
					},
				),
			},
		},
	}
}
