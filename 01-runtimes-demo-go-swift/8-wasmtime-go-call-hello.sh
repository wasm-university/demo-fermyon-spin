#!/bin/bash

########################
# include the magic
########################
. ../demo-magic.sh

# hide the evidence
clear

# Put your stuff here
pei "wasmtime --invoke hello hello-go.wasm \"Bob Morane\""

