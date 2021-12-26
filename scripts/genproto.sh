 #!/usr/bin/env bash
#
# Generate all protobuf bindings.
#
set -ex

if ! [[ "$0" =~ scripts/genproto.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

mkdir -p ./pkg/proto
protoc -I ./api \
  --go_out ./pkg/proto \
  --go-grpc_out ./pkg/proto \
  --grpc-gateway_out  ./pkg/proto \
  ./api/proto/v1/hello_world.proto
