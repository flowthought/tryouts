use std::io;

fn main() {
  let mut buffer = String::new();
  let mut inc_count: i32 = 0;
  let mut prev: i32 = std::i32::MAX;
  
  loop {
    match io::stdin().read_line(&mut buffer) {
      Ok(0) => {
        break;
      }
      Ok(_) => {
        // println!("{buffer:?}");
        match buffer.trim_end().parse::<i32>() {
          Ok(curr) => {
            if curr > prev {
              inc_count += 1;
            }
            prev = curr;
          }
          Err(error) => {
            println!("Parse error: {error}");
          }
        }
      }
      Err(error) => {
        println!("error: {error}");
        break;
      }
    }
    buffer.clear();
  }
  
  println!("{inc_count} measurements increased");

}

