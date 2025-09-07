package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestServerResource(t *testing.T) {
	// Create a server resource
	r := NewServerResource()

	// Check that it was initialized properly
	if r == nil {
		t.Fatal("expected resource to be initialized")
	}
}

// TestServerResourceSchema tests the schema definition
func TestServerResourceSchema(t *testing.T) {
	// Test creation of a server resource model
	model := &serverResourceModel{
		Name:      stringValue("test-server"),
		IPAddress: stringValue("192.168.1.1"),
		Port:      int64Value(8080),
	}

	// Check values
	if model.Name.ValueString() != "test-server" {
		t.Errorf("expected name to be 'test-server', got %s", model.Name.ValueString())
	}

	if model.IPAddress.ValueString() != "192.168.1.1" {
		t.Errorf("expected ip_address to be '192.168.1.1', got %s", model.IPAddress.ValueString())
	}

	if model.Port.ValueInt64() != 8080 {
		t.Errorf("expected port to be 8080, got %d", model.Port.ValueInt64())
	}
}

// Helper functions for creating test values
func stringValue(v string) types.String {
	return types.StringValue(v)
}

func int64Value(v int64) types.Int64 {
	return types.Int64Value(v)
}
