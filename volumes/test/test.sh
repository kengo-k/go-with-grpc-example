#!/bin/bash

#
# NOTE: 事前にserverを起動しておくこと
#
result=$(grpcurl -plaintext -d '{"values": [1,2,3,4,5,6,7,8,9,10]}' localhost:3000 sample.SampleService/Add)

got=$(echo $result | jq .sum)
expected=55

if [ ${got} = ${expected} ]; then
  echo OK
else
  echo NG
fi
