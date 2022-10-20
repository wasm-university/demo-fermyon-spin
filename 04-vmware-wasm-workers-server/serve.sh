#!/bin/bash
########################
# include the magic
########################
. ../demo-magic.sh

# hide the evidence
clear

# Put your stuff here
echo "open http://127.0.0.1:8080/worker"
pei "cd worker; wws ."


