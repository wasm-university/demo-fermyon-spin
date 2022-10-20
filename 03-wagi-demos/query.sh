#!/bin/bash
########################
# include the magic
########################
. ../demo-magic.sh

# hide the evidence
clear

# Put your stuff here
pei "curl -i localhost:3000/hello-swift"
echo ""
pei "curl -i localhost:3000/hello-go?name=bob"
echo ""
pei "curl -i -X POST localhost:3000/hello-go -d \"😀 Hello there\""
echo ""
