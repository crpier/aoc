fn read_input() -> String {
    std::fs::read_to_string("input.txt").unwrap()
}

fn get_first_digit(text: &str) -> Option<char> {
    for char in text.chars() {
        if char.is_numeric() {
            return Some(char.clone());
        }
    }
    return None;
}

fn get_last_digit(text: &str) -> Option<char> {
    for char in text.chars().rev() {
        if char.is_numeric() {
            return Some(char.clone());
        }
    }
    return None;
}

fn main() {
    let mut sum = 0;

    let input = read_input();
    for line in input.lines() {
        let first_digit = get_first_digit(line).unwrap();
        let last_digit = get_last_digit(line).unwrap();
        let new_number: i32 = format!("{}{}", first_digit, last_digit).parse::<i32>().unwrap();
        println!("{new_number}");
        sum += new_number;
    }
    println!("{sum}");
}
