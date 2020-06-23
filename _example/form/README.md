This example illustrates how to use the `plex` server to serve `application/x-www-form-urlencoded` requests.

The client should provide the following headers in order the request to be served:

- `Content-Type: application/x-www-form-urlencoded` that determines the format of the content send to the server
- `Accept: application/json` that determines the format of the content that the client expect. If this header is missing the server will use the value of `Content-Type` header to serve the response. **We do not support `application/x-www-form-urlencoded` format in the response body**

### Example

The form post request:

```bash
$ curl -v -X POST -H 'Content-Type: application/x-www-form-urlencoded' -H 'Accept: application/json' http://localhost:8080/v2/bar -d 'name=Jack' 
```

The request can be made with a `JSON` body as well:

```bash
$ curl -v -X POST -H 'Content-Type: application/json' http://localhost:8080/v2/bar -d '{"name":"Jack"}' 
```

The response for both formats produces the same response:

```bash
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> POST /v2/bar HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Content-Type: application/x-www-form-urlencoded
> Accept: application/json
> Content-Length: 9
>
* upload completely sent off: 9 out of 9 bytes
< HTTP/1.1 200 OK
< Cache-Control: no-cache, no-store, no-transform, must-revalidate, private, max-age=0
< Content-Type: application/json
< Expires: Thu, 01 Jan 1970 01:00:00 BST
< Pragma: no-cache
< X-Accel-Expires: 0
< Date: Tue, 23 Jun 2020 12:17:05 GMT
< Content-Length: 25
<
* Connection #0 to host localhost left intact
{"body":"Welcome, Jack!"}* Closing connection 0
```
