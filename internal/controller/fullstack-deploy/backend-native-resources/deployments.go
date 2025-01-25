package backendNativeResource

import (
	quickopsv1Controllerapi "aasourav/fullstackdeploymentoperator/api/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/utils/ptr"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

func BackendDeploymentResource(deploymentData quickopsv1Controllerapi.FullStackDeploy) *appsv1.Deployment {

	backendEnv := []corev1.EnvVar{}
	if deploymentData.Spec.BackendEnv != nil && len(deploymentData.Spec.BackendEnv) > 0 {
		for key, value := range deploymentData.Spec.BackendEnv {
			env := corev1.EnvVar{
				Name:  key,
				Value: value,
			}
			backendEnv = append(backendEnv, env)
		}
	}

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "be",
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
				"apps": "quickopsbe",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"apps": "quickopsbe",
				},
			},
			Replicas: &deploymentData.Spec.BackendReplicas,
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      deploymentData.Name + "be-deployment",
					Namespace: deploymentData.Namespace,
					Labels: map[string]string{
						"apps": "quickopsfe",
					},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  deploymentData.Name + "be-deployment",
							Image: deploymentData.Spec.BackendImage,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: deploymentData.Spec.BackendPort,
									Protocol:      corev1.ProtocolTCP,
									Name:          "be",
								},
							},
							Env: backendEnv,
						},
					},
				},
			},
		},
	}

	return deployment
}
