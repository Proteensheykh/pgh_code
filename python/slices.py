def alternate_elements(items:list) -> list:
    return items[::2]

# Rotate elements in a list to thhe right(clockwise)
def circular_shift(items:list, shift:int) -> list:
    for i in range(0, shift):
        last = items[-1] # store the last element of the list

        for j in range(len(items)-1, -1, -1):
            items[j] = items[j-1] # shift elements in the list one step to the right

        items[0] = last     # replace the last element in the beginning   

    return items


print(alternate_elements([1,2,3,4,5]))

print(circular_shift([1,2,3,4,5], 2))