#!/usr/bin/env bash

set -euo pipefail

KEYS[0]="dev0"
KEYS[1]="dev1"
KEYS[2]="dev2"
CHAINID="${CHAINID:-mud_168169-1}"
MONIKER="${MONIKER:-localtestnet}"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="${KEYRING:-test}"
KEYALGO="${KEYALGO:-eth_secp256k1}"
LOGLEVEL="${LOGLEVEL:-info}"
# Set dedicated home directory for the mudd instance
HOMEDIR="${HOMEDIR:-$HOME/.tmp-mudd}"
# to trace evm
#TRACE="--trace"
TRACE="${TRACE:-}"
SKIP_START="${LOCAL_NODE_SKIP_START:-0}"
MODE="${1:-}"

# Path variables
CONFIG=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# validate dependencies are installed
command -v jq >/dev/null 2>&1 || {
	echo >&2 "jq not installed. More info: https://stedolan.github.io/jq/download/"
	exit 1
}

command -v bc >/dev/null 2>&1 || {
	echo >&2 "bc not installed."
	exit 1
}

# Reinstall daemon
make install

resolve_daemon() {
	if command -v mudd >/dev/null 2>&1; then
		command -v mudd
		return
	fi

	local gobin
	gobin="$(go env GOBIN)"
	if [[ -z "$gobin" ]]; then
		gobin="$(go env GOPATH)/bin"
	fi

	if [[ -x "$gobin/mudd" ]]; then
		echo "$gobin/mudd"
		return
	fi

	echo >&2 "mudd binary not found after make install."
	exit 1
}

DAEMON="$(resolve_daemon)"

# User prompt if an existing local node configuration is found.
if [ -d "$HOMEDIR" ]; then
	printf "\nAn existing folder at '%s' was found. You can choose to delete this folder and start a new local node with new keys from genesis. When declined, the existing local node is started. \n" "$HOMEDIR"
	echo "Overwrite the existing configuration and start a new local node? [y/n]"
	read -r overwrite
else
	overwrite="Y"
fi


# Setup local node if overwrite is set to Yes, otherwise skip setup
if [[ $overwrite == "y" || $overwrite == "Y" ]]; then
	# Remove the previous folder
	rm -rf "$HOMEDIR"

	# Set client config
	"$DAEMON" config keyring-backend "$KEYRING" --home "$HOMEDIR"
	"$DAEMON" config chain-id "$CHAINID" --home "$HOMEDIR"

	# If keys exist they should be deleted
	for KEY in "${KEYS[@]}"; do
		"$DAEMON" keys add "$KEY" --keyring-backend "$KEYRING" --algo "$KEYALGO" --home "$HOMEDIR"
	done

	# Set moniker and chain-id for MUD (Moniker can be anything, chain-id must be an integer)
	"$DAEMON" init "$MONIKER" -o --chain-id "$CHAINID" --home "$HOMEDIR"

	# Change parameter token denominations to amud
	jq '.app_state["staking"]["params"]["bond_denom"]="amud"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["crisis"]["constant_fee"]["denom"]="amud"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="amud"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["evm"]["params"]["evm_denom"]="amud"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
	jq '.app_state["inflation"]["inflation_amount"]["denom"]="amud"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Set gas limit in genesis
	jq '.consensus_params["block"]["max_gas"]="10000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	if [[ "$OSTYPE" == "darwin"* ]]; then
		sed -i '' 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "$CONFIG"
	else
		sed -i 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "$CONFIG"
	fi

	if [[ "$MODE" == "pending" ]]; then
		if [[ "$OSTYPE" == "darwin"* ]]; then
			sed -i '' 's/timeout_propose = "3s"/timeout_propose = "30s"/g' "$CONFIG"
			sed -i '' 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' "$CONFIG"
			sed -i '' 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' "$CONFIG"
			sed -i '' 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' "$CONFIG"
			sed -i '' 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' "$CONFIG"
			sed -i '' 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' "$CONFIG"
			sed -i '' 's/timeout_commit = "1s"/timeout_commit = "150s"/g' "$CONFIG"
			sed -i '' 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' "$CONFIG"
		else
			sed -i 's/timeout_propose = "3s"/timeout_propose = "30s"/g' "$CONFIG"
			sed -i 's/timeout_propose_delta = "500ms"/timeout_propose_delta = "5s"/g' "$CONFIG"
			sed -i 's/timeout_prevote = "1s"/timeout_prevote = "10s"/g' "$CONFIG"
			sed -i 's/timeout_prevote_delta = "500ms"/timeout_prevote_delta = "5s"/g' "$CONFIG"
			sed -i 's/timeout_precommit = "1s"/timeout_precommit = "10s"/g' "$CONFIG"
			sed -i 's/timeout_precommit_delta = "500ms"/timeout_precommit_delta = "5s"/g' "$CONFIG"
			sed -i 's/timeout_commit = "1s"/timeout_commit = "150s"/g' "$CONFIG"
			sed -i 's/timeout_broadcast_tx_commit = "10s"/timeout_broadcast_tx_commit = "150s"/g' "$CONFIG"
		fi
	fi

    # enable prometheus metrics
    if [[ "$OSTYPE" == "darwin"* ]]; then
        sed -i '' 's/prometheus = false/prometheus = true/' "$CONFIG"
        sed -i '' 's/prometheus-retention-time = 0/prometheus-retention-time  = 1000000000000/g' "$APP_TOML"
        sed -i '' 's/enabled = false/enabled = true/g' "$APP_TOML"
    else
        sed -i 's/prometheus = false/prometheus = true/' "$CONFIG"
        sed -i 's/prometheus-retention-time  = "0"/prometheus-retention-time  = "1000000000000"/g' "$APP_TOML"
        sed -i 's/enabled = false/enabled = true/g' "$APP_TOML"
    fi
	
	# Change proposal periods to pass within a reasonable time for local testing
	sed -i.bak 's/"max_deposit_period": "172800s"/"max_deposit_period": "30s"/g' "$HOMEDIR"/config/genesis.json
	sed -i.bak 's/"voting_period": "172800s"/"voting_period": "30s"/g' "$HOMEDIR"/config/genesis.json

	# set custom pruning settings
	sed -i.bak 's/pruning = "default"/pruning = "custom"/g' "$APP_TOML"
	sed -i.bak 's/pruning-keep-recent = "0"/pruning-keep-recent = "2"/g' "$APP_TOML"
	sed -i.bak 's/pruning-interval = "0"/pruning-interval = "10"/g' "$APP_TOML"

	# Allocate genesis accounts (cosmos formatted addresses)
	for KEY in "${KEYS[@]}"; do
		"$DAEMON" add-genesis-account "$KEY" 100000000000000000000000000amud --keyring-backend "$KEYRING" --home "$HOMEDIR"
	done

	# bc is required to add these big numbers
	total_supply=$(echo "${#KEYS[@]} * 100000000000000000000000000" | bc)
	jq -r --arg total_supply "$total_supply" '.app_state["bank"]["supply"] |= (map(select(.denom != "amud")) + [{"denom":"amud","amount":$total_supply}])' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

	# Sign genesis transaction
	"$DAEMON" gentx "${KEYS[0]}" 1000000000000000000000amud --keyring-backend "$KEYRING" --chain-id "$CHAINID" --home "$HOMEDIR"
	## In case you want to create multiple validators at genesis
	## 1. Back to `mudd keys add` step, init more keys
	## 2. Back to `mudd add-genesis-account` step, add balance for those
	## 3. Clone this ~/.mudd home directory into some others, let's say `~/.clonedMudd`
	## 4. Run `gentx` in each of those folders
	## 5. Copy the `gentx-*` folders under `~/.clonedMudd/config/gentx/` folders into the original `~/.mudd/config/gentx`

	# Collect genesis tx
	"$DAEMON" collect-gentxs --home "$HOMEDIR"

	# Run this to ensure everything worked and that the genesis file is setup correctly
	"$DAEMON" validate-genesis --home "$HOMEDIR"

	if [[ "$MODE" == "pending" ]]; then
		echo "pending mode is on, please wait for the first block committed."
	fi
fi

if [[ "$SKIP_START" == "1" ]]; then
	exit 0
fi

# Expose JSON-RPC to the host network for local development.
if [[ "$OSTYPE" == "darwin"* ]]; then
	sed -i '' 's/address = "127.0.0.1:8545"/address = "0.0.0.0:8545"/g' "$APP_TOML"
else
	sed -i 's/address = "127.0.0.1:8545"/address = "0.0.0.0:8545"/g' "$APP_TOML"
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
START_ARGS=(
	start
	--metrics
	--log_level "$LOGLEVEL"
	--minimum-gas-prices=0.0001amud
	--json-rpc.api eth,txpool,personal,net,debug,web3
	--api.enable
	--home "$HOMEDIR"
)

if [[ -n "$TRACE" ]]; then
	START_ARGS+=("$TRACE")
fi

"$DAEMON" "${START_ARGS[@]}"
