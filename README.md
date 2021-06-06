# Lightweight static file server
![publish container workflow](https://github.com/jkandasa/static-file-server/actions/workflows/publish_container.yaml/badge.svg)
![publish executable workflow](https://github.com/jkandasa/static-file-server/actions/workflows/publish_executable.yaml/badge.svg)

Lightweight static file server serves static files.

### Setup in docker
```bash
docker run --rm --name static-file-server \
  --publish 8080:8080 \
  --volume $PWD/data:/data \
  --env TZ="Asia/Kolkata" \
  --env BRAND_NAME="Lightweight Static File Server" \
  jkandasa/static-file-server:master
```

### Binary execution
```bash
$ ./static-file-server-master-linux-amd64 -port 8080 -dir /data
```

## Download
* [Container image](https://hub.docker.com/r/jkandasa/static-file-server)
* [Executable binaries](https://github.com/jkandasa/static-file-server/releases)
