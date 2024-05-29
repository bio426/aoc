package main

import (
	"strconv"
	"strings"
)

func solutionD1P1(lines []string) (int32, error) {
	var floor int32
	line := lines[0]
	for _, char := range line {
		if char == '(' {
			floor += 1
		} else {
			floor -= 1
		}
	}
	return floor, nil
}

func solutionD1P2(lines []string) (int32, error) {
	var floor int32
	var pos int32
	line := lines[0]
	for _, char := range line {
		pos += 1
		if char == '(' {
			floor += 1
		} else {
			floor -= 1
		}
		if floor == -1 {
			break
		}
	}
	return pos, nil
}

func solutionD2P1(lines []string) (int32, error) {
	var total int32

	for _, line := range lines {
		var dimensions [3]int32
		strDims := strings.Split(line, "x")
		for idx, str := range strDims {
			val, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				return 0, err
			}
			dimensions[idx] = int32(val)

		}

		s1 := dimensions[0] * dimensions[1]
		s2 := dimensions[1] * dimensions[2]
		s3 := dimensions[2] * dimensions[0]

		boxArea := 2*s1 + 2*s2 + 2*s3
		extra := min(min(s1, s2), s3)

		total += boxArea + extra
	}

	return total, nil
}

func solutionD2P2(lines []string) (int32, error) {
	var total int32

	for _, line := range lines {
		var dimensions [3]int32
		strDims := strings.Split(line, "x")
		for idx, str := range strDims {
			val, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				return 0, err
			}
			dimensions[idx] = int32(val)

		}

		smallestSide := min(min(dimensions[0], dimensions[1]), dimensions[2])
		var smallerSide int32
		if smallestSide == dimensions[0] {
			smallerSide = min(dimensions[1], dimensions[2])
		} else if smallestSide == dimensions[1] {
			smallerSide = min(dimensions[0], dimensions[2])
		} else {
			smallerSide = min(dimensions[0], dimensions[1])
		}

		perimeter := 2*smallestSide + 2*smallerSide
		bow := dimensions[0] * dimensions[1] * dimensions[2]

		total += perimeter + bow
	}

	return total, nil
}

func solutionD3P1(lines []string) (int32, error) {
	history := make(map[[2]int32]bool)
	var posX int32
	var posY int32

	history[[2]int32{posX, posY}] = true
	var visitedHouses int32 = 1

	line := lines[0]

	for _, char := range line {
		if char == '^' {
			posY += 1
		} else if char == '>' {
			posX += 1
		} else if char == 'v' {
			posY -= 1
		} else if char == '<' {
			posX -= 1
		}

		key := [2]int32{posX, posY}
		_, exist := history[key]
		if !exist {
			visitedHouses += 1
		}

		history[key] = true
	}

	return visitedHouses, nil
}

func solutionD3P2(lines []string) (int32, error) {
	historySanta := make(map[[2]int32]bool)
	var santaX int32
	var santaY int32
	historySanta[[2]int32{santaX, santaY}] = true

	historyRobo := make(map[[2]int32]bool)
	var roboX int32
	var roboY int32
	historyRobo[[2]int32{roboX, roboY}] = true

	var visitedHouses int32 = 1

	line := lines[0]

	for idx, char := range line {
		command := [2]int32{}
		if char == '^' {
			command[1] += 1
		} else if char == '>' {
			command[0] += 1
		} else if char == 'v' {
			command[1] -= 1
		} else if char == '<' {
			command[0] -= 1
		}

		isSanta := idx%2 == 0
		var key [2]int32
		if isSanta {
			santaX += command[0]
			santaY += command[1]
			key = [2]int32{santaX, santaY}
		} else {
			roboX += command[0]
			roboY += command[1]
			key = [2]int32{roboX, roboY}
		}

		_, santaExist := historySanta[key]
		_, roboExist := historyRobo[key]
		if !(santaExist || roboExist) {
			visitedHouses += 1
		}

		if isSanta {
			historySanta[key] = true
		} else {
			historyRobo[key] = true
		}

	}

	return visitedHouses, nil
}

func solutionD4P1(lines []string) (int32, error) {
	return 0, nil
}

func solutionD4P2(lines []string) (int32, error) {
	return 0, nil
}

func solutionD5P1(lines []string) (int32, error) {
	var niceStrings int32

	for _, line := range lines {
		var prev rune
		hasProhibited := false
		vowelCount := 0
		hasTwice := false

		for _, char := range line {
			if prev == 'a' && char == 'b' ||
				prev == 'c' && char == 'd' ||
				prev == 'p' && char == 'q' ||
				prev == 'x' && char == 'y' {
				hasProhibited = true
				break
			}

			if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
				vowelCount += 1
			}

			if !hasTwice && (prev == char) {
				hasTwice = true
			}
			prev = char

		}
		if !hasProhibited && vowelCount > 2 && hasTwice {
			niceStrings += 1
		}

	}
	return niceStrings, nil
}

func solutionD5P2(lines []string) (int32, error) {

	var niceStrings int32

	for _, line := range lines {
		hasPairTwice := false
		hasRepeatedBetween := false

		for idx, char := range line {
			// pair check
			if !hasPairTwice && idx < len(line)-2 {
				base := line[idx : idx+2]
				toCompare := line[idx+2:]
				// check if base exist in compare
				// el -1 esta para omitir el ultimo caracter porque no existriria el cIdx+1
				for cIdx := 0; cIdx < len(toCompare)-1; cIdx++ {
					match := toCompare[cIdx] == base[0] && toCompare[cIdx+1] == base[1]
					if match {
						hasPairTwice = true
						break
					}
				}
			}

			// repeat check
			if !hasRepeatedBetween && idx < len(line)-2 {
				next2 := line[idx+2]
				if char == rune(next2) {
					hasRepeatedBetween = true
				}
			}

		}

		if hasPairTwice && hasRepeatedBetween {
			niceStrings += 1
		}

	}
	return niceStrings, nil
}
