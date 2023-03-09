### Steps ###

## Clone the wasm github repo 
git clone https://github.com/CosmWasm/wasmd


## clone the cosmwasm contract templates repo
git clone https://github.com/cosmwasm/cw-examples


## Run the docker container
docker-compose up -d


## Interact with the container
docker exec -it wasmd bash

## Change dir to the wasm source code
cd /wasm

## Build the wasmd binary
make install 

## Lets work in the user home dir
cd /home

## Install the rust cross-compile env for wasm32
rustup target add wasm32-unknown-unknown

## Change the dir to the contract templates 
cd /home/cw/contracts 

## As for this tutorial, we are building the nameservice contract, 
## Lets move to the nameservice dir 
cd ./nameservice 

## Build the nameserivce contract binary
RUSTFLAGS='-C link-arg=-s' cargo wasm

## Lets work in the user home dir
cd /home/

## Run the init script to run the wasm chain 
./init_script

## Change the minimum-gas-price in the following file
vim /root/.wasmd/config/app.toml


## Restart the chain 
wasmd start



## Deploy the smart contract 
## Note: cw_nameservice.wasm is the rust binary built for the  wasm chain 
wasmd tx wasm store /home/cw/contracts/nameservice/target/wasm32-unknown-unknown/release/cw_nameservice.wasm --from main  $TXFLAG -y -b block


### Intitialize the smart contract with some params 
INIT='{"purchase_price":{"amount":"100","denom":"stake"},"transfer_price":{"amount":"999","denom":"stake"}}'
wasmd tx wasm instantiate 1 "$INIT" --from main - --label "awesome name service" --no-admin $TXFLAG



### Sample query to the smart contract
### Note: The query over here is just querying the contract version
CONTRACT=$(wasmd query wasm list-contract-by-code $CODE_ID --output json | jq -r '.contracts[-1]')
wasmd query wasm contract $CONTRACT
