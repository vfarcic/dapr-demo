```bash
export VERSION=0.1.1

docker image build --tag ghcr.io/vfarcic/dapr-speech-to-text .

docker image tag ghcr.io/vfarcic/dapr-speech-to-text ghcr.io/vfarcic/dapr-speech-to-text:$VERSION

docker image push ghcr.io/vfarcic/dapr-speech-to-text

docker image push ghcr.io/vfarcic/dapr-speech-to-text:$VERSION
```