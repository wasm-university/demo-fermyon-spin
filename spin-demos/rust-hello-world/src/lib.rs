use anyhow::Result;
use spin_sdk::{
    http::{Request, Response},
    http_component,
};

/// A simple Spin HTTP component.
#[http_component]
fn rust_hello_world(req: Request) -> Result<Response> {

  println!("{:?}", req.headers());
  println!("{:?}", req.body());
  // logs:
  // bat $HOME/.spin/rust-hello-world/logs/rust-hello-world_stdout.txt
  // bat $HOME/.spin/rust-hello-world/logs/rust-hello-world_stderr.txt

    Ok(http::Response::builder()
        .status(200)
        .header("MESSAGE", "🖖 Vulcan salutation")
        .body(Some("👋 Hello World! 🌍".into()))?)
}
