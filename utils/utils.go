package utils

import (
	pluginTypes "github.com/argoproj/argo-rollouts/utils/plugin/types"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewKubeConfig() (*rest.Config, error) {
	config, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		return nil, pluginTypes.RpcError{ErrorString: err.Error()}
	}
	return config, nil
}

// func NewClientset(config *rest.Config) (*kubernetes.Clientset, error) {
// 	config, err := rest.InClusterConfig()
// 	if err != nil {
// 		return pluginTypes.RpcError{ErrorString: err.Error()}
// 	}

// 	clientset, err := kubernetes.NewForConfig(config)
// 	if err != nil {
// 		return pluginTypes.RpcError{ErrorString: err.Error()}
// 	}
// 	r.Clienset = clientset
// 	return clientset, nil
// }
