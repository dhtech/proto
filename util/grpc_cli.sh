#!/usr/bin/env bash

cd $(dirname $0)

export GRPC_DEFAULT_SSL_ROOTS_FILE_PATH=ca-certificates.crt
if [[ "$(uname)" == "Linux" ]]; then
  ./grpc_cli_linux -enable_ssl "$@"
elif [[ "$(uname)" == "FreeBSD" ]]; then
  ./grpc_cli_freebsd -enable_ssl "$@"
elif [[ "$(uname)" == "MINGW64_NT-10.0" ]]; then
  ./grpc_cli.exe -enable_ssl "$@"
else
  echo "Unknown OS $(uname)"
fi
