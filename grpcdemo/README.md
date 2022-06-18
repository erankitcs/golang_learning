### Vacation Tracker based on GRPC and GO Lang


### Tools to install
```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
go env GOPATH
$env:PATH +="{GO PATH HERE}\bin"
```
### Commands

1. Generating Server Code from protoc message
```
cd server
go mod init github.com/erankitcs/golang_learning/grpcdemo/server
protoc --proto_path=../pb ../pb/*.proto --go_out=. --go-grpc_out=.
```

2. Certificate Generate commands
```
cd cert
# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=AU/ST=Vic/L=Melbourne/O=Go Grpc/OU=Testing/CN=*.gogrpc.test/emailAddress=gogrpctest@gmail.com"
# View CA certificate
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem -subj "/C=AU/ST=Vic/L=Melbourne/O=Go Grpc/OU=ServerTesting/CN=*.gogrpcserver.com/emailAddress=gogrpcserver@gmail.com"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in server-req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem -extfile server-ext.cnf

echo "Server's signed certificate"
openssl x509 -in server-cert.pem -noout -text
 
 # 4. Move certs starts with server into cert folder inside server. 
```

3. Build and Run server
```
cd server
go build .
./server.exe
```

4. Call server via grpccurl

- You need to add below line for local DNS resolution.

127.0.0.1  test.gogrpcserver.com

- Need to set Environment variable 
```
$env:GODEBUG="x509ignoreCN=0"
```
- Run below command to call grpc server. 
```
grpcurl -cacert ca-cert.pem test.gogrpcserver.com:9000 messages.EmployeeService/GetAll
```

*Postman support grpc now*
https://blog.postman.com/postman-now-supports-grpc/