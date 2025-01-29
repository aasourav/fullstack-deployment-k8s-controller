package controller

import (
	quickopsv1Controllerapi "aasourav/fullstackdeploymentoperator/api/v1"
	backend "aasourav/fullstackdeploymentoperator/internal/controller/fullstack-deploy/backend-native-resources"
	"context"

	controllerUtils "aasourav/fullstackdeploymentoperator/internal/controller/utils"

	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (r *FullStackDeployReconciler) backendReconciler(fullStackDeploymentData *quickopsv1Controllerapi.FullStackDeploy) error {
	//=============== Deployment =======================
	// deployment := &appsv1.Deployment{}
	// r.Get(context.TODO(), types.NamespacedName{Name: fullStackDeploymentData.Name + "fe-deployment", Namespace: fullStackDeploymentData.Namespace}, deployment, &client.GetOptions{})

	if _, err := r.KubernetesClientSet.AppsV1().Deployments(fullStackDeploymentData.Namespace).Get(context.TODO(), fullStackDeploymentData.Name+"-backend", metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return nil
		}
		deployment := backend.BackendDeploymentResource(*fullStackDeploymentData)
		if err := r.Create(context.TODO(), deployment); err != nil {
			return err
		}
	}

	// ====================  Service ===========================

	if _, err := r.KubernetesClientSet.CoreV1().Services(fullStackDeploymentData.Namespace).Get(context.TODO(), fullStackDeploymentData.Name+"-backend-service", metav1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		service := backend.BackendService(*fullStackDeploymentData)
		if err := r.Create(context.TODO(), service); err != nil {
			return err
		}
	}

	// ====================  Ingress ==============================

	ingress := &networkingv1.Ingress{}
	ingress, err := r.KubernetesClientSet.NetworkingV1().Ingresses(fullStackDeploymentData.Namespace).Get(context.TODO(), fullStackDeploymentData.Name+"fullstack-ing", metav1.GetOptions{})
	if err != nil {
		if !errors.IsNotFound(err) {
			return err
		}
		ingress := backend.BackendIngressService(*fullStackDeploymentData)
		if err := r.Create(context.TODO(), ingress); err != nil {
			return err
		}
	} else {
		if isFound := controllerUtils.IsHTTPIngressPathExist(ingress.Spec.Rules[0].HTTP.Paths, "/be?(.*)"); isFound {
			return nil
		}
		ingress = backend.UpdateBackendIngressService(*fullStackDeploymentData, *ingress)
		if err := r.Update(context.TODO(), ingress); err != nil {
			return err
		}
	}

	return nil
}
