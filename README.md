```bash
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/jaycchiu524/go-api-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/jaycchiu524/go-api-grpc greet/proto/dummy.proto
```