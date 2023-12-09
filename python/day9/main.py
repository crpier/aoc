def list_is_zero(values: list[int]):
    return all(value == 0 for value in values)


def list_is_not_zero(values: list[int]):
    return not list_is_zero(values)


def get_derivates(values: list[int]):
    derivates_list: list[int] = []
    for idx in range(len(values) - 1):
        derivates_list.append(values[idx + 1] - values[idx])
    return derivates_list


def calculate_next_value(values: list[int]):
    lists: list[list[int]] = [values]
    while list_is_not_zero(lists[-1]):
        lists.append(get_derivates(lists[-1]))
    for lists_idx in reversed(range(len(lists))):
        if list_is_zero(lists[lists_idx]):
            lists[lists_idx].append(0)
        else:
            lists[lists_idx].append(lists[lists_idx][-1] + lists[lists_idx + 1][-1])
    return lists[0][-1]


def calculate_previous_value(values: list[int]):
    lists: list[list[int]] = [values]
    while list_is_not_zero(lists[-1]):
        lists.append(get_derivates(lists[-1]))
    for lists_idx in reversed(range(len(lists))):
        if list_is_zero(lists[lists_idx]):
            lists[lists_idx].insert(0, 0)
        else:
            lists[lists_idx].insert(0, lists[lists_idx][0] - lists[lists_idx + 1][0])
    return lists[0][0]


# next_values_sum = 0
prev_values_sum = 0
with open("input.txt") as f:
# with open("example.txt") as f:
    for line in f:
        values: list[int] = [int(val) for val in line.strip().split()]
        # next_value = calculate_next_value(values)
        prev_value = calculate_previous_value(values)
        # print(values, next_value)
        # next_values_sum += next_value
        prev_values_sum += prev_value
print(prev_values_sum)
