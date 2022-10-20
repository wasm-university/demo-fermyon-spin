#!/bin/bash
hey -n 1000 -c 50 -m GET \
-H "Content-Type: application/json" \
"http://localhost:9090"
