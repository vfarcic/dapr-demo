```bash
docker image build --tag ghcr.io/vfarcic/dapr-tweet .

docker image tag ghcr.io/vfarcic/dapr-tweet ghcr.io/vfarcic/dapr-tweet:0.1.0

docker image push ghcr.io/vfarcic/dapr-tweet

docker image push ghcr.io/vfarcic/dapr-tweet:0.1.0
```