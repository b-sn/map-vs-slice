#!/usr/bin/env python3

with open('./1984.txt', 'r') as file:
    text = file.read()

def count_characters_dict(text):
    characters = {}
    for character in text:
        ch = ord(character)
        if ch not in characters:
            characters[ch] = 0
        characters[ch] += 1
    return characters

def count_characters_array(text):
    characters = [0] * 123
    for character in text:
        ch = ord(character)
        characters[ch] += 1
    return characters


import timeit
print('count_characters_dict')
t = timeit.Timer('count_characters_dict(text)', 'from __main__ import count_characters_dict, text')
print(t.timeit(number=1000))
print('count_characters_array')
t = timeit.Timer('count_characters_array(text)', 'from __main__ import count_characters_array, text')
print(t.timeit(number=1000))

