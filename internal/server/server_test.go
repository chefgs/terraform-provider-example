package server

import (
	"context"
	"testing"
)

func TestNewServer(t *testing.T) {
	s := NewServer()
	
	if s == nil {
		t.Fatal("expected server to be initialized")
	}
	
	if s.Port != 80 {
		t.Errorf("expected default port to be 80, got %d", s.Port)
	}
}

func TestServer_Create(t *testing.T) {
	s := NewServer()
	ctx := context.Background()
	
	err := s.Create(ctx, "test-server", "192.168.1.1", 8080)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	
	if s.Name != "test-server" {
		t.Errorf("expected name to be 'test-server', got %s", s.Name)
	}
	
	if s.IPAddress != "192.168.1.1" {
		t.Errorf("expected IP address to be '192.168.1.1', got %s", s.IPAddress)
	}
	
	if s.Port != 8080 {
		t.Errorf("expected port to be 8080, got %d", s.Port)
	}
	
	if s.ID != "server-test-server" {
		t.Errorf("expected ID to be 'server-test-server', got %s", s.ID)
	}
}

func TestServer_Update(t *testing.T) {
	s := NewServer()
	ctx := context.Background()
	
	// First create the server
	err := s.Create(ctx, "test-server", "192.168.1.1", 8080)
	if err != nil {
		t.Fatalf("unexpected error during create: %v", err)
	}
	
	// Then update it
	err = s.Update(ctx, s.ID, "updated-server", "192.168.1.2", 9090)
	if err != nil {
		t.Fatalf("unexpected error during update: %v", err)
	}
	
	if s.Name != "updated-server" {
		t.Errorf("expected name to be 'updated-server', got %s", s.Name)
	}
	
	if s.IPAddress != "192.168.1.2" {
		t.Errorf("expected IP address to be '192.168.1.2', got %s", s.IPAddress)
	}
	
	if s.Port != 9090 {
		t.Errorf("expected port to be 9090, got %d", s.Port)
	}
}

func TestServer_Delete(t *testing.T) {
	s := NewServer()
	ctx := context.Background()
	
	// First create the server
	err := s.Create(ctx, "test-server", "192.168.1.1", 8080)
	if err != nil {
		t.Fatalf("unexpected error during create: %v", err)
	}
	
	// Then delete it
	err = s.Delete(ctx, s.ID)
	if err != nil {
		t.Fatalf("unexpected error during delete: %v", err)
	}
	
	// In a real implementation, we might check that the server no longer exists
	// But for this example, the Delete method doesn't actually change anything
}
