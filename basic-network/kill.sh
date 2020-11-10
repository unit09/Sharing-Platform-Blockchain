set -ev

#docker-compose down
docker-compose -f docker-compose-ca.yaml stop
docker-compose stop
docker-compose -f docker-compose-ca.yaml kill && docker-compose -f docker-compose-ca.yaml down --volumes --remove-orphans
docker-compose -f docker-compose.yaml kill && docker-compose -f docker-compose.yaml down --volumes --remove-orphans

docker stop $(docker ps -aq)
docker rm -f $(docker ps -aq)

rm -rf $GOPATH/src/Sharing-Platform/application/wallet/

#docker images
#docker rmi [ID]

docker rmi -f $(docker images dev-* -q)