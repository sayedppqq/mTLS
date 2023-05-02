#### Generating our rootCA file.
- `openssl req -new -x509 -days 365 -key ca.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=Acme Root CA" -out ca.crt
`
#### Generating server certificate.
- `openssl genrsa -out ca.key 2048`
- `openssl req -new -key server.key -days 365 -out server.csr -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=*.example.com"`
- `openssl x509 -req -extfile <(printf "subjectAltName=DNS:example.com,DNS:www.example.com") -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt`

#### Generating the client certificate.
- `openssl genrsa -out client.key 2048`
- `openssl req -new -key client.key -days 365 -out client.csr -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=*.example.com"`
- `openssl x509 -req -extfile <(printf "subjectAltName=DNS:example.com,DNS:www.example.com") -in client.csr -CA ca.crt -CAkey ca.key -out client.crt -days 365 -sha256 -CAcreateserial`

### Resource
[How to MTLS in golang](https://kofo.dev/how-to-mtls-in-golang)

[mTLS Golang Example](https://github.com/haoel/mTLS)