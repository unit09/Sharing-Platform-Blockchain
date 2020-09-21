set -ev

docker-compose down

docker stop $(docker ps -aq)

docker rm -f $(docker ps -aq)

#docker images

#docker rmi [ID]