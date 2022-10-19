#!/bin/bash
hey -n 10 -c 10 -m GET \
-H "Content-Type: application/json" \
"http://localhost:8080/worker"
