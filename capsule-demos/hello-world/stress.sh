#!/bin/bash
hey -n 1000 -c 100 -m GET \
-H "Content-Type: application/json" \
"http://localhost:9090"
