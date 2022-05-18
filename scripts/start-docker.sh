#!/bin/bash

KEY="mykey"
CHAINID="enron_3000-3"
MONIKER="mymoniker"
DATA_DIR=$(mktemp -d -t enron-datadir.XXXXX)

echo "create and add new keys"
./enrond keys add $KEY --home $DATA_DIR --no-backup --chain-id $CHAINID --algo "eth_secp256k1" --keyring-backend test
echo "init Enron with moniker=$MONIKER and chain-id=$CHAINID"
./enrond init $MONIKER --chain-id $CHAINID --home $DATA_DIR
echo "prepare genesis: Allocate genesis accounts"
./enrond add-genesis-account \
"$(./enrond keys show $KEY -a --home $DATA_DIR --keyring-backend test)" 1000000000000000000aenron,1000000000000000000stake \
--home $DATA_DIR --keyring-backend test
echo "prepare genesis: Sign genesis transaction"
./enrond gentx $KEY 1000000000000000000stake --keyring-backend test --home $DATA_DIR --keyring-backend test --chain-id $CHAINID
echo "prepare genesis: Collect genesis tx"
./enrond collect-gentxs --home $DATA_DIR
echo "prepare genesis: Run validate-genesis to ensure everything worked and that the genesis file is setup correctly"
./enrond validate-genesis --home $DATA_DIR

echo "starting enron node $i in background ..."
./enrond start --pruning=nothing --rpc.unsafe \
--keyring-backend test --home $DATA_DIR \
>$DATA_DIR/node.log 2>&1 & disown

echo "started enron node"
tail -f /dev/null