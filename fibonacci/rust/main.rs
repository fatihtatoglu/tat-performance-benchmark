use std::time::{SystemTime, UNIX_EPOCH};

fn fibonacci(n: i32) -> u128 {
    if n == 0 {
        return 0;
    }

    if n == 1 {
        return 1;
    }

    return fibonacci(n - 1) + fibonacci(n - 2);
}

fn main() {
    for i in 1..51 {
        let result = fibonacci(i);

        let start = SystemTime::now();
        let since_the_epoch = start
            .duration_since(UNIX_EPOCH)
            .expect("Time went backwards");

        println!("{}: {} ({:?})", i, result, since_the_epoch);
    }
}
