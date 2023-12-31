def alternate_elements(items:list) -> list:
    return items[::2]

# Rotate elements in a list to thhe right(clockwise)
def circular_right_shift(items:list, shift:int) -> list:
    for i in range(0, shift):
        last = items[-1] # store the last element of the list

        items[1:] = items[:-1] # shift elements in the list one step to the right

        items[0] = last     # replace the last element in the beginning   

    return items

# Rotate elements in a list to the left(anti-clockwise)
def circular_left_shift(items:list, shift:int) -> list:
    for i in range(0, shift):
        first = items[0] # store the first element of the list

        items[:-1] = items[1:] # shift elements in the list one step to the left

        items[-1] = first     # replace the first element at the end 

    return items


print(alternate_elements([1,2,3,4,5]))

print(circular_right_shift([1,2,3,4,5], 2))

print(circular_left_shift([1,2,3,4,5], 2))