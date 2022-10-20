#!/bin/bash
cd worker
cargo build --target wasm32-wasi --release
