docker exec cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"getAllShareRecord","Args":[""]}'
sleep 2
docker exec cli peer chaincode invoke -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"startShare","Args":["1Q2W","A1","Daegu"]}'
sleep 2
docker exec cli peer chaincode invoke -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"endShare","Args":["3E4R","A1","Daegu"]}'
sleep 2
docker exec cli peer chaincode invoke -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"startShare","Args":["2W5T","B2","Seoul"]}'
sleep 2
docker exec cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"getAllShareRecord","Args":[""]}'
sleep 2
docker exec cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"getShareStart","Args":[""]}'
sleep 2
docker exec cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"getShareRecordByLocation","Args":["Daegu"]}'
sleep 2
docker exec cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"countAllShareRecordByLocation","Args":[""]}'
sleep 2
docker exec cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"getAllPlace","Args":[""]}'
sleep 2
docker exec -e "CORE_PEER_LOCALMSPID=AgencyOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/agency.sharing.com/users/Admin@agency.sharing.com/msp" -e "CORE_PEER_ADDRESS=peer0.agency.sharing.com:7051" cli peer chaincode invoke -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"setPlace","Args":["KKK","Mars"]}'
sleep 2
docker exec cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"getAllPlace","Args":[""]}'
sleep 2
docker exec cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"getUserShareRecord","Args":["2W5T"]}'
#docker exec -e "CORE_PEER_LOCALMSPID=AgencyOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/agency.sharing.com/users/Admin@agency.sharing.com/msp" -e "CORE_PEER_ADDRESS=peer0.agency.sharing.com:7051" cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"getAllShareRecord","Args":[""]}'
#sleep 2
#docker exec -e "CORE_PEER_LOCALMSPID=AgencyOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/agency.sharing.com/users/Admin@agency.sharing.com/msp" -e "CORE_PEER_ADDRESS=peer0.agency.sharing.com:7051" cli peer chaincode invoke -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"startShare","Args":["1Q3W","B2","Seoul"]}'
#sleep 2
#docker exec -e "CORE_PEER_LOCALMSPID=AgencyOrg" -e "CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/agency.sharing.com/users/Admin@agency.sharing.com/msp" -e "CORE_PEER_ADDRESS=peer0.agency.sharing.com:7051" cli peer chaincode query -o orderer.sharing.com:7050 -C sharechannel1 -n share-cc-ch1 -c '{"function":"getAllShareRecord","Args":[""]}'