"""
O == rounded rock
# == square rock
. == open space

If the place of (row, item) contains a O or a # 
    Check (row, item) around to see if there is an empty space
        If there is an empty space -> move rock
        Otherwise -> rock stays in place

EXAMPLE
    START    | UP ONE
    O..O.#.. | O.OO.##.
    O.OO..#. | O.OO.#.# <-2 rocks stay in place
    ..O..#.# | ........


"""


def part_one():
    file = open("input.txt", "r")
    beams = [list(r.strip()) for r in file.readlines()]
    print(beams)

    weight = 0
    for idx, beam in enumerate(beams):
        for col in range(len(beam)):
            pos = beams[idx][col]  # get the current position
            if pos == "O":
                beams[idx][col] = "."  # set the index to be empty
                tilt = idx - 1  # tilt the platform
                # while the platform is tilted and the index is empty
                while tilt >= 0 and beams[tilt][col] == ".":
                    # subtract the position
                    tilt -= 1
                # when the position breaks
                beams[tilt + 1][col] = "O"  # change the previous row to a stone
                weight += len(beams) - tilt - 1  # add the weight
        print(beams[idx])

    print(weight)


def part_two():
    file = open("input.txt", "r")
    beams = [list(r.strip()) for r in file.readlines()]
    d, cycles, states = 0, 0, {}

    while True:
        if d in {0, 2}:
            if d == 2:
                beams.reverse()
            else:
                if states is not None:
                    _new = tuple(
                        [
                            (row, col)
                            for row, beam in enumerate(beams)
                            for col, item in enumerate(beam)
                            if item == "O"
                        ]
                    )
                    if _new in states:
                        # print(1000000000 - (1000000000 - cycles) % (cycles - states[_new]))
                        cycles = 1000000000 - (1000000000 - cycles) % (
                            cycles - states[_new]
                        )
                        states = None
                    else:
                        states[_new] = cycles
            for idx, beam in enumerate(beams):
                for col in range(len(beam)):
                    pos = beams[idx][col]  # get the current position
                    if pos == "O":
                        beams[idx][col] = "."  # set the index to be empty
                        tilt = idx - 1  # tilt the platform
                        # while the platform is tilted and the index is empty
                        while tilt >= 0 and beams[tilt][col] == ".":
                            # subtract the position
                            tilt -= 1
                        # when the position breaks
                        beams[tilt + 1][col] = "O"  # change the previous row to a stone

            if d == 2:
                beams.reverse()

        # process the east & west tilting
        elif d == 1:
            for col in range(len(beams[0])):
                for idx in range(len(beams)):
                    pos = beams[idx][col]  # get the current position
                    if pos == "O":
                        beams[idx][col] = "."  # set the index to be empty
                        tilt = col - 1
                        while tilt >= 0 and beams[idx][tilt] == ".":
                            tilt -= 1
                        # when the position breaks
                        beams[idx][tilt + 1] = "O"  # change the previous row to a stone
        else:
            for col in range(len(beams[0]) - 1, -1, -1):
                for idx in range(len(beams)):
                    pos = beams[idx][col]  # get the current position
                    if pos == "O":
                        beams[idx][col] = "."  # set the index to be empty
                        tilt = col + 1
                        while tilt < len(beams[0]) and beams[idx][tilt] == ".":
                            tilt += 1
                        beams[idx][tilt - 1] = "O"
            cycles += 1
            if cycles == 1_000_000_000:
                print(
                    sum(
                        [
                            len(beams) - row
                            for row, beam in enumerate(beams)
                            for col, item in enumerate(beam)
                            if item == "O"
                        ]
                    )
                )
                quit()
        d = (d + 1) % 4
        # print(cycles, end='\r')


def example():
    file = open("example.txt", "r")
    beams = [list(r.strip()) for r in file.readlines()]
    _new = tuple(
        [
            (col, row)
            for row, beam in enumerate(beams)
            for col, item in enumerate(beam)
            if item == "O"
        ]
    )
    print(beams)
    print(_new)


# example()
part_two()
