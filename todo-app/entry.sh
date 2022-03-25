#!/bin/sh
echo hoge
if [ -z "${AWS_LAMBDA_RUNTIME_API}" ]; then
  exec /usr/bin/aws-lambda-rie "/src/app/main"
else
  exec "/src/app/main"
fi
