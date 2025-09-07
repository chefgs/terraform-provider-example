package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

// testAccProtoV6ProviderFactories are used to instantiate a provider during
// acceptance testing. The factory function will be invoked for every Terraform
// CLI command executed to create a provider server to which the CLI can
// reattach.
var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"example": providerserver.NewProtocol6WithError(New("test")()),
}

func TestProviderConfigure(t *testing.T) {
	// Create a provider instance
	p := New("test")()

	// Verify the provider was initialized properly
	if p == nil {
		t.Fatal("expected provider to be initialized")
	}
}
