set -ev

docker exec cli peer chaincode install -n share-cc-1 -v 1.0 -p chaincode/go/channel1
sleep 1

docker exec cli peer chaincode instantiate -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-1 -v 1.0 -c '{"Args":[""]}' -P "OR ('AgencyOrg.member','Share1Org.member','Share2Org.member')"
sleep 10

docker exec cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-1 -c '{"function":"checkUser","Args":["test","test"]}'

