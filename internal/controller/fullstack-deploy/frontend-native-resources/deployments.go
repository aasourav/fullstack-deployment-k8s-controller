package frontendNativeResource

import (
	quickopsv1Controllerapi "aasourav/fullstackdeploymentoperator/api/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func FrontendDeploymentResource(deploymentData quickopsv1Controllerapi.FullStackDeploy) *appsv1.Deployment {
	frontendEnv := []corev1.EnvVar{}
	if deploymentData.Spec.FrontendEnv != nil && len(deploymentData.Spec.FrontendEnv) > 0 {
		for key, value := range deploymentData.Spec.FrontendEnv {
			env := corev1.EnvVar{
				Name:  key,
				Value: value,
			}
			frontendEnv = append(frontendEnv, env)
		}
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentData.Name + "fe-deployment",
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
			Labels: map[string]string{
				"apps": "quickopsfe",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"apps": "quickopsfe",
				},
			},
			Replicas: &deploymentData.Spec.FrontendReplicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      deploymentData.Name + "fe-deployment",
					Namespace: deploymentData.Namespace,
					Labels: map[string]string{
						"apps": "quickopsfe",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentData.Name + "fe-deployment",
							Image: deploymentData.Spec.FrontendImage,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: deploymentData.Spec.FrontendPort,
									Protocol:      corev1.ProtocolTCP,
									Name:          "fe",
								},
							},
							Env: frontendEnv,
						},
					},
				},
			},
		},
	}

	return deployment
}
