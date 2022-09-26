use suborbital::runnable::*;

use suborbital::http; // 👋 HOST FUNCTION
use suborbital::log; // 👋 HOST FUNCTION
use std::str;

struct Hello{}

/*
./demo hello/hello.wasm http://localhost:9090
*/

impl Runnable for Hello {
    fn run(&self, input: Vec<u8>) -> Result<Vec<u8>, RunErr> {
        let in_string = String::from_utf8(input).unwrap();

        // "http://localhost:9090"
        let query_result = match http::get(&in_string, None) { // 👋 HOST FUNCTION
            Ok(data) => data,
            Err(e) => e.message.as_bytes().to_vec()
        };

        let message = match str::from_utf8(&query_result) {
            Ok(v) => v,
            Err(e) => panic!("Invalid UTF-8 sequence: {}", e),
        };

        log::info(&message); // 👋 HOST FUNCTION


        Ok(String::from(format!("hello {}", in_string)).as_bytes().to_vec())
    }
}


// initialize the runner, do not edit below //
static RUNNABLE: &Hello = &Hello{};

#[no_mangle]
pub extern fn _start() {
    use_runnable(RUNNABLE);
}
