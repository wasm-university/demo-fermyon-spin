#!/bin/bash
########################
# include the magic
########################
. ../demo-magic.sh

# hide the evidence
clear

# Put your stuff here
pei "tinygo build -o hello-go.wasm -target wasi ./hello.go"
pei "ls -lh *.wasm"
