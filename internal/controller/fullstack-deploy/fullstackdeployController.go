/*
Copyright 2025 Ahsan Amin.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"

	quickopsv1Controllerapi "aasourav/fullstackdeploymentoperator/api/v1"
	controllerUtil "aasourav/fullstackdeploymentoperator/internal/controller/utils"
)

// FullStackDeployReconciler reconciles a FullStackDeploy object
type FullStackDeployReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	controllerUtil.KubeClients
	Log logr.Logger
}

// +kubebuilder:rbac:groups=quickops.sand.tech,resources=fullstackdeploys,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=quickops.sand.tech,resources=fullstackdeploys/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=quickops.sand.tech,resources=fullstackdeploys/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the FullStackDeploy object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *FullStackDeployReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// l := r.Log.WithValues("reconciling frontend: ")

	fullStackDeploy := &quickopsv1Controllerapi.FullStackDeploy{}
	err := r.Get(context.TODO(), types.NamespacedName{Name: req.Name, Namespace: req.Namespace}, fullStackDeploy, &client.GetOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			fmt.Println("FullstackDeployError: ", err)
			return ctrl.Result{}, err
		} else {
			return ctrl.Result{}, nil
		}
	}

	// fullStackDeploy.Finalizers

	// Frontend Reconciller
	if err := r.frontendReconciler(fullStackDeploy); err != nil {
		return ctrl.Result{}, err
	}

	// Backend Reconciller
	if err := r.backendReconciler(fullStackDeploy); err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FullStackDeployReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(controller.Options{MaxConcurrentReconciles: 2}).
		For(&quickopsv1Controllerapi.FullStackDeploy{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Complete(r)
}
