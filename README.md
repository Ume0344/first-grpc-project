# first-grpc-project

### Proposal

- Client will send the request to to server to create a new user.
- Server will create a user and acknowledge back to client with a unique id.

### Implementation
- Define the service in .proto file.
- Generate protobuf files by running from root (first-grpc-project);
    ```
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative user_mgmt/usermgmt.proto
    ```
- Above command will generate two files;
    - usermgmt.pb.go - provides logic for serializing or deserializing the messages.
    - usermgmt_grpc.pb.go - includes generated code for server and client for rpc.

### Connecting two different physical/virtual servers
- Make sure you can ping servers from each other.
- Just change the address to `<ipv4 address of server:<server port>>` i.e, `192.168.0.4:5005` in client implementation.
