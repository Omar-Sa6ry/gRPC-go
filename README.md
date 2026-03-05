# gRPC Go Microservices Example

This project demonstrates various gRPC concepts and implementations using Go. It contains three main microservices: **Greet**, **Calculator**, and **Blog**. Each service explores different aspects of gRPC, from basic unary calls to complex bidirectional streaming, deadlines, TLS/SSL security, and database integration using MongoDB.

## Table of Contents
1. [Prerequisites](#prerequisites)
2. [Project Structure](#project-structure)
3. [Services Overview](#services-overview)
   - [Greet Service](#greet-service)
   - [Calculator Service](#calculator-service)
   - [Blog Service](#blog-service)
4. [Security (SSL/TLS)](#security-ssltls)
5. [How to Compile and Run](#how-to-compile-and-run)
6. [Testing with Evans REPL](#testing-with-evans-repl)

---

## Prerequisites

To compile and run this project, ensure you have the following installed on your machine:
- **Go** (Programming language runtime)
- **Protocol Buffers Compiler (protoc)** (Used to generate Go code from `.proto` files)
- **protoc-gen-go** and **protoc-gen-go-grpc** (Go plugins for protoc)
- **Make** (To run the `Makefile` commands for compilation)
- **Docker & Docker Compose** (Required for the MongoDB instance used in the Blog service)
- **OpenSSL** (Required for generating SSL certificates)
- **Evans** (Optional, for gRPC reflection and manual testing via CLI)

---

## Project Structure

The project is structured into modular directories representing individual services and utilities:

```text
.
├── blog/
│   ├── client/       # gRPC Client implementations for the Blog service
│   ├── proto/        # Protocol Buffers definitions for the Blog service
│   ├── server/       # gRPC Server and MongoDB implementations for the Blog service
│   └── docker-compose.yml # Docker configuration for MongoDB and Mongo Express
├── calculator/
│   ├── client/       # gRPC Client implementations for the Calculator service
│   ├── proto/        # Protocol Buffers definitions for the Calculator service
│   └── server/       # gRPC Server implementations for the Calculator service
├── greet/
│   ├── client/       # gRPC Client implementations for the Greet service
│   ├── proto/        # Protocol Buffers definitions for the Greet service
│   └── server/       # gRPC Server implementations for the Greet service
├── ssl/
│   ├── ssl.ps1       # PowerShell script to generate self-signed certificates
│   └── ssl.ps1.conf  # OpenSSL configuration file for the script
├── bin/              # Compiled server and client execution binaries (ignored in version control)
├── Makefile          # Contains build instructions for compiling protos and Go binaries
├── go.mod            # Go module definition and dependencies list
└── go.sum            # Checksums for Go dependencies
```

---

## Services Overview

### 1. Greet Service
The Greet service explores all four types of gRPC communication paradigms, along with Context Deadlines.

* **Files and Capabilities:**
  * `greet.go`: Unary RPC. The client sends a single request (a First Name), and the server responds with a simple greeting.
  * `greet_many_times.go`: Server Streaming RPC. The client sends one request, and the server responds with a stream of multiple messages.
  * `long_greet.go`: Client Streaming RPC. The client sends a stream of multiple requests, and the server responds once after receiving all messages.
  * `greet_everyone.go`: Bi-directional Streaming RPC. Both client and server send a stream of messages to each other concurrently.
  * `greet_with_deadline.go`: Demonstrates `context.WithTimeout`. The client requests a response within a certain time frame. If the server takes too long, the client cancels the request with a `DEADLINE_EXCEEDED` error.

### 2. Calculator Service
The Calculator service serves as a practical, math-oriented application of gRPC.

* **Files and Capabilities:**
  * `sum.go`: Unary RPC. Accepts two integers and returns their sum.
  * `primes.go`: Server Streaming RPC. Accepts an integer and streams its prime number factors back to the client.
  * `avg.go`: Client Streaming RPC. Accepts a stream of integers from the client and returns the calculated average once the stream ends.
  * `max.go`: Bi-directional Streaming RPC. Accepts a stream of numbers and responds with the maximum number seen so far whenever a new number is sent.
  * `sqrt.go`: Unary RPC with Error Handling. Accepts a number and returns its square root. If the number is negative, it returns an `INVALID_ARGUMENT` gRPC error code.

### 3. Blog Service
The Blog service demonstrates a complete CRUD (Create, Read, Update, Delete) API integrating gRPC with a MongoDB database.

* **Database (MongoDB):**
  * Spun up using `blog/docker-compose.yml`.
  * Includes `mongo-express` accessible at `http://localhost:8081` (default credentials: `admin` / `pass`) for viewing records via a web UI.
* **Files and Capabilities:**
  * `blog_item.go`: Defines the MongoDB `BlogItem` BSON struct and mapping functions to convert it to a gRPC `Blog` struct.
  * `create.go`: Unary RPC. Inserts a new blog post into MongoDB and returns the assigned Object ID.
  * `read.go`: Unary RPC. Fetches a single blog post from MongoDB by its Object ID securely. Returns a `NOT_FOUND` error if the blog does not exist.
  * `update.go`: Unary RPC. Modifies an existing blog post entirely based on the provided Object ID.
  * `delete.go`: Unary RPC. Removes a blog post from the MongoDB database permanently.
  * `list.go`: Server Streaming RPC. Queries the MongoDB collection for all blog posts and streams them one by one back to the client.

---

## Security (SSL/TLS)

The project includes an infrastructure to generate dummy SSL/TLS certificates for testing secure gRPC connections.

* **Generating Certificates:**
  Navigate to the `ssl` directory and run the PowerShell script:
  ```powershell
  cd ssl
  .\ssl.ps1
  ```
  This will use OpenSSL to generate a Root Certificate Authority (`ca.crt`, `ca.key`) and sign a server certificate (`server.crt`, `server.pem`, `server.key`).

* **Applying SSL to Services:**
  The servers and clients (e.g., in `main.go` of the Greet service) have code blocks that can be uncommented or switched via a `tls := true` variable to utilize `credentials.NewServerTLSFromFile` and `credentials.NewClientTLSFromFile` for fully encrypted HTTP/2 communication.

---

## How to Compile and Run

### Step 1: Install Dependencies
Run the following command at the root of the project to download all required packages (gRPC, Protocol Buffers, Mongo Driver):
```powershell
go mod tidy
```

### Step 2: Protocol Buffer Compilation & Go Build
The project uses a `Makefile` to compile the `.proto` files into Go code and then build the `server.exe` and `client.exe` binaries into the `bin/` directory.

To build a specific service, run:
```powershell
make greet
```
or
```powershell
make calculator
```
or 
```powershell
make blog
```

### Step 3: Start the Databases (If using Blog Service)
Navigate to the `blog` folder and start the Docker containers:
```powershell
cd blog
docker-compose up -d
cd ..
```

### Step 4: Run the Server and Client
First, open a terminal and run the server explicitly from the generated binaries folder:
```powershell
./bin/blog/server
```
*(Replace `blog` with `greet` or `calculator` depending on the service you want to test).*

Next, open a new terminal window and run the client to invoke the RPC calls configured in the client's `main.go`:
```powershell
./bin/blog/client
```

---

## Testing with Evans REPL

The Calculator service server has `reflection.Register(grpcServer)` enabled. This instructs the gRPC server to expose its APIs, allowing tools like **Evans** to interact with it seamlessly without needing the underlying `.proto` files on the client machine.

1. Ensure the Calculator server is running:
   ```powershell
   ./bin/calculator/server
   ```
2. Start the Evans REPL interface:
   ```powershell
   evans --host localhost --port 50051 -r repl
   ```
3. Inside Evans, you can view the available services and invoke RPCs manually:
   ```text
   # Show available services
   > show service
   
   # Select the Calculator service
   > service CalculatorService
   
   # Execute an RPC, for example, the Sum function
   > call Sum
   ```
   Evans will then prompt you interactively to input the required fields for the request.
