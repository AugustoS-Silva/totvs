arr = [22,7,9,12,8,36,45,92,54,82,63,15,4,3]


def ord(array):
    if len(array) <= 1:
        return array   

    primeiro_valor = array[0]
    menor = [i for i in array[1:] if i <= primeiro_valor]
    maior = [i for i in array[1:] if i > primeiro_valor]
    return ord(menor) + [primeiro_valor] + ord(maior)


print(ord(arr))
