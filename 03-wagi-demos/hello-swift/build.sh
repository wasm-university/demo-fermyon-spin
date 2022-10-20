#!/bin/bash
########################
# include the magic
########################
. ../../demo-magic.sh

# hide the evidence
clear

# Put your stuff here
pei "swiftc -target wasm32-unknown-wasi hello.swift -o hello.wasm"
pei "ls -lh *.wasm"


