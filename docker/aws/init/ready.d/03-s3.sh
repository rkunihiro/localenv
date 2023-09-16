#!/bin/bash

awslocal s3 mb s3://hello
awslocal s3 cp /init/keys/s3/hello/ s3://hello/ --recursive
awslocal s3 ls s3://hello

awslocal s3api put-bucket-policy \
  --bucket hello \
  --policy file:///init/keys/s3/bucket-policy.json
awslocal s3api put-public-access-block \
  --bucket hello \
  --public-access-block-configuration \
  "BlockPublicAcls=true,IgnorePublicAcls=true,BlockPublicPolicy=true,RestrictPublicBuckets=true"

# browse
# https://<bucket-name>.s3.<region>.localhost.localstack.cloud:4566/<key-name>
# https://hello.s3.ap-northeast-1.localhost.localstack.cloud:4566/index.html
