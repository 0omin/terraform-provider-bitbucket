// Package provider implements the Bitbucket Cloud Terraform provider.
package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = (*bitbucketProvider)(nil)

type bitbucketProvider struct {
	version string
}

// New returns a factory for the provider. version is reported to Terraform.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &bitbucketProvider{version: version}
	}
}

func (p *bitbucketProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "bitbucket"
	resp.Version = p.version
}

func (p *bitbucketProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Manage Bitbucket Cloud resources through the Bitbucket Cloud REST API 2.0.",
	}
}

func (p *bitbucketProvider) Configure(context.Context, provider.ConfigureRequest, *provider.ConfigureResponse) {
}

func (p *bitbucketProvider) Resources(context.Context) []func() resource.Resource {
	return nil
}

func (p *bitbucketProvider) DataSources(context.Context) []func() datasource.DataSource {
	return nil
}
