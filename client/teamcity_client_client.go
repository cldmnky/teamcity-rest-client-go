// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/cldmnky/teamcity-rest-client-go/client/agent"
	"github.com/cldmnky/teamcity-rest-client-go/client/agent_pool"
	"github.com/cldmnky/teamcity-rest-client-go/client/build"
	"github.com/cldmnky/teamcity-rest-client-go/client/build_queue"
	"github.com/cldmnky/teamcity-rest-client-go/client/build_type"
	"github.com/cldmnky/teamcity-rest-client-go/client/change"
	"github.com/cldmnky/teamcity-rest-client-go/client/cloud_instance"
	"github.com/cldmnky/teamcity-rest-client-go/client/debug"
	"github.com/cldmnky/teamcity-rest-client-go/client/federation"
	"github.com/cldmnky/teamcity-rest-client-go/client/group"
	"github.com/cldmnky/teamcity-rest-client-go/client/investigation"
	"github.com/cldmnky/teamcity-rest-client-go/client/mute"
	"github.com/cldmnky/teamcity-rest-client-go/client/operations"
	"github.com/cldmnky/teamcity-rest-client-go/client/problem"
	"github.com/cldmnky/teamcity-rest-client-go/client/project"
	"github.com/cldmnky/teamcity-rest-client-go/client/server"
	"github.com/cldmnky/teamcity-rest-client-go/client/test"
	"github.com/cldmnky/teamcity-rest-client-go/client/test_occurrence"
	"github.com/cldmnky/teamcity-rest-client-go/client/user"
	"github.com/cldmnky/teamcity-rest-client-go/client/vcs_root"
	"github.com/cldmnky/teamcity-rest-client-go/client/vcs_root_instance"
)

// Default teamcity client HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "teamcit:8111"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new teamcity client HTTP client.
func NewHTTPClient(formats strfmt.Registry) *TeamcityClient {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new teamcity client HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *TeamcityClient {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new teamcity client client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *TeamcityClient {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(TeamcityClient)
	cli.Transport = transport
	cli.Agent = agent.New(transport, formats)
	cli.AgentPool = agent_pool.New(transport, formats)
	cli.Build = build.New(transport, formats)
	cli.BuildQueue = build_queue.New(transport, formats)
	cli.BuildType = build_type.New(transport, formats)
	cli.Change = change.New(transport, formats)
	cli.CloudInstance = cloud_instance.New(transport, formats)
	cli.Debug = debug.New(transport, formats)
	cli.Federation = federation.New(transport, formats)
	cli.Group = group.New(transport, formats)
	cli.Investigation = investigation.New(transport, formats)
	cli.Mute = mute.New(transport, formats)
	cli.Operations = operations.New(transport, formats)
	cli.Problem = problem.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.Server = server.New(transport, formats)
	cli.Test = test.New(transport, formats)
	cli.TestOccurrence = test_occurrence.New(transport, formats)
	cli.User = user.New(transport, formats)
	cli.VcsRoot = vcs_root.New(transport, formats)
	cli.VcsRootInstance = vcs_root_instance.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// TeamcityClient is a client for teamcity client
type TeamcityClient struct {
	Agent agent.ClientService

	AgentPool agent_pool.ClientService

	Build build.ClientService

	BuildQueue build_queue.ClientService

	BuildType build_type.ClientService

	Change change.ClientService

	CloudInstance cloud_instance.ClientService

	Debug debug.ClientService

	Federation federation.ClientService

	Group group.ClientService

	Investigation investigation.ClientService

	Mute mute.ClientService

	Operations operations.ClientService

	Problem problem.ClientService

	Project project.ClientService

	Server server.ClientService

	Test test.ClientService

	TestOccurrence test_occurrence.ClientService

	User user.ClientService

	VcsRoot vcs_root.ClientService

	VcsRootInstance vcs_root_instance.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *TeamcityClient) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Agent.SetTransport(transport)
	c.AgentPool.SetTransport(transport)
	c.Build.SetTransport(transport)
	c.BuildQueue.SetTransport(transport)
	c.BuildType.SetTransport(transport)
	c.Change.SetTransport(transport)
	c.CloudInstance.SetTransport(transport)
	c.Debug.SetTransport(transport)
	c.Federation.SetTransport(transport)
	c.Group.SetTransport(transport)
	c.Investigation.SetTransport(transport)
	c.Mute.SetTransport(transport)
	c.Operations.SetTransport(transport)
	c.Problem.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.Server.SetTransport(transport)
	c.Test.SetTransport(transport)
	c.TestOccurrence.SetTransport(transport)
	c.User.SetTransport(transport)
	c.VcsRoot.SetTransport(transport)
	c.VcsRootInstance.SetTransport(transport)
}
