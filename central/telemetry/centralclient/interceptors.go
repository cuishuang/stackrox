package centralclient

import (
	"strings"

	"github.com/stackrox/rox/pkg/clientconn"
	"github.com/stackrox/rox/pkg/telemetry/phonehome"
)

var (
	trackedPaths []string
	ignoredPaths = []string{"/v1/ping", "/v1.PingService/Ping", "/v1/metadata", "/static/*"}

	trackedUserAgents []string

	interceptors = map[string][]phonehome.Interceptor{
		"API Call": {apiCall, addDefaultProps},
		"roxctl":   {roxctl, addDefaultProps},
	}
)

func addDefaultProps(rp *phonehome.RequestParams, props map[string]any) bool {
	props["Path"] = rp.Path
	props["Code"] = rp.Code
	props["Method"] = rp.Method
	props["User-Agent"] = rp.UserAgent
	if cmd := phonehome.GetFirst(rp.Headers, clientconn.RoxctlCommandHeader); cmd != "" {
		props["roxctl Command"] = cmd
	}
	if index := phonehome.GetFirst(rp.Headers, clientconn.RoxctlCommandIndexHeader); index != "" {
		props["roxctl Command Index"] = index
	}
	if execEnv := phonehome.GetFirst(rp.Headers, clientconn.ExecutionEnvironment); execEnv != "" {
		props["Execution Environment"] = execEnv
	}
	return true
}

// apiCall enables API Call events for the API paths specified in the
// trackedPaths ("*" value enables all paths) or for the calls with the
// User-Agent containing the substrings specified in the trackedUserAgents, and
// have no match in the ignoredPaths list.
func apiCall(rp *phonehome.RequestParams, _ map[string]any) bool {
	return !rp.HasPathIn(ignoredPaths) && (rp.HasPathIn(trackedPaths) ||
		rp.HasUserAgentWith(trackedUserAgents))
}

// roxctl enables the roxctl event.
func roxctl(rp *phonehome.RequestParams, _ map[string]any) bool {
	return strings.Contains(rp.UserAgent, "roxctl")
}
