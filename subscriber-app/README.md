```bash
docker image build --tag vfarcic/dapr-subscriber-app .

docker image tag vfarcic/dapr-subscriber-app vfarcic/dapr-subscriber-app:0.1.5

docker image push vfarcic/dapr-subscriber-app

docker image push vfarcic/dapr-subscriber-app:0.1.5
```