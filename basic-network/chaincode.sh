set -ev

docker exec cli1 peer chaincode install -n share-cc-1 -v 1.0 -p chaincode/go/channel1
#docker exec cli2 peer chaincode install -n share-cc-2 -v 1.0 -p chaincode/go/channel2
sleep 1

docker exec cli1 peer chaincode instantiate -o orderer.sharing.com:7050 -C channel1share -n share-cc-1 -v 1.0 -c '{"Args":[""]}' -P "OR ('AgencyOrg.member')"
#docker exec cli2 peer chaincode instantiate -o orderer.sharing.com:7050 -C channel2share -n share-cc-2 -v 1.0 -c '{"Args":[""]}' -P "OR ('Share1Org.member','Share2Org.member')"
sleep 10

docker exec cli1 peer chaincode query -o orderer.sharing.com:7050 -C channel1share -n share-cc-1 -c '{"function":"checkUser","Args":["test","test"]}'

