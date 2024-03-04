str_count = int(input())


def count_taps(value: int) -> int:
    return value // 4 + min(value % 4, 2)


spaces = [count_taps(int(input())) for _ in range(str_count)]

print(sum(spaces))
