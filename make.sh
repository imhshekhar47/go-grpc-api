#!/usr/bin/env bash

ACTION=${1:-help}


clean_generarted() {
    find pb -name "*.pb.go" -exec rm {} \;
}

generate_pb() {
    protoc -I="proto"  --go_out=plugins=grpc:"pb" proto/*.proto
}

build() {
    [[ -d "build" ]] || mkdir -p "build"
    go build -o ./build/go-grpc-api 
}

case "${ACTION}" in
    run)
        go run main.go
        ;;

    clean)
        [[ -d "build" ]] && rm -rf "build"
        clean_generarted
        ;;

    generate)
        clean_generarted && generate_pb
        ;;

    build)
        clean_generarted && generate_pb && build
        ;;
    *)
        echo "Bad usage"
        exit 1
        ;;
esac

exit $?
