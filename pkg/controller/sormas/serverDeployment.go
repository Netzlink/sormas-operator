package sormas

import (
	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// newPodForCR returns a busybox pod with the same name/namespace as the cr
func newDeploymentForCR(cr *sormasv1alpha1.Sormas) *appsv1.Deployment {
	sormasLabels := map[string]string{
		"app": cr.Name,
	}
	serverMetaData := metav1.ObjectMeta{
		Name:      cr.Name + "-sormas-server",
		Namespace: cr.Namespace,
		Labels:    sormasLabels,
	}
	return &appsv1.Deployment{
		ObjectMeta: serverMetaData,
		Spec: appsv1.DeploymentSpec{
			Replicas: &cr.Spec.Server.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: sormasLabels,
			},
			Template: v1.PodTemplateSpec{
				ObjectMeta: serverMetaData,
				Spec: v1.PodSpec{
					Containers: []v1.Container{
						v1.Container{
							Name: "sormas-server",
							Image: cr.Spec.Server.Image,
							Ports: []v1.ContainerPort{
								v1.ContainerPort{
									Name: "web",
									ContainerPort: 6080,
									Protocol: "TCP",
								},
							},
							LivenessProbe: &v1.Probe{
								InitialDelaySeconds: 300,
								FailureThreshold: 3,
								TimeoutSeconds: 3,
								SuccessThreshold: 1,
								Handler: v1.Handler{
									HTTPGet: &v1.HTTPGetAction{
										Path: "/sormas-ui/login",
										Scheme: "HTTP",
										Host: "localhost",
										Port: intstr.FromInt(6080),
									},
								},
							},
							Env: []v1.EnvVar{
								v1.EnvVar{
									Name: "DB_HOST",
									Value: func(cr *sormasv1alpha1.Sormas) string {
										// check if host is empty then generate service of postgres 
										if cr.Spec.Database.Host != ""  {
											return cr.Spec.Database.Host
										}
										return cr.Name + "-db"
									}(cr),
								},
								v1.EnvVar{
									Name: "SORMAS_POSTGRES_USER",
									Value: cr.Spec.Database.User,
								},
								v1.EnvVar{
									Name: "SORMAS_POSTGRES_PASSWORD",
									ValueFrom: &v1.EnvVarSource{
										SecretKeyRef: &v1.SecretKeySelector{
											LocalObjectReference: v1.LocalObjectReference{
												Name: cr.Name + "-config",
											},
											Key: "database",
										},
									},
								},
								v1.EnvVar{
									Name: "DB_NAME",
									Value: cr.Spec.Database.Name,
								},
								v1.EnvVar{
									Name: "DB_NAME_AUDIT",
									Value: cr.Spec.Database.AuditName,
								},
								v1.EnvVar{
									Name: "SORMAS_SERVER_URL",
									Value: cr.Spec.Server.ServerURL,
								},
								v1.EnvVar{
									Name: "DOMAIN_NAME",
									Value: cr.Spec.Server.DomainName,
								},
								v1.EnvVar{
									Name: "JVM_MAX",
									Value: cr.Spec.Server.JvmMax,
								},
								v1.EnvVar{
									Name: "SORMAS_VERSION",
									Value: cr.Spec.Server.Version,
								},
								v1.EnvVar{
									Name: "DEVMODE",
									Value: cr.Spec.Server.DevMode,
								},
								v1.EnvVar{
									Name: "MAIL_HOST",
									Value: cr.Spec.Mail.MailHost,
								},
								v1.EnvVar{
									Name: "MAIL_FROM",
									Value: cr.Spec.Mail.MailFrom,
								},
								v1.EnvVar{
									Name: "EMAIL_SENDER_ADDRESS",
									Value: cr.Spec.Mail.SenderAddr,
								},
								v1.EnvVar{
									Name: "EMAIL_SENDER_NAME",
									Value: cr.Spec.Mail.SenderName,
								},
								v1.EnvVar{
									Name: "LATITUDE",
									Value: cr.Spec.Config.Locale.Latitude,
								},
								v1.EnvVar{
									Name: "LONGITUDE",
									Value: cr.Spec.Config.Locale.Longitude,
								},
								v1.EnvVar{
									Name: "LOCALE",
									Value: cr.Spec.Config.Locale.Locale,
								},
								v1.EnvVar{
									Name: "MapZoom",
									Value: cr.Spec.Config.Locale.MapZoom,
								},
								v1.EnvVar{
									Name: "TZ",
									Value: cr.Spec.Config.Locale.Timezone,
								},
								v1.EnvVar{
									Name: "GEO_UUID",
									Value: cr.Spec.Config.Locale.GeoUUID,
								},
								v1.EnvVar{
									Name: "EPIDPREFIX",
									Value: cr.Spec.Config.Epidprefix,
								},
								v1.EnvVar{
									Name: "SEPERATOR",
									Value: cr.Spec.Config.Seperator,
								},
							},
							VolumeMounts: func(cr *sormasv1alpha1.Sormas) []v1.VolumeMount {
								if cr.Spec.Server.Custom {
									return []v1.VolumeMount{
										v1.VolumeMount{
											Name: "custom",
											MountPath: "/opt/sormas/custom",
										},
									}
								}
								return []v1.VolumeMount{}
							}(cr),
						},
					},
					Volumes: []v1.Volume{
						v1.Volume{
							Name: "custom",
							VolumeSource: v1.VolumeSource{
								PersistentVolumeClaim: &v1.PersistentVolumeClaimVolumeSource{
									ClaimName: cr.Name + "-custom",
								},
							},
						},
					},
				},
			},
		},
	}
}