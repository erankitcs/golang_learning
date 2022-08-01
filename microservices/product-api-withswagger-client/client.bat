@ECHO off    
WHERE /q swagger
IF %ERRORLEVEL% NEQ 0 go get -u github.com/go-swagger/go-swagger/cmd/swagger 
go mod init github.com/erankitcs/golang_learning/microservices/productapiclient
swagger generate client -f ./../product-api-withswagger/swagger.yaml -A productapi
go mod tidy