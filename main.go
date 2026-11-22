package main

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"

	"github.com/0omin/terraform-provider-bitbucket/internal/provider"
)

// version is set at build time with -ldflags "-X main.version=...".
var version = "dev"

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/0omin/bitbucket",
	}

	if err := providerserver.Serve(context.Background(), provider.New(version), opts); err != nil {
		log.Fatal(err)
	}
}
