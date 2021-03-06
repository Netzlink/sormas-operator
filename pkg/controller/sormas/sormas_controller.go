package sormas

import (
	"context"
	"fmt"

	sormasv1alpha1 "github.com/Netzlink/sormas-operator/pkg/apis/sormas/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubernetes/pkg/apis/core"
	networking "k8s.io/kubernetes/pkg/apis/networking"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_sormas")

/**k8s.io/apimachinery/pkg/labels"pi
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new Sormas Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileSormas{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("sormas-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource Sormas
	err = c.Watch(&source.Kind{Type: &sormasv1alpha1.Sormas{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource Pods and requeue the owner Sormas
	// Watching Deployment, PVC, Statefulset, Secret, Configmap, Ingress
	err = c.Watch(&source.Kind{Type: &appsv1.Deployment{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &sormasv1alpha1.Sormas{},
	})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &v1.PersistentVolumeClaim{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &sormasv1alpha1.Sormas{},
	})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &appsv1.StatefulSet{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &sormasv1alpha1.Sormas{},
	})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &v1.Secret{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &sormasv1alpha1.Sormas{},
	})
	if err != nil {
		return err
	}

	// TODO: make this shit working again
	err = c.Watch(&source.Kind{Type: &networking.Ingress{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &sormasv1alpha1.Sormas{},
	})
	if err != nil {
		return err
	}

	err = c.Watch(&source.Kind{Type: &core.Service{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &sormasv1alpha1.Sormas{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileSormas implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileSormas{}

// ReconcileSormas reconciles a Sormas object
type ReconcileSormas struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a Sormas object and makes changes based on the state read
// and what is in the Sormas.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileSormas) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling Sormas")

	// Fetch the Sormas instance
	instance := &sormasv1alpha1.Sormas{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// ignore on pause
	if instance.Spec.Paused {
		reqLogger.Info(fmt.Sprintf("Instance %s is paused in namespace %s", instance.Name, instance.Namespace))
		return reconcile.Result{}, nil
	}

	// Define a new Pod object
	// pod := newPodForCR(instance)
	sormasSecret := newSecretForCR(instance)
	sormasDeployment := newDeploymentForCR(instance)
	sormasService := newServiceForCR(instance)
	sormasPVC := newPVCForCR(instance)
	sormasStatefulSet := newStatefulSetForCR(instance)
	sormasIngress := newIngressForCR(instance)

	for _, obj := range []metav1.Object{
		sormasSecret,
		sormasDeployment,
		sormasService,
		sormasPVC,
		sormasStatefulSet,
		sormasIngress,
	} {
		if err := controllerutil.SetControllerReference(instance, obj, r.scheme); err != nil {
			return reconcile.Result{}, err
		}
	}

	// Check if this Pod already exists
	foundSecret := &v1.Secret{}
	foundDeployment := &appsv1.Deployment{}
	foundService := &core.Service{}
	foundPVC := &v1.PersistentVolumeClaim{}
	foundStatefulSet := &appsv1.StatefulSet{}
	foundIngress := &networking.Ingress{}

	err = r.client.Get(context.TODO(), types.NamespacedName{Name: sormasSecret.Name, Namespace: sormasSecret.Namespace}, foundSecret)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Secret", "Namespace", sormasSecret.Namespace, "Name", sormasSecret.Name)
		err = r.client.Create(context.TODO(), foundSecret)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Secret created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	err = r.client.Get(context.TODO(), types.NamespacedName{Name: sormasDeployment.Name, Namespace: sormasDeployment.Namespace}, foundDeployment)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Deployment", "Namespace", sormasDeployment.Namespace, "Name", sormasDeployment.Name)
		err = r.client.Create(context.TODO(), foundDeployment)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Deployment created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	err = r.client.Get(context.TODO(), types.NamespacedName{Name: sormasService.Name, Namespace: sormasService.Namespace}, foundService)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Service", "Namespace", sormasService.Namespace, "Name", sormasService.Name)
		err = r.client.Create(context.TODO(), foundService)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Service created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	err = r.client.Get(context.TODO(), types.NamespacedName{Name: sormasPVC.Name, Namespace: sormasPVC.Namespace}, foundPVC)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new PVC", "Namespace", sormasPVC.Namespace, "Name", sormasPVC.Name)
		err = r.client.Create(context.TODO(), foundPVC)
		if err != nil {
			return reconcile.Result{}, err
		}
		// PVC created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	err = r.client.Get(context.TODO(), types.NamespacedName{Name: sormasStatefulSet.Name, Namespace: sormasStatefulSet.Namespace}, foundStatefulSet)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new StatefulSet", "Namespace", sormasStatefulSet.Namespace, "Name", sormasStatefulSet.Name)
		err = r.client.Create(context.TODO(), foundStatefulSet)
		if err != nil {
			return reconcile.Result{}, err
		}
		// StatefulSet created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	err = r.client.Get(context.TODO(), types.NamespacedName{Name: sormasIngress.Name, Namespace: sormasIngress.Namespace}, foundIngress)
	if err != nil && errors.IsNotFound(err) {
		reqLogger.Info("Creating a new Ingress", "Namespace", sormasIngress.Namespace, "Name", sormasIngress.Name)
		err = r.client.Create(context.TODO(), foundIngress)
		if err != nil {
			return reconcile.Result{}, err
		}
		// Ingress created successfully - don't requeue
		return reconcile.Result{}, nil
	} else if err != nil {
		return reconcile.Result{}, err
	}

	// Everything already exists - don't requeue
	reqLogger.Info("Skip reconcile: Sormas already exists", "Namespace", foundSecret.Namespace, "Name", foundSecret.Name)
	return reconcile.Result{}, nil
}
