package sormas

import (
	"k8s.io/kubernetes/pkg/apis/core"
	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	routev1 "github.com/openshift/client-go/route/listeners/route/v1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/apis/core"
	"k8s.io/kubernetes/staging/src/k8s.io/client-go/informers/auditregistration/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
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