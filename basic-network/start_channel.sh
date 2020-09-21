set -ev

export MSYS_NO_PATHCONV=1

docker-compose -f docker-compose.yaml down
docker-compose -f docker-compose.yaml up -d orderer.sharing.com peer0.agency.sharing.com peer1.agency.sharing.com peer0.share1.sharing.com peer1.share1.sharing.com peer0.share2.sharing.com peer1.share2.sharing.com cli1 cli2

export FABRIC_START_TIMEOUT=10
sleep ${FABRIC_START_TIMEOUT}

# Create the channel
docker exec cli1 peer channel create -o orderer.sharing.com:7050 -c channel1share -f /etc/hyperledger/configtx/channel1.tx
docker exec cli2 peer channel create -o orderer.sharing.com:7050 -c channel2share -f /etc/hyperledger/configtx/channel2.tx

# Join & Update Anchor node (channel1)
docker exec cli1 peer channel join -b channel1share.block
docker exec cli1 peer channel update -o orderer.sharing.com:7050 -c channel1share -f /etc/hyperledger/configtx/AgencyOrgAnchors.tx

docker exec -e "CORE_PEER_ADDRESS=peer1.agency.sharing.com:7051" cli1 peer channel join -b channel1share.block

# Join & Update Anchor node (channel2)
docker exec cli2 peer channel join -b channel2share.block
docker exec cli2 peer channel update -o orderer.sharing.com:7050 -c channel2share -f /etc/hyperledger/configtx/Share1OrgAnchors.tx

docker exec -e "CORE_PEER_ADDRESS=peer1.share1.sharing.com:7051" cli2 peer channel join -b channel2share.block

docker exec -e "CORE_PEER_LOCALMSPID=Share2Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/share2.sharing.com/users/Admin@share2.sharing.com/msp" -e "CORE_PEER_ADDRESS=peer0.share2.sharing.com:7051" cli2 peer channel join -b channel2share.block
docker exec -e "CORE_PEER_LOCALMSPID=Share2Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/share2.sharing.com/users/Admin@share2.sharing.com/msp" -e "CORE_PEER_ADDRESS=peer0.share2.sharing.com:7051" cli2 peer channel update -o orderer.sharing.com:7050 -c channel2share -f /etc/hyperledger/configtx/Share2OrgAnchors.tx

docker exec -e "CORE_PEER_LOCALMSPID=Share2Org" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/share2.sharing.com/users/Admin@share2.sharing.com/msp" -e "CORE_PEER_ADDRESS=peer1.share2.sharing.com:7051" cli2 peer channel join -b channel2share.block

