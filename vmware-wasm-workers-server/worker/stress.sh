#!/bin/bash
hey -n 10000 -c 50 -m GET \
-H "Content-Type: application/json" \
"http://localhost:8080/worker"
