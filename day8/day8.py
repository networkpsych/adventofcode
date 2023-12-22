def part_one(instruction: str, parts: dict):
    steps = 0
    node = {list(parts.keys())[0] : parts[list(parts.keys())[0]]}
    curr_key = list(parts.keys())[0]
    prev_key = ''
    while curr_key != "ZZZ" and steps < 20000000:
        print(steps, end="\r")
        if curr_key == prev_key:
            break
        for step in instruction:
            # print(node)
            # print(step, end="\r")
            steps += 1
            for key, val in node.items():
                prev_key = key
                if key == "ZZZ":
                    print(steps)
                    break
                elif step == "R":
                    node = {val[1]: parts[val[1]]}
                    curr_key = val[1]
                elif step == "L":
                    node = {val[0]: parts[val[0]]}
                    curr_key = val[0] 
    print(node)
    return steps

def part_two(instructions: str, parts: dict):
    # need to make this work lol
    steps = 0
    start = 'AAA'
    node = {list(parts.keys())[0] : parts[list(parts.keys())[0]]}
    curr_key = list(parts.keys())[0]
    prev_key = ''

    while not curr_key.endswith('Z'):
        steps += 1
        for i in instructions:
            for key, val in node.items():
                prev_key = key
                if i == "R":
                    node = {val[1]: parts[val[1]]}
                    curr_key = val[1]
                elif i == "L":
                    node = {val[0]: parts[val[0]]}
                    curr_key = val[0] 
    print(node)
    return steps
        

p1 = {}
with open('input.txt', 'r') as file:
    p = [i.strip().split(" = ") for i in file]
    for i in p:
        i[1] = i[1].replace("(", "")
        i[1] = i[1].replace(")", "")
        p1[i[0]] = i[1].split(", ")

with open("instructions.txt", 'r') as inst:
    pi = inst.read()


example = {
    "AAA": ("BBB", "CCC"),
    "BBB": ("DDD", "EEE"),
    "CCC": ("ZZZ", "GGG"),
    "DDD": ("DDD", "DDD"),
    "EEE": ("EEE", "EEE"),
    "GGG": ("GGG", "GGG"),
    "ZZZ": ("ZZZ", "ZZZ")
    }
tmp = {}
# print(p1)
#instructions = "RL"
#node = {list(example.keys())[0] : example[list(example.keys())[0]]}
# print(node)
# print(type(node))
# print(node.items())
for item in sorted(p1):
    tmp[item] = p1[item]
print(part_two(pi, tmp))
