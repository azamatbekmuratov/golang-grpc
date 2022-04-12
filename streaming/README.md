## gRPC Server-Side Streaming with Go

This showcase demonstrates abilities of gRPC Server-Side Streaming using Golang language

### There are 3 types of streaming:

**Client-side streaming:** Where the client will have multiple requests and the server will only return one response.

**Bidirectional streaming:** Where both client and server can have multiple requests and responses together within a single connection.

**Server-side streaming:** Where the client sends a single request and the server can return several responses together. This is the one I will be showing you how to implement today.