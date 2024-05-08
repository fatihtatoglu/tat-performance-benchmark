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
    for i in 1..101 {
        let result = fibonacci(i);
        println!("{}: {}", i, result);
    }
}
