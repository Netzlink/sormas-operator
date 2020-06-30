package sormas

import (
	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func newStatefulSetForCR(cr *sormasv1alpha1.Sormas) *appsv1.StatefulSet {
	sormasLabels := map[string]string{
		"app": cr.Name,
	}
	statefulsetMetaData := metav1.ObjectMeta{
		Name:      cr.Name + "-sormas",
		Namespace: cr.Namespace,
		Labels:    sormasLabels,
	}
	replicas := int32(1)
	return &appsv1.StatefulSet{
		ObjectMeta: statefulsetMetaData,
		Spec: appsv1.StatefulSetSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: sormasLabels,
			},
			ServiceName: cr.Name + "-db",
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						v1.Container{
							Name:  "postgres",
							Image: cr.Spec.Database.Image,
							VolumeMounts: []v1.VolumeMount{
								v1.VolumeMount{
									Name:      "data",
									MountPath: "/var/lib/postgresql/data",
								},
							},
							Env: []v1.EnvVar{
								v1.EnvVar{
									Name: "POSTGRES_PASSWORD",
									ValueFrom: &v1.EnvVarSource{
										SecretKeyRef: &v1.SecretKeySelector{
											LocalObjectReference: v1.LocalObjectReference{
												Name: cr.Name + "-sormas",
											},
											Key: "password",
										},
									},
								},
								v1.EnvVar{
									Name:  "DB_NAME",
									Value: cr.Spec.Database.Name,
								},
								v1.EnvVar{
									Name:  "DB_NAME_AUDIT",
									Value: cr.Spec.Database.AuditName,
								},
								v1.EnvVar{
									Name: "SORMAS_POSTGRES_PASSWORD",
									ValueFrom: &v1.EnvVarSource{
										SecretKeyRef: &v1.SecretKeySelector{
											LocalObjectReference: v1.LocalObjectReference{
												Name: cr.Name + "-sormas",
											},
											Key: "password",
										},
									},
								},
								v1.EnvVar{
									Name:  "SORMAS_POSTGRES_USER",
									Value: cr.Spec.Database.User,
								},
								v1.EnvVar{
									Name:  "TZ",
									Value: cr.Spec.Config.Locale.Timezone,
								},
							},
							Resources: v1.ResourceRequirements{
								Limits: v1.ResourceList(
									map[v1.ResourceName]resource.Quantity{
										"cpu":    *resource.NewMilliQuantity(225, resource.DecimalSI),
										"memory": *resource.NewMilliQuantity(525, resource.BinarySI),
									},
								),
								Requests: v1.ResourceList(
									map[v1.ResourceName]resource.Quantity{
										"cpu":    *resource.NewMilliQuantity(25, resource.DecimalSI),
										"memory": *resource.NewMilliQuantity(25, resource.BinarySI),
									},
								),
							},
							Ports: []v1.ContainerPort{
								v1.ContainerPort{
									Name:          "postgres",
									ContainerPort: 5432,
								},
							},
							LivenessProbe: &v1.Probe{
								Handler: v1.Handler{
									Exec: &v1.ExecAction{
										Command: []string{
											"psql",
											"-U",
											cr.Spec.Database.User,
											"-c",
											"SELECT 1;",
											cr.Spec.Database.Name,
										},
									},
								},
								InitialDelaySeconds: 30,
								TimeoutSeconds:      3,
								FailureThreshold:    3,
								PeriodSeconds:       30,
								SuccessThreshold:    1,
							},
						},
					},
				},
			},
			VolumeClaimTemplates: []v1.PersistentVolumeClaim{
				v1.PersistentVolumeClaim{
					ObjectMeta: metav1.ObjectMeta{
						Name: "data",
					},
					Spec: v1.PersistentVolumeClaimSpec{
						Resources: v1.ResourceRequirements{
							Requests: v1.ResourceList(
								map[v1.ResourceName]resource.Quantity{
									"storage": *resource.NewQuantity(cr.Spec.Database.Size, resource.DecimalSI),
								},
							),
						},
					},
				},
			},
		},
	}
}
