def get_x(data, k):
    """

    :param data:
    :param k:
    :return:
    """

    def list_of_add_values(k, element):
        """

        :return:
        """

        add_ = lambda k, x: k - (x % k)
        first_value = add_(k, element)
        list_ = [first_value, ]

        i = 0
        z = 0
        while i < 50:
            z = list_[i] + k
            list_.append(z)
            i += 1

        return list_

    add_ = lambda i, k: i % k == 0
    found = all([add_(i, k) for i in data])
    loop = 0
    x = 0

    while not found:
        i = 0
        while i < len(data):
            values_to_add = list_of_add_values(k, data[i])
            if k not in values_to_add:
                if x in values_to_add:
                    data[i] += x
                    loop = 0
                    break
            else:
                loop += 1
                if loop >= len(data):
                    if (data[i] + x) % k == 0:
                        data[i] += x
                        loop = 0
                        break
            i += 1
        x += 1
        found = all([add_(i, k) for i in data])


    return x




x = int(input())
data = []

for i in range(x):
    temp_data = []
    for j in range(2):
        d = [int(i) for i in input().split(" ")]
        temp_data.append(d)
    data.append(temp_data)



for i in data:
    print(get_x(i[1],i[0][1]))
