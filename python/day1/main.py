def get_number(line: str):
    for char in line:
        if char.isdigit():
            first_digit = int(char)
            break
    for char in line[::-1]:
        if char.isdigit():
            second_digit = int(char)
            break
    number = first_digit * 10 + second_digit
    return number

    


sum = 0
with open("input.txt") as f:
    for line in f:
        sum += get_number(line)

print(sum)
