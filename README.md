## Multi gRPC Application

This multi gRPC application brings together a Go gRPC server, a Node.js gRPC client (`nodejs_grpc_client`), and a Python gRPC client (`py_grpc_client`). The components are orchestrated using Docker Compose for streamlined deployment and testing in a containerized environment.

## Overview

### Go gRPC Server (`go_grpc_server`)
The central component, hosted on port 8080, implements the `UserServiceServer` interface, managing various user-related functionalities.

### Node.js gRPC Client (`nodejs_grpc_client`)
Communicates bidirectionally with the Go server using `@grpc/proto-loader`. This client issues requests to retrieve specific user data and the entire user list from the Go server.

### Python gRPC Client (`py_grpc_client`)
Connects to the Go server using the `grpc` module and generated gRPC classes (`user_pb2` and `user_pb2_grpc`). This client fetches user details from the server.

## Project Structure
```plaintext
/multi_grpc_env
|-- docker-compose.yml
|-- go_grpc_client
|-- go_grpc_server
|-- nodejs_grpc_client
|-- py_grpc_client
```

## Getting Started
1. Clone the repository.
2. Run the Docker Compose configuration using `docker-compose up`.
3. Explore the interaction between the Go server, Node.js client, and Python client in this multi gRPC environment.

Feel free to contribute, report issues, or provide feedback. Let's enhance and optimize this multi gRPC application together!
