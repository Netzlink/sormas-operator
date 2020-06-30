package sormas

import (
	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/apis/core"
)

func newPVCForCR(cr *sormasv1alpha1.Sormas) *core.PersistentVolumeClaim {
	sormasLabels := map[string]string{
		"app": cr.Name,
	}
	secretMetaData := metav1.ObjectMeta{
		Name:      cr.Name + "-sormas",
		Namespace: cr.Namespace,
		Labels:    sormasLabels,
	}
	return &core.PersistentVolumeClaim{
		ObjectMeta: secretMetaData,
		Spec: core.PersistentVolumeClaimSpec{
			AccessModes: []core.PersistentVolumeAccessMode{
				core.PersistentVolumeAccessMode("RWO"),
			},
			Resources: core.ResourceRequirements{
				Requests: core.ResourceList(
					map[core.ResourceName]resource.Quantity{
						"storage": *resource.NewQuantity(1*1000000000, resource.DecimalSI),
					},
				),
			},
		},
	}
}
