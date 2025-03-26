package plugin

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/aburan28/rolloutplugin-controller/api/v1alpha1"
	"github.com/aburan28/rolloutplugin-controller/pkg/plugin/rpc"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	goPlugin "github.com/hashicorp/go-plugin"
)

func TestRollout(t *testing.T) {
	// Path to the plugin binary. Ensure you have built the goPlugin.
	pluginPath := filepath.Join(".", "statefulset")
	if _, err := os.Stat(pluginPath); os.IsNotExist(err) {
		t.Fatalf("plugin binary not found at %s. Build it with:\n  go build -o rollout_plugin ./rollout_plugin", pluginPath)
	}

	handshake := goPlugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "ARGO_ROLLOUTS_RPC_PLUGIN",
		MagicCookieValue: "rolloutplugin",
	}

	// On the client side, we use an empty PluginWrapper.
	pluginMap := map[string]goPlugin.Plugin{
		"RpcRolloutPlugin": &rpc.RpcRolloutPlugin{},
	}

	client := goPlugin.NewClient(&goPlugin.ClientConfig{
		HandshakeConfig:  handshake,
		Plugins:          pluginMap,
		Cmd:              exec.Command("./" + pluginPath),
		AllowedProtocols: []goPlugin.Protocol{goPlugin.ProtocolNetRPC},
	})
	defer client.Kill()

	// Connect via RPC.
	rpcClient, err := client.Client()
	if err != nil {
		t.Fatalf("error creating RPC client: %v", err)
	}

	// Dispense the "rollout" goPlugin.
	raw, err := rpcClient.Dispense("RpcRolloutPlugin")
	if err != nil {
		t.Fatalf("error dispensing rollout plugin: %v", err)
	}

	rollout, ok := raw.(rpc.RolloutPlugin)
	if !ok {
		t.Fatalf("unexpected type from plugin")
	}

	got := rollout.InitPlugin()
	if got.ErrorString != "" {
		t.Fatalf("error initializing plugin: %s", got.ErrorString)
	}
	v := &v1alpha1.RolloutPlugin{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "test",
			Namespace: "default",
		},
		Spec: v1alpha1.RolloutPluginSpec{},
	}
	rpcErr := rollout.SetWeight(v)
	require.EqualError(t, rpcErr, "statefulsets.apps \"test\" not found")

	rpcErr = rollout.SetCanaryScale(v)

}
