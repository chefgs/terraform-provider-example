terraform {
  required_providers {
    example = {
      source  = "chefgs/example"
      version = "1.0.0"
    }
  }
}

resource "example_server" "my_server" {
  name       = "my-server"
  ip_address = "192.168.1.100"
  port       = 8080
}

output "server_id" {
  value = example_server.my_server.id
}

output "server_name" {
  value = example_server.my_server.name
}

output "server_ip" {
  value = example_server.my_server.ip_address
}

output "server_port" {
  value = example_server.my_server.port
}
