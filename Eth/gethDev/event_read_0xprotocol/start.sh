#!/bin/bash

CONTRACT="exchange"
solc --abi "$CONTRACT".sol | sed -n '4,$p' > "$CONTRACT".abi

PKG_NAME="exchange"
mkdir "$PKG_NAME"
if [ -f $GOBIN/abigen ]
then
    $GOBIN/abigen --abi="$CONTRACT".abi --pkg="$PKG_NAME" --out="$PKG_NAME"/"$PKG_NAME".go
fi
