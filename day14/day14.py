"""
O == rounded rock
# == square rock
. == open space

If the place of (row, col) contains a O or a # 
    Check (row, col) around to see if there is an empty space
        If there is an empty space -> move rock
        Otherwise -> rock stays in place

EXAMPLE
    START    | UP ONE
    O..O.#.. | O.OO.##.
    O.OO..#. | O.OO.#.# <-2 rocks stay in place
    ..O..#.# | ........


"""
def part_one():
    file = open('input.txt', 'r')
    x = [list(r.strip()) for r in file.readlines()]
    print(x)

    weight = 0
    for idx, row in enumerate(x):
        for cidx in range(len(row)):
            pos = x[idx][cidx] # get the current position
            if pos == "O":
                x[idx][cidx] = "." # set the index to be empty
                tilt = idx - 1 # tilt the platform
                # while the platform is tilted and the index is empty
                while tilt >= 0 and x[tilt][cidx] == ".":
                    # subtract the position
                    tilt -= 1
                # when the position breaks
                x[tilt+1][cidx] = "O" # change the previous row to a stone
                weight += len(x) - tilt - 1 # add the weight
        print(x[idx])

    print(weight)

def part_two():
    file = open('input.txt', 'r')
    x = [list(r.strip()) for r in file.readlines()]
    d, cycles, states = 0, 0, {}

    while True:
        if d in {0, 2}:
            if d == 2:
                x.reverse()
            else:
                if states is not None:
                    _new = tuple([(ridx, cidx) for ridx, row in enumerate(x) for cidx, col in enumerate(row) if col == "O"])
                    if _new in states:
                        cycles = 1_000_000_000 - (1_000_000_000 - cycles) % (cycles - states[_new])
                        states = None
                    else:
                        states[_new] = cycles
            for idx, row in enumerate(x):
                for cidx in range(len(row)):
                    pos = x[idx][cidx] # get the current position
                    if pos == "O":
                        x[idx][cidx] = "." # set the index to be empty
                        tilt = idx - 1 # tilt the platform
                        # while the platform is tilted and the index is empty
                        while tilt >= 0 and x[tilt][cidx] == ".":
                            # subtract the position
                            tilt -= 1
                        # when the position breaks
                        x[tilt+1][cidx] = "O" # change the previous row to a stone
            
            if d == 2:
                x.reverse()
        elif d == 1:
            for idx, row in enumerate(x[0]):
                for cidx in range(len(row)):
                    pos = x[idx][cidx] # get the current position
                    if pos == "O":
                        x[idx][cidx] = "." # set the index to be empty
                        tilt = cidx - 1
                        while tilt >= 0 and x[idx][tilt] == ".":
                            tilt -= 1
                    # when the position breaks
                    x[idx][tilt+1] = "O" # change the previous row to a stone  
        else:
            for idx in range(len(x[0]) - 1, -1, -1):
                for cidx in range(len(x)):
                    pos = x[idx][cidx] # get the current position
                    if pos == "O":
                        x[idx][cidx] = "." # set the index to be empty
                        tilt = cidx + 1
                        while tilt < len(x[0]) and x[idx][tilt] == ".":
                            tilt += 1
                        x[idx][tilt-1] = "O"
            cycles += 1
            if cycles == 1_000_000_000:
                print(sum([len(x) - ridx for ridx, row in enumerate(x) for cidx, col in enumerate(row) if col == "O"]))
                quit()
        d = (d + 1) % 4


def example():
    file = open('example.txt', 'r')
    x = [list(r.strip()) for r in file.readlines()]
    _new = tuple([(cidx, ridx) for ridx, row in enumerate(x) for cidx, col in enumerate(row) if col == "O"])
    print(x)
    print(_new)
    
example()
part_two()