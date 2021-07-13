package internal

import (
	apicorev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	appsv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

// KubernetesClient
type KubernetesClient struct {
	clientset *kubernetes.Clientset
	namespace string
}

// initialize KubernetesClient
func (kc *KubernetesClient) Init() {
	clientSet, err := kubernetes.NewForConfig(KubeConfigFromDir())
	if err != nil {
		panic(err)
	}

	kc.clientset = clientSet
	kc.namespace = apicorev1.NamespaceDefault
}

// gets KubernetesClient namespace
func (kc *KubernetesClient) GetNamespace() string {
	return string(kc.namespace)
}

// DeploymentClientSet - clientset to create, update, read, list & delete deployment
func (kc *KubernetesClient) DeploymentClientSet(namespace string) appsv1.DeploymentInterface {
	return kc.clientset.AppsV1().Deployments(namespace)
}

// PodClientSet - clientset to create, update, read, list & delete pod
func (kc *KubernetesClient) PodClientSet(namespace string) corev1.PodInterface {
	return kc.clientset.CoreV1().Pods(namespace)
}

// ConfigMapClientSet - clientset to create, update, read, list & delete configmap
func (kc *KubernetesClient) ConfigMapClientSet(namespace string) corev1.ConfigMapInterface {
	return kc.clientset.CoreV1().ConfigMaps(namespace)
}

// SecretClientSet - clientset to create, update, read, list & delete secret
func (kc *KubernetesClient) SecretClientSet(namespace string) corev1.SecretInterface {
	return kc.clientset.CoreV1().Secrets(namespace)
}
