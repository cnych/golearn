#!/bin/bash
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
if [ -d $(dirname $DIR)/stardust.go ]; then
    export GOPATH="$DIR:$(dirname $DIR)/stardust.go"
else
    export GOPATH="$DIR:$(dirname $DIR)/stardust-go"
fi
