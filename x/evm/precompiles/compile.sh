#!/usr/bin/env bash

set -eo pipefail

commands=(git yarn jq abigen)
for cmd in "${commands[@]}"; do
  if ! command -v "$cmd" &>/dev/null; then
    echo "$cmd command not found, please install $cmd first" && exit 1
  fi
done

project_dir="$(git rev-parse --show-toplevel)"
if [ ! -d "$project_dir/solidity/contracts/node_modules" ]; then
  echo "===> Installing node modules"
  (cd "$project_dir/solidity" && yarn install)
fi

if [ -d "$project_dir/solidity/artifacts" ]; then
  echo "===> Cleaning artifacts"
  (cd "$project_dir/solidity" && yarn clean)
fi

echo "===> Compiling contracts"
(cd "$project_dir/solidity" && yarn compile)

[[ ! -d "$project_dir/x/evm/precompiles/contracts/artifacts" ]] && mkdir -p "$project_dir/x/evm/precompiles/contracts/artifacts"

# add core contracts
contracts=(IStaking IDistribution IBank ISlashing IEvidence IEpochs IGov IAuth IInflation)
# add 3rd party contracts

for contract in "${contracts[@]}"; do
  echo "===> Ethereum ABI wrapper code generator: $contract"
  pkg=$(echo "$contract" | tr '[:upper:]' '[:lower:]')
  pkg=${pkg:1}
  file_path=$(find "$project_dir/solidity/artifacts" -name "${contract}.json" -type f)
  jq -c '.abi' "$file_path" >"$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.abi"
  jq -r '.bytecode' "$file_path" >"$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.bin"
  abigen --abi "$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.abi" \
    --bin "$project_dir/x/evm/precompiles/contracts/artifacts/${contract}.bin" \
    --type "${contract}" --pkg ${pkg} \
    --out "$project_dir/x/evm/precompiles/${pkg}/${contract}.go"

  # Copy ABI files
  test_file_name="${contract:1}Test.json"
  test_file_path=$(find "$project_dir/solidity/artifacts" -name $test_file_name -type f)
  destination_path="$project_dir/x/evm/precompiles/${pkg}/abi"
  # Check if path exists
  if [ -d "${destination_path}" ]; then
    echo "===> Copy ABI files for module: $contract"
    cp "$file_path" "$destination_path/${contract}.json"
    cp "$test_file_path" "$destination_path/$test_file_name"
  fi
done

rm -rf "$project_dir/x/evm/precompiles/contracts"