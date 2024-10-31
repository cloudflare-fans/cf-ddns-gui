#!/bin/sh

if [ -z "$GOPATH" ]; then
    echo GOPATH environment variable not set
    exit
fi

if [ ! -e "$GOPATH/bin/2goarray" ]; then
    echo "Installing 2goarray..."
    go install github.com/cloudflare-fans/2goarray
    if [ $? -ne 0 ]; then
        echo Failure executing go install github.com/cloudflare-fans/2goarray
        exit
    fi
fi

if [ -z "$1" ]; then
    echo Please specify a PNG file
    exit
fi

if [ ! -f "$1" ]; then
    echo $1 is not a valid file
    exit
fi

OUTPUT_DIR=${1%.*}
mkdir "${OUTPUT_DIR}"

BASENAME=$(basename "${1}")
PACKAGE=${BASENAME%.*}

OUTPUT=$OUTPUT_DIR/icon_unix.go
echo Generating $OUTPUT
echo "//+build linux darwin" > $OUTPUT
echo >> $OUTPUT
cat "$1" | $GOPATH/bin/2goarray Data "${PACKAGE}" >> $OUTPUT
if [ $? -ne 0 ]; then
    echo Failure generating $OUTPUT
    exit
fi
echo Finished
