package frontendNativeResource

import (
	quickopsv1Controllerapi "aasourav/fullstackdeploymentoperator/api/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"

	corev1 "k8s.io/api/core/v1"
)

func FrontendService(deploymentData quickopsv1Controllerapi.FullStackDeploy) *corev1.Service {
	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentData.Name + "-frontend-service",
			Namespace: deploymentData.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: deploymentData.APIVersion,
					Kind:       deploymentData.Kind,
					Name:       deploymentData.Name,
					Controller: ptr.To(true),
					UID:        deploymentData.UID,
				},
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"apps": "quickopsfe",
			},
			Ports: []corev1.ServicePort{
				{
					Protocol:   corev1.ProtocolTCP,
					TargetPort: intstr.FromString("fe"),
					Port:       deploymentData.Spec.FrontendPort,
				},
			},
			Type: corev1.ServiceTypeClusterIP,
		},
	}

	return service
}
