#!/usr/bin/env bash

ACTION=${1:-help}

case "${ACTION}" in
    run)
        go run main.go
        ;;

    clean)
        [[ -d "build" ]] && rm -rf "build"
        ;;
        
    build)
        mkdir -p build \
          && go build -o ./build/go-grpc-api \
        ;;
    *)
        echo "Bad usage"
        exit 1
        ;;
esac

exit $?
