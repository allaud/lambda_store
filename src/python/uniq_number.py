def uniq_number_generator(count_from):
    def f():
        nonlocal count_from
        count_from = count_from + 1
        return count_from
    return f
uniq_number = uniq_number_generator(0)
