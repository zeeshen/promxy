package http

import (
	"git.llsapp.com/hunter/oc-wrapper/defines"
	"git.llsapp.com/hunter/oc-wrapper/env"
	"go.opencensus.io/trace"
)

// set default attributes according to our spec
// https://git.llsapp.com/hunter/hunter-spec/blob/master/spec/trace.md
func clientDefaultAttributes(addr string) []trace.Attribute {
	return []trace.Attribute{
		trace.StringAttribute(defines.SERVICE_NAME, env.GetServiceName()),
		trace.StringAttribute(defines.SERVICE_REMOTE_KIND, defines.SERVICE_KIND_HTTP),
		trace.StringAttribute(defines.SERVICE_HOST_NAME, env.GetHostName()),
		trace.StringAttribute(defines.SERVICE_REMOTE_ADDR, addr),
	}
}

func serverDefaultAttributes() []trace.Attribute {
	return []trace.Attribute{
		trace.StringAttribute(defines.SERVICE_NAME, env.GetServiceName()),
		trace.StringAttribute(defines.SERVICE_KIND, defines.SERVICE_KIND_HTTP),
		trace.StringAttribute(defines.SERVICE_HOST_NAME, env.GetHostName()),
	}
}
