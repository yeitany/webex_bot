
### build and run locally

make build

docker build -t webex_bot:latest .

### test github request:
docker run -i -p 127.0.0.1:8080:8080 webex_bot:latest /github

curl -i -X POST localhost:8080/webhooks/github -H "Content-type: application/json" -d "@githubPayloadExample.json" -H "X-GitHub-Event: pull_request"

### test gcp request:
docker run -i -p 127.0.0.1:8080:8080 webex_bot:latest /gcp

curl -i -X POST localhost:8080/webhooks/gcp -H "application/json" -d "@gcpNotificationPayload.json"

