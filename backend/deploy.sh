bash backend/build.sh
$(aws ecr get-login --no-include-email --region eu-west-1)
docker tag fixtheplanet:latest 515136113268.dkr.ecr.eu-west-1.amazonaws.com/fixtheplanet:latest
docker push 515136113268.dkr.ecr.eu-west-1.amazonaws.com/fixtheplanet:latest