fn main() {
    // let world: &'static str = "prithvi";
    let num = 238479238;
    println!("Hello, {}! -Stranger{}", foo("bar"), num);
}

fn foo(_x: &'static str) -> &'static str {
    _x
}
