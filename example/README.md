# Example API

This example demonstrates how to use the `plex` server.

The service supports HTTP requests with content-type:

- `application/json`
- `application/x-www-form-urlencoded`

In order to use the desired format you should pass set the `Content-Type` header
to one of the value above. Note that the response format is always
`application/json`.

## Getting Started

You can make a form-post request with the following command:

```bash
$ curl -v -X POST http://localhost:8080/v1/users \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Accept: application/json' \
  -d 'email=john.doe@example.com&password=swordfish'
```

You can make a json-post request the the following command:

```bash
$ curl -v -X POST http://localhost:8080/v1/users \
  -H 'Content-Type: application/json' \
  -d '{"email":"john.doe@example.com", "password": "swordfish}'
```

The response for both formats produces the same response:

```bash
* TCP_NODELAY set
* Connected to localhost (::1) port 8080 (#0)
> POST /v1/users HTTP/1.1
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
{"id":"457a1114-8832-4ad2-b950-33eee5fab920"}* Closing connection 0
```
