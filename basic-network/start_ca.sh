set -ev

docker-compose -f docker-compose-ca.yaml up -d ca.share1.sharing.com ca.agency.sharing.com
sleep 1
cd $GOPATH/src/Sharing-Platform/application/sdk
node enrollAdmin_share1.js
sleep 2
node enrollAdmin_agency.js
sleep 2
node registerUser_share1.js
sleep 2
node registerUser_agency.js