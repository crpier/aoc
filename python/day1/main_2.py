import contextlib


def get_last_5_chars(line: str, idx: int) -> str:
    if idx < 1:
        return ""
    if idx < 5:
        start = 0
    else:
        start = idx - 5
    chars = line[start : idx + 1]
    chars = chars.rjust(5)
    return chars


def get_next_digit(line: str) -> tuple[int, str]:
    for idx, char in enumerate(line):
        if char.isdigit():
            return int(char), line[idx + 1 :]
        else:
            last_chars = get_last_5_chars(line, idx)
            if "one" in last_chars:
                return 1, line[idx + 1 :]
            elif "two" in last_chars:
                return 2, line[idx + 1 :]
            elif "three" in last_chars:
                return 3, line[idx + 1 :]
            elif "four" in last_chars:
                return 4, line[idx + 1 :]
            elif "five" in last_chars:
                return 5, line[idx + 1 :]
            elif "six" in last_chars:
                return 6, line[idx + 1 :]
            elif "seven" in last_chars:
                return 7, line[idx + 1 :]
            elif "eight" in last_chars:
                return 8, line[idx + 1 :]
            elif "nine" in last_chars:
                return 9, line[idx + 1 :]

    raise ValueError("No numbers in line")


def get_number(line: str) -> int:
    original_line = line
    first_number, line = get_next_digit(line)
    numbers = [first_number]
    while True:
        try:
            new_number, line = get_next_digit(line)
            numbers.append(new_number)
        except ValueError:
            break

    if len(numbers) < 2:
        print(f"line {original_line} gave result {numbers}")
        print(original_line[::-1])

    number = numbers[0] * 10 + numbers[-1]
    print(f"{original_line} --> {numbers} => {number}")
    return number


sum = 0
# with open("example_2.txt") as f:
with open("input.txt") as f:
    for line in f:
        line = line.strip()
        sum += get_number(line)
print(sum)
