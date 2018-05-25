@echo off
set GRPC_DEFAULT_SSL_ROOTS_FILE_PATH=ca-certificates.crt
grpc_cli.exe -enable_ssl %*
