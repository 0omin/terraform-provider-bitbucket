package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/provider"
)

func TestProviderMetadataAndSchema(t *testing.T) {
	ctx := context.Background()
	p := New("test")()

	var meta provider.MetadataResponse
	p.Metadata(ctx, provider.MetadataRequest{}, &meta)
	if meta.TypeName != "bitbucket" {
		t.Errorf("TypeName = %q, want %q", meta.TypeName, "bitbucket")
	}
	if meta.Version != "test" {
		t.Errorf("Version = %q, want %q", meta.Version, "test")
	}

	var resp provider.SchemaResponse
	p.Schema(ctx, provider.SchemaRequest{}, &resp)
	if diags := resp.Schema.ValidateImplementation(ctx); diags.HasError() {
		t.Fatalf("schema is invalid: %v", diags)
	}
}
