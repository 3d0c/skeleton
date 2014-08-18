## Martini based MVC skeleton

This is an application skeleton for typical API server, based on martini framework. You can quickly bootstrap your environment, using some MVC principles.  
This is a renewed version, so if you're looking for full featured example, that was here before, please checkout `v1.0` branch. 

### Usage

Get it:  

```sh
go get github.com/3d0c/skeleton
```

Generate ssl keys:  

```sh
go run $GOROOT/src/pkg/crypto/tls/generate_cert.go --host localhost --ca
```
it will generate `cert.pem` and `key.pem` files.

Check out `config.json` for base settings. Default one:

```json
{
    "application" : {
        "listen_on" : ":5500",
        "https_on"  : ":5443",
        "ssl_cert"  : "./cert.pem",
        "ssl_key"   : "./key.pem"
    }
}
```

Run it by simple `go run app.go` or using some code reloader `gin -p 5000 -a 5500`, so it will be available on `5500` or `5000` port.  
Check it:

```sh
~ curl -i http://localhost:5500/user

HTTP/1.1 200 OK
Access-Control-Allow-Headers: Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token
Access-Control-Allow-Methods: POST, GET, OPTIONS, PUT, DELETE
Access-Control-Allow-Origin: *
Cache-Control: max-age=2592000
Cache-Control: public
Content-Type: application/json
Pragma: public
Date: Mon, 18 Aug 2014 21:37:25 GMT
Content-Length: 19

{"name":"test one"}
```
