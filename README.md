# How to use

## Output executable

```bash
protoc -Igreet/proto --go_out=. --go_opt=module=github.com/jaycchiu524/go-api-grpc --go-grpc_out=. --go-grpc_opt=module=github.com/jaycchiu524/go-api-grpc greet/proto/greet.proto

go build -o bin/greet/server ./greet/server
go build -o bin/greet/client ./greet/client

# OR 
make greet # See Makefile

./bin/greet/server
./bin/greet/client

# output: Greet: Hello Jay
```
