/*
 * Copyright (c) Matthew Hartstonge <matt@mykro.co.nz>
 * SPDX-License-Identifier: MPL-2.0
 */

package provider

import (
	"context"
	"net/url"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	faClient "github.com/matthewhartstonge/terraform-provider-fusionauth/internal/client"
)

const (
	envApiToken = "FUSIONAUTH_API_TOKEN"
	envEndpoint = "FUSIONAUTH_ENDPOINT"
	envTenant   = "FUSIONAUTH_TENANT"
)

// Ensure FusionAuthProvider satisfies various provider interfaces.
var _ provider.Provider = &FusionAuthProvider{}

// FusionAuthProvider defines the provider implementation.
type FusionAuthProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// FusionAuthProviderModel describes the provider data model.
type FusionAuthProviderModel struct {
	ApiToken types.String `tfsdk:"api_token"`
	Endpoint types.String `tfsdk:"endpoint"`
	Tenant   types.String `tfsdk:"tenant"`
}

func (p *FusionAuthProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "fusionauth"
	resp.Version = p.version
}

func (p *FusionAuthProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description:         "",
		MarkdownDescription: "",

		Attributes: map[string]schema.Attribute{
			"api_token": schema.StringAttribute{
				MarkdownDescription: "FusionAuth API Token. Can also be set with the `" + envApiToken + "` environment variable.",
				Optional:            true,
			},
			"endpoint": schema.StringAttribute{
				MarkdownDescription: "Address of the FusionAuth instance to use. Can also be set with the `" + envEndpoint + "` environment variable.",
				Optional:            true,
			},
			"tenant": schema.StringAttribute{
				MarkdownDescription: "FusionAuth tenant to scope requests to. Can also be set with the `" + envTenant + "` environment variable.",
				Optional:            true,
			},
		},
	}
}

func (p *FusionAuthProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Check environment variables
	apiToken := os.Getenv(envApiToken)
	endpoint := os.Getenv(envEndpoint)
	tenant := os.Getenv(envTenant)

	var data FusionAuthProviderModel

	// Read configuration data into model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Configuration values are now available.
	if data.ApiToken.String() != "" {
		apiToken = data.ApiToken.String()
	}
	if apiToken == "" {
		resp.Diagnostics.AddError(
			"Missing API Token Configuration",
			"While configuring the provider, the API token was not found in "+
				"the "+envApiToken+" environment variable or provider "+
				"configuration block api_token attribute.",
		)
	}

	if data.Endpoint.String() != "" {
		endpoint = data.Endpoint.String()
	}
	if endpoint == "" {
		resp.Diagnostics.AddError(
			"Missing Endpoint Configuration",
			"While configuring the provider, the API endpoint was not found in "+
				"the "+envEndpoint+" environment variable or provider "+
				"configuration block endpoint attribute.",
		)
	}

	if data.Tenant.String() != "" {
		tenant = data.Tenant.String()
	}

	baseURL, err := url.Parse(endpoint)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to parse FusionAuth API Endpoint",
			"While configuring the provider, the API endpoint '"+endpoint+"'was unable "+
				"to be parsed as a URL.",
		)
	}

	// Finalized validating config, return errors
	if resp.Diagnostics.HasError() {
		return
	}

	// client configuration for data sources and resources
	client := faClient.New(baseURL, apiToken, tenant)

	// Bind in the client data
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *FusionAuthProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTenantResource,
	}
}

func (p *FusionAuthProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &FusionAuthProvider{
			version: version,
		}
	}
}
