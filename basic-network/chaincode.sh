set -ev

docker exec cli peer chaincode install -n share-cc-ch1 -v 1.0 -p chaincode/go/channel1
docker exec -e "CORE_PEER_LOCALMSPID=AgencyOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/agency.sharing.com/users/Admin@agency.sharing.com/msp" -e "CORE_PEER_ADDRESS=peer0.agency.sharing.com:7051" cli peer chaincode install -n share-cc-ch1 -v 1.0 -p chaincode/go/channel1
sleep 1

docker exec cli peer chaincode instantiate -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -v 1.0 -c '{"Args":[""]}' -P "OR ('AgencyOrg.member','Share1Org.member','Share2Org.member')"
docker exec -e "CORE_PEER_LOCALMSPID=AgencyOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/agency.sharing.com/users/Admin@agency.sharing.com/msp" -e "CORE_PEER_ADDRESS=peer0.agency.sharing.com:7051" cli peer chaincode instantiate -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -v 1.0 -c '{"Args":[""]}' -P "OR ('AgencyOrg.member','Share1Org.member','Share2Org.member')"
sleep 10
