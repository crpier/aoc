fn read_input() -> String {
    std::fs::read_to_string("src/input.txt").unwrap();
}

fn main() {
    let input = read_input();
    println!("{input}");
}
