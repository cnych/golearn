#!/bin/bash
DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
export GOPATH="$DIR:$(dirname $DIR)/starjazz-go"
export PATH="$PATH:$GOPATH/bin"
