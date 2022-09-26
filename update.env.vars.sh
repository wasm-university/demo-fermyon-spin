#!/bin/bash
SWIFT_VERSION="5.7.0"
export SWIFT_ROOT=$HOME/swift-wasm-${SWIFT_VERSION}-RELEASE
export PATH=$SWIFT_ROOT/usr/bin:"${PATH}"
swiftc --version
