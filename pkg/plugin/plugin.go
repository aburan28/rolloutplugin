package plugin

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/aburan28/rolloutplugin-controller/api/v1alpha1"
	"github.com/aburan28/rolloutplugin-controller/pkg/plugin/rpc"
	pluginTypes "github.com/aburan28/rolloutplugin-controller/pkg/types"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type StatefulSetRpcPlugin struct {
	Clienset *kubernetes.Clientset
	LogCtx   *logrus.Entry
	IsTest   bool

	dynamicInformerFactory        dynamicinformer.DynamicSharedInformerFactory
	clusterDynamicInformerFactory dynamicinformer.DynamicSharedInformerFactory
	istioDynamicInformerFactory   dynamicinformer.DynamicSharedInformerFactory
}

var _ rpc.RolloutPlugin = (*StatefulSetRpcPlugin)(nil)

func (r *StatefulSetRpcPlugin) InitPlugin() pluginTypes.RpcError {
	// setup informer??
	r.LogCtx.Info("InitPlugin")
	if r.IsTest {
		return pluginTypes.RpcError{}
	}

	// Use the default kubeconfig file from the user's home directory.
	var kubeconfig string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = filepath.Join(home, ".kube", "config")
	}

	// Build the config from the kubeconfig file.
	// This automatically uses the current/default context set in your kubeconfig.
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		return pluginTypes.RpcError{ErrorString: fmt.Sprintf("failed to build kubeconfig: %v", err)}
	}

	// Create the Kubernetes clientset.
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return pluginTypes.RpcError{ErrorString: fmt.Sprintf("failed to create clientset: %v", err)}
	}
	r.Clienset = clientset

	return pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) CheckForRollouts(clientset *kubernetes.Clientset) pluginTypes.RpcError {
	return pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) SetWeight(rolloutplugin *v1alpha1.RolloutPlugin) pluginTypes.RpcError {

	r.LogCtx.Info("SetWeight")
	r.LogCtx.Info(rolloutplugin.Name)

	ctx := context.TODO()

	ss, err := r.lookupStatefulSet(ctx, rolloutplugin.Spec.Selector.MatchLabels, rolloutplugin.Name, rolloutplugin.Namespace)
	if err != nil {
		r.LogCtx.Errorf("Error looking up StatefulSet: %v", err)
		return pluginTypes.RpcError{ErrorString: fmt.Sprintf("Error looking up StatefulSet: %v", err)}
	}

	curWeight := rolloutplugin.Status.CurrentWeight
	// do the math on updating the weight/replicas
	desiredReplicas := ss.Spec.Replicas
	currentReplicas := ss.Status.Replicas
	partition := ss.Spec.UpdateStrategy.RollingUpdate.Partition

	// Update the StatefulSet

	if ss == nil {
		return pluginTypes.RpcError{ErrorString: "StatefulSet not found"}
	}
	// need to know what to do here
	return pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) SetCanaryScale(rolloutplugin *v1alpha1.RolloutPlugin) pluginTypes.RpcError {
	r.LogCtx.Info("SetCanaryScale")
	return pluginTypes.RpcError{}
}

func (r *StatefulSetRpcPlugin) Type() string {
	return "RpcRolloutPlugin"
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
