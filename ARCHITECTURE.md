# Ignite

## MVP Features

- Expose the container to the internet
- Connecting to a domain
- Very basic metrics (req/responses on ur function)
- Scaling
- Stack

## Reverse Proxy
Envoy Proxy will be used for exposing the container and connecting to a domain. Envoy is an open-source edge and service proxy that provides load balancing, routing, and observability for distributed systems. It can be used to route traffic to the container hosting the HTTP function.

## Prometheus

used for function metrics and stuff :3 think they have a Go SDK

## Podman

Podman is a tool for managing containers and is used for container creation where the function will be deployed. Podman is an alternative to Docker and provides a similar containerization experience, without requiring a daemon to run in the background.
