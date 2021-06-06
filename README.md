# Lightweight static file server

```bash
docker run --rm --name static-file-server \
  --publish 8080:8080 \
  --volume $PWD/data:/data \
  --env TZ="Asia/Kolkata" \
  --env BRAND_NAME="Lightweight Static File Server" \
  jkandasa/static-file-server:master
```