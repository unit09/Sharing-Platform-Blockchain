# remove previous crypto material and config transactions
rm -fr config/*
rm -fr crypto-config/*

# generate crypto material
./bin/cryptogen generate --config=./crypto-config.yaml

# generate genesis block for orderer
mkdir config
./bin/configtxgen -profile OrdererGenesis -outputBlock ./config/genesis.block

# generate channel configuration transaction
./bin/configtxgen -profile Channel1 -outputCreateChannelTx ./config/channel1.tx -channelID channel1share
./bin/configtxgen -profile Channel2 -outputCreateChannelTx ./config/channel2.tx -channelID channel2share

# generate anchor peer transaction
./bin/configtxgen -profile Channel1 -outputAnchorPeersUpdate ./config/AgencyOrgAnchors.tx -channelID channel1share -asOrg AgencyOrg

./bin/configtxgen -profile Channel2 -outputAnchorPeersUpdate ./config/Share1OrgAnchors.tx -channelID channel2share -asOrg Share1Org
./bin/configtxgen -profile Channel2 -outputAnchorPeersUpdate ./config/Share2OrgAnchors.tx -channelID channel2share -asOrg Share2Org
