## To build the program

```
go build
```

## To run

```
./go-api
```

On another terminal we can curl the API using https

```
curl https://localhost:8080/hello --cacert ./cert/CA/CA.pem

`Hello All`
```
## Steps to change scheme from `http` to `https` using self signed certificate

### Generate a Private Key and Root CA certificate for Certificate Authority (CA).

```
 mkdir cert
 cd cert
 mkdir CA
 cd CA
 openssl genrsa -out CA.key -des3 2048 # private key generated
 # generate Root CA certificate
 openssl req -x509 -sha256 -new -nodes -days 3650 -key CA.key -out CA.pem 
```

 ### Generate a Private key and metadata information for localhost server.

``` 
 mkdir localhost
 cd localhost
 touch localhost.ext
 # Add the metadata information as in the repo
 # generate Private Key for the server
 openssl genrsa -out localhost.key -des3 2048
```

 ### Generate CSR for the server for the go-api server running on localhost.
```
 openssl req -new -key localhost.key -out localhost.csr # use localhost in the Common Name (CN) or any address but that must be present in /etc/hosts
```
 ### Request to CA using the CSR
```
 openssl x509 -req -in localhost.csr -CA ../CA.pem -CAkey ../CA.key -CAcreateserial -days 3650 -sha256 -extfile localhost.ext -out localhost.crt
```

 ### Decrypt the localhost key as server needs it.

```
openssl rsa -in localhost.key -out localhost.decrypted.key
```

Now we can provide the localhost certificate and decryption key to the server.