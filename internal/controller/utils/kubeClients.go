package utils

import (
	"context"
	"flag"
	"fmt"

	"github.com/go-logr/logr"
	"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

type KubeClients struct {
	CRDStructuredClientSet   *clientset.Clientset
	KubernetesClientSet      *kubernetes.Clientset
	CRDUnstructuredClientSet *dynamic.DynamicClient
}

var kubeconfig = flag.String("kubeconfig", "/home/aasourav/.kube/config", "location to your kubeconfig file")

func GetCRDStructuredClientSet() (*clientset.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("[SYS]: Trying to Get Config from InCluster config..")
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Println("[ERROR]: Getting in-cluster configuration...")
		}
	}

	client, err := clientset.NewForConfig(config)
	if err != nil {
		return client, err
	}
	return client, nil
}

func GetKubernetesClientSet() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("[SYS]: Trying to Get Config from InCluster config..")
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Println("[ERROR]: Getting in-cluster configuration...")
		}
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return client, err
	}
	return client, nil
}

func GetCRDUnstructuredClientSet() (*dynamic.DynamicClient, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println("[SYS]: Trying to Get Config from InCluster config..")
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Println("[ERROR]: Getting in-cluster configuration...")
		}
	}
	unstructuredClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return unstructuredClient, err
	}
	return unstructuredClient, nil
}

func LoadCRDs(l logr.Logger) {
	crdStructuredClientSet, err := GetCRDStructuredClientSet()
	if crdStructuredClientSet == nil {
		l.Error(err, "Failed to get structured CRD clientSet")
	}
	crdStructuredList, err := crdStructuredClientSet.ApiextensionsV1().CustomResourceDefinitions().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		l.Error(err, "CRD structured list not found")
	}

	// 3. get new empty schema holder
	scheme := runtime.NewScheme()

	// 4. loop over all the crd and add to the schema
	for _, crd := range crdStructuredList.Items {
		for _, v := range crd.Spec.Versions {
			//fmt.Printf("GROUP = %s        VERSION = %s     KIND = %s\n", crd.Spec.Group, v.Name, crd.Spec.Names.Kind)
			scheme.AddKnownTypeWithName(
				schema.GroupVersionKind{
					Group:   crd.Spec.Group,
					Version: v.Name,
					Kind:    crd.Spec.Names.Kind,
				},
				&unstructured.Unstructured{},
			)
		}
	}
}

func GetAllClients() (KubeClients, error) {
	CrdStructuredClient, err := GetCRDStructuredClientSet()
	if err != nil {
		return KubeClients{}, err
	}
	KubernetesClient, err := GetKubernetesClientSet()
	if err != nil {
		return KubeClients{}, err
	}
	UnstructuredClient, err := GetCRDUnstructuredClientSet()
	if err != nil {
		return KubeClients{}, err
	}
	return KubeClients{
		CRDStructuredClientSet:   CrdStructuredClient,
		KubernetesClientSet:      KubernetesClient,
		CRDUnstructuredClientSet: UnstructuredClient,
	}, nil

}
