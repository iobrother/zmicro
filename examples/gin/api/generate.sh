#!/bin/sh

echo 'Generating api'
PROTOS=$(find ./api -type f -name '*.proto')

for PROTO in $PROTOS; do
  echo $PROTO
  protoc \
    -I. -I$(dirname $PROTO) \
    -I../../third_party \
    --gofast_out=. \
    --gofast_opt paths=source_relative \
    --zmicro-gin_out=. \
    --zmicro-gin_opt paths=source_relative \
    --zmicro-gin_opt allow_empty_patch_body=true \
    --zmicro-resty_out=. \
    --zmicro-resty_opt paths=source_relative \
    --zmicro-resty_opt use_invoke2=true \
    $PROTO
done
