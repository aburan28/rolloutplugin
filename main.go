package main

import (
	"flag"
	"fmt"
	"os"

	"rolloutplugin/pkg/plugin"

	rpc "github.com/aburan28/rolloutplugin-controller/pkg/plugin/rpc"
	goPlugin "github.com/hashicorp/go-plugin"
	"github.com/hashicorp/serf/version"
	log "github.com/sirupsen/logrus"
)

// handshakeConfigs are used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var handshakeConfig = goPlugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "ARGO_ROLLOUTS_RPC_PLUGIN",
	MagicCookieValue: "rolloutplugin",
}

func main() {
	// Create a flag to print the version of the plugin
	// This is useful for debugging and support
	versionFlag := flag.Bool("version", false, "Print the version of the plugin")
	flag.Parse()
	if *versionFlag {
		fmt.Fprintln(os.Stderr, version.GetHumanVersion()) // print to stderr
		return
	}
	// log.SetOutput(os.Stderr)
	log.SetOutput(os.Stderr)

	logCtx := log.WithFields(log.Fields{"plugin": "rolloutpl2222ugin"})

	rpcPluginImp := &plugin.StatefulSetRpcPlugin{
		LogCtx: logCtx,
	}

	var pluginMap = map[string]goPlugin.Plugin{
		"RpcRolloutPlugin": &rpc.RpcRolloutPlugin{
			Impl: rpcPluginImp,
		},

		// "RpcTrafficRouterPlugin": &rolloutsPlugin.RpcTrafficRouterPlugin{Impl: rpcPluginImp},
	}

	goPlugin.Serve(&goPlugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
