# Conformance Testing

This document outlines the conformance testing setup for this repository. The goal of these tests is to ensure that all server implementations (`standardlib`, `gorilla`, and `gin`) adhere to the same API specification.

## Components

The conformance testing setup consists of several key components:

### 1. OpenAPI Specification (`openapi.yaml`)

The `openapi.yaml` file, located in the root of the repository, serves as the single source of truth for the API definition. It describes all the available endpoints, their expected request and response formats, and the data models used.

Any changes to the API should be made to this file first.

### 2. Generated Go Client (`pkg/client`)

A Go client is automatically generated from the `openapi.yaml` specification. This client provides a typed interface for interacting with the API, which is used by the conformance tests to make requests to the servers.

The client is generated using the `oapi-codegen` tool.

### 3. Taskfile (`Taskfile.yaml`)

The `Taskfile.yaml` file provides a simple way to manage common development tasks. It includes a `generate-client` task, which runs `oapi-codegen` to regenerate the Go client whenever the `openapi.yaml` file changes.

To regenerate the client, run:

```bash
go tool task generate-client
```

### 4. Tool Dependencies

The specific versions of the tools used for development, such as `task` and `oapi-codegen`, are managed in the `go.mod` file using the `tool` directive.

### 5. Conformance Tests (`test/conformance_test.go`)

The conformance tests themselves are located in `test/conformance_test.go`. These tests are designed to be implementation-agnostic.

For each server implementation, the test suite will:
1.  Start the server as a background process.
2.  Use the generated Go client to perform a series of tests against the running server, covering all API endpoints (CRUD operations).
3.  Shut down the server process.

## Running the Tests

To run the full suite of conformance tests against all implementations, use the following command:

```bash
go test -v ./test
```
Or use the `test` task in the `Taskfile.yaml`:
```bash
go tool task test
```
