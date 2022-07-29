### Product API Gorilla Toolkit

### Usage

1. Get Products API-
```
http://localhost:9090/
```
2. Add a Product
```
http://localhost:9090/
Body-
{"id":3,"name":"tea","description":"A Tea","price":111.21,"sku":"ssdff-sfff-sfff"}
```
3. Update a Product
```
http://localhost:9090/3
Body-
{"id":3,"name":"tea","description":"A Tea of new type","price":122.21,"sku":"ssdff-sfff-sfff"}
```

### Swagger Documentation
- With Docker on Windows.
```
docker pull quay.io/goswagger/swagger
./swagger.bat
```
- Without Docker on Windows.
```
./make.bat
```
- Without Docker on Linux/Mac.
```
./make.bat
```