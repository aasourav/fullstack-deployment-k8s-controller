package backendNativeResource

import (
	quickopsv1Controllerapi "aasourav/fullstackdeploymentoperator/api/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"

	networkingv1 "k8s.io/api/networking/v1"
)

func UpdateBackendIngressService(deploymentData quickopsv1Controllerapi.FullStackDeploy, ingress networkingv1.Ingress) *networkingv1.Ingress {
	networkingv1Paths := ingress.Spec.Rules[0].HTTP.Paths
	backendPath := networkingv1.HTTPIngressPath{
		Path:     "/be/?(.*)",
		PathType: ptr.To(networkingv1.PathTypeImplementationSpecific),
		Backend: networkingv1.IngressBackend{
			Service: &networkingv1.IngressServiceBackend{
				Name: deploymentData.Name + "-backend-service",
				Port: networkingv1.ServiceBackendPort{
					Number: deploymentData.Spec.BackendPort,
				},
			},
		},
	}
	networkingv1Paths = append(networkingv1Paths, backendPath)
	ingress.Spec.Rules[0].HTTP.Paths = networkingv1Paths

	return &ingress
}

func BackendIngressService(deploymentData quickopsv1Controllerapi.FullStackDeploy) *networkingv1.Ingress {
	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentData.Name + "fullstack-ing",
			Namespace: deploymentData.Namespace,
			Annotations: map[string]string{
				"nginx.ingress.kubernetes.io/rewrite-target": "/$1",
				"nginx.ingress.kubernetes.io/use-regex":      "true",
			},
			OwnerReferences: []metav1.OwnerReference{
				{
					APIVersion: deploymentData.APIVersion,
					Kind:       deploymentData.Kind,
					Name:       deploymentData.Name,
					UID:        deploymentData.UID,
					Controller: ptr.To(true),
				},
			},
		},
		Spec: networkingv1.IngressSpec{
			IngressClassName: ptr.To("nginx"),
			Rules: []networkingv1.IngressRule{
				{
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/be/?(.*)",
									PathType: ptr.To(networkingv1.PathTypeImplementationSpecific),
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: deploymentData.Name + "-backend-service",
											Port: networkingv1.ServiceBackendPort{
												Number: deploymentData.Spec.BackendPort,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	return ingress
}
