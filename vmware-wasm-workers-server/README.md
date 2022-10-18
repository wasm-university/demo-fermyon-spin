
## Install

```bash
curl https://raw.githubusercontent.com/vmware-labs/wasm-workers-server/main/install.sh | bash && \
  wws --help
```

```bash
Wasm Workers Server will be installed in /usr/local/bin.
This requires sudo permissions. If you prefer to install it
in your current directory, run the installer with --local.
If you want it to be global, just type your password:
🧹 Cleaning up
🚀 Wasm Workers Server (wws) was installed correctly!
You can now try it: wws --help
Usage: wws [OPTIONS] [PATH]

Arguments:
  [PATH]  Folder to read WebAssembly modules from [default: .]

Options:
      --host <HOSTNAME>  Hostname to initiate the server [default: 127.0.0.1]
  -p, --port <PORT>      Port to initiate the server [default: 8080]
  -h, --help             Print help information
  -V, --version          Print version information
```

```bash
curl https://raw.githubusercontent.com/vmware-labs/wasm-workers-server/main/examples/js-basic/handler.js \
  -o ./index.js && \
  wws .
```

## Create a Rust Worker

```bash
mkdir worker
cd worker
cargo init
cargo build --target wasm32-wasi --release

cd target/wasm32-wasi/release && \
  wws .
```

http://127.0.0.1:8080/worker
