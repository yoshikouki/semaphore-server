# semaphore-server

SEMAPI is SEmaphore MAnagement API.

## Usage

```
package main

import (
	"github.com/yoshikouki/semapi/server"
)

func main() {
	semapi.Launch(server.Config{
		Port: 9876,
	})
}

```

## Look and feel

```
$ curl -X POST localhost:8686/semapi/lock \
    -H "Content-Type: application/json" \
    -d '{"lock_target": "org-repo-stage", "user":"test", "ttl":"10s"}'
{"expireDate":"2021/05/24 23:59:03","getLocked":"true","user":"test"}

$ curl -X POST localhost:8686/semapi/unlock \
    -H "Content-Type: application/json" \
    -d '{"unlock_target": "org-repo-stage", "user":"test"}'
{"getUnlock":"true","message":""}
```

And more: https://github.com/yoshikouki/semapi/blob/main/integration_test.go
