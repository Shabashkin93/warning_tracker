# Curl requests

## Create record

```bash
curl --request POST \
  --url http://${SERVER_ADDRESS}:${SERVER_PORT}/v1/warning \
  --header 'Content-Type: multipart/form-data' \
  --header 'User-Agent: insomnia/8.5.1' \
  --form commit=beffe2b9a727c481c8a4896edb1783a054ac084c \
  --form branch=develop \
  --form build_log=@/run/media/shabashkin/pool/projects/warning_tracker/resources/testfile \
  --form created_by=Shabashkin \
  --form created_at=2023-12-06T20:07:41.137Z
```

## Get status

```bash
curl --request GET \
  --url http://${SERVER_ADDRESS}:${SERVER_PORT}/v1/status \
  --header 'User-Agent: insomnia/8.5.1'
```

## Get metrics for Prometheus

```bash
curl --request GET \
  --url http://${SERVER_ADDRESS}:${SERVER_PORT}/metrics \
  --header 'User-Agent: insomnia/8.4.1' \
  --header 'content-type: multipart/form-data'
```