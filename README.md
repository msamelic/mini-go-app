### mini-go-app

Minimal application  
with a few basic libraries (zap, gin, env)  
for copy-paste usage instead of creating empty projects

### Env vars

| Name       | Example Value | Default Value | Description |
|------------|---------------|---------------|-------------|
| `APP_NAME` |               | `mini-go-app` |             |
| `PORT`     |               | `8080`        |             |

### API

#### GET /health
```bash
$ curl 'localhost:8080/health'
{"status":"UP"}
```