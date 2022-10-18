use anyhow::Result;
use wasm_workers_rs::{
    handler,
    http::{self, Request, Response},
};

#[handler]
fn reply(_req: Request<String>) -> Result<Response<String>> {
    // Applied changes here to use the Response method. This requires changes
    // on signature and how it returns the data.
    let response = format!(
        "<!DOCTYPE html>
<body>
<h1>Hello World</h1>
</body>"
    );

    Ok(http::Response::builder()
        .status(200)
        .header("x-generated-by", "wasm-workers-server")
        .body(response.into())?)
}
