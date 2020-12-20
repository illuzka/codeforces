def main(k, data):
    if all([i % k == 0 for i in data]):
        return 0

    data_length = len(data)
    x = 0
    done = False
    done_indexes = set()

    while not done:
        i = 0
        while i < len(data):
            if data[i] % k == 0:
                done_indexes.add(i)
            else:
                if (data[i] + x) % k == 0:
                    data[i] += x
                    done_indexes.add(i)
                    break
            i += 1
        x += 1

        if len(done_indexes) == data_length:
            return x
