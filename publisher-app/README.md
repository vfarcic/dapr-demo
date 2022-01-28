```bash
docker image build --tag vfarcic/dapr-publisher-app .

docker image tag vfarcic/dapr-publisher-app vfarcic/dapr-publisher-app:0.1.0

docker image push vfarcic/dapr-publisher-app

docker image push vfarcic/dapr-publisher-app:0.1.0
```