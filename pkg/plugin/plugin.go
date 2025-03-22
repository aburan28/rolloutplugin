package plugin

import (
	"fmt"

	"github.com/aburan28/rolloutplugin-controller/api/v1alpha1"
	pluginTypes "github.com/aburan28/rolloutplugin-controller/pkg/types"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type StatefulSetRpcPlugin struct {
	Clienset *kubernetes.Clientset
	LogCtx   *logrus.Entry
	IsTest   bool
}

func (r *StatefulSetRpcPlugin) InitPlugin() pluginTypes.RpcError {
	if r.IsTest {
		return pluginTypes.RpcError{}
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		return pluginTypes.RpcError{ErrorString: err.Error()}
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return pluginTypes.RpcError{ErrorString: err.Error()}
	}
	r.Clienset = clientset

	return pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) SetWeight(rolloutplugin *v1alpha1.RolloutPlugin) pluginTypes.RpcError {
	fmt.Println("SetWeight")
	// need to know what to do here
	return pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) SetCanaryScale(rolloutplugin *v1alpha1.RolloutPlugin) pluginTypes.RpcError {
	return pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) Type() string {
	return "rolloutplugin"
}

func (r *StatefulSetRpcPlugin) UpdateHash(rolloutplugin *v1alpha1.RolloutPlugin) pluginTypes.RpcError {
	return pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) SetHeaderRoute(rolloutplugin *v1alpha1.RolloutPlugin) pluginTypes.RpcError {
	return pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) Rollback(rolloutplugin *v1alpha1.RolloutPlugin) pluginTypes.RpcError {
	return pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) Terminate(rolloutplugin *v1alpha1.RolloutPlugin, roCtx pluginTypes.RpcRolloutContext) (pluginTypes.RpcRolloutResult, pluginTypes.RpcError) {
	return pluginTypes.RpcRolloutResult{}, pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) Abort(rolloutplugin *v1alpha1.RolloutPlugin, roCtx pluginTypes.RpcRolloutContext) (pluginTypes.RpcRolloutResult, pluginTypes.RpcError) {
	return pluginTypes.RpcRolloutResult{}, pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) Run(rolloutplugin *v1alpha1.RolloutPlugin, roCtx pluginTypes.RpcRolloutContext) (pluginTypes.RpcRolloutResult, pluginTypes.RpcError) {
	return pluginTypes.RpcRolloutResult{}, pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) SetMirrorRoute(rolloutplugin *v1alpha1.RolloutPlugin) pluginTypes.RpcError {
	return pluginTypes.RpcError{}
}
