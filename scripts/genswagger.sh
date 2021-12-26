 #!/usr/bin/env bash
#
# Generate all swagger bindings.
#
set -ex


mkdir -p ./assets/swagger

protoc -I ./api \
  --openapiv2_out ./assets --openapiv2_opt logtostderr=true \
  ./api/proto/v1/hello_world.proto

mv  assets/proto/*/*.json ./assets/swagger
rm -rf  assets/proto
cp -r third_party/swagger-ui/* ./assets/swagger/

if [[ "$(uname)" -eq "Darwin" ]]
then
  sed -i '' 's/https:\/\/petstore.swagger.io\/v2\/swagger.json/.\/hello_world.swagger.json/g'  ./assets/swagger/index.html 
else
  sed -i  's/https:\/\/petstore.swagger.io\/v2\/swagger.json/.\/hello_world.swagger.json/g'  ./assets/swagger/index.html 
fi

alias swagger='docker run --rm -it  --user $(id -u):$(id -g) -e GOPATH=$(go env GOPATH):/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger:v0.28.0'
# swagger generate client -f ./assets/swagger/hello_world.swagger.json

mkdir -p ./pkg/sdk
swagger generate client -f ./assets/swagger/hello_world.swagger.json -t ./pkg/sdk

# swagger generate client -h