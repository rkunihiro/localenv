#!/bin/bash

awslocal dynamodb create-table \
  --cli-input-json file:///init/keys/dynamodb/todo.json

awslocal dynamodb list-tables
