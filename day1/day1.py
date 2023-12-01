import re

def part_one(input_file):
    nums = []
    with open(input_file, 'r') as file:
        for line in file.readlines():
            m = re.findall('\d', line)
            nums.append(int(''.join([m[0],m[len(m)-1]])))
    print(sum(nums))

def part_two(input_file):
    nums = []
    num_conv = {
        "one": "one1one","two": "two2two",
        "three": "three3three","four": "four4four",
        "five": "five5five", "six": "six6six",
        "seven": "seven7seven","eight": "eight8eight",
        "nine": "nine9nine",
    }
    with open(input_file, 'r') as file:
        for line in file.readlines():
            for key, val in reversed(num_conv.items()):
                line = line.replace(key, val)
            #print(line)
            m = re.findall('\d', line)
            #print(m)
            nums.append(int(''.join([m[0],m[len(m)-1]])))
    print(sum(nums))
            
# part_two('example.txt')
part_two('input_2.txt')