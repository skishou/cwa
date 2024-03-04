p, v = map(int, input().split())
q, m = map(int, input().split())

all_variants = 2 * (v + m + 1)
min_high_tree = min(p + v, q + m)
max_low_tree = max(p - v, q - m)

result = all_variants - max(0, min_high_tree - max_low_tree + 1)
print(result)
