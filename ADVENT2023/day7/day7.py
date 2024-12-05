

hand_eval = 0b000000000000000000000000000000000000000000000000000000000000
ranks = {
    "Five": ''
}

file = open("example.txt", "r").readlines()
hands = []

for game in file:
    hand, bid = game.strip().split(" ")
    hands.append([hand, int(bid.strip())])
# hands = [i.strip().split(" ") for i in open("example.txt", "r").readlines()]
print(hands)