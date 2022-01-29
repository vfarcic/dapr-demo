```bash
export VERSION=0.1.19

docker image build --tag ghcr.io/vfarcic/dapr-publications .

docker image tag ghcr.io/vfarcic/dapr-publications ghcr.io/vfarcic/dapr-publications:$VERSION

docker image push ghcr.io/vfarcic/dapr-publications

docker image push ghcr.io/vfarcic/dapr-publications:$VERSION
```