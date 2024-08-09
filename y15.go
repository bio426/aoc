package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math"
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
	secret := lines[0]
	result := 0

	for result < math.MaxInt32 {
		hashInput := secret + strconv.Itoa(result)
		hash := md5.Sum([]byte(hashInput))
		str := hex.EncodeToString(hash[:])
		if str[:5] == "00000" {
			break
		}
		result += 1
	}

	return int32(result), nil
}

func solutionD4P2(lines []string) (int32, error) {
	secret := lines[0]
	result := 0

	for result < math.MaxInt32 {
		hashInput := secret + strconv.Itoa(result)
		hash := md5.Sum([]byte(hashInput))
		str := hex.EncodeToString(hash[:])
		if str[:6] == "000000" {
			break
		}
		result += 1
	}

	return int32(result), nil
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

func solutionD6P1(lines []string) (int32, error) {
	grid := [1000][1000]bool{}
	for _, line := range lines {
		words := strings.Fields(line)
		command := words[0]
		coordinates := [4]int32{}
		startStrs := []string{}
		endStrs := []string{}
		if command == "toggle" {
			startStrs = strings.Split(words[1], ",")
			endStrs = strings.Split(words[3], ",")
		} else {
			command = strings.Join(words[:2], " ")
			startStrs = strings.Split(words[2], ",")
			endStrs = strings.Split(words[4], ",")
		}
		strCoordinates := append(startStrs, endStrs...)
		for idx, str := range strCoordinates {
			v, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				return 0, err
			}
			coordinates[idx] = int32(v)
		}

		// ejecutar comandos
		for i := coordinates[0]; i <= coordinates[2]; i++ {
			for j := coordinates[1]; j <= coordinates[3]; j++ {
				switch command {
				case "toggle":
					grid[i][j] = !grid[i][j]
				case "turn on":
					grid[i][j] = true
				case "turn off":
					grid[i][j] = false
				}

			}
		}

	}

	littedLights := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				littedLights++
			}
		}
	}

	return int32(littedLights), nil
}

func solutionD6P2(lines []string) (int32, error) {
	grid := [1000][1000]int32{}
	for _, line := range lines {
		words := strings.Fields(line)
		command := words[0]
		coordinates := [4]int32{}
		startStrs := []string{}
		endStrs := []string{}
		if command == "toggle" {
			startStrs = strings.Split(words[1], ",")
			endStrs = strings.Split(words[3], ",")
		} else {
			command = strings.Join(words[:2], " ")
			startStrs = strings.Split(words[2], ",")
			endStrs = strings.Split(words[4], ",")
		}
		strCoordinates := append(startStrs, endStrs...)
		for idx, str := range strCoordinates {
			v, err := strconv.ParseInt(str, 10, 32)
			if err != nil {
				return 0, err
			}
			coordinates[idx] = int32(v)
		}

		// ejecutar comandos
		for i := coordinates[0]; i <= coordinates[2]; i++ {
			for j := coordinates[1]; j <= coordinates[3]; j++ {
				switch command {
				case "toggle":
					grid[i][j] += 2
				case "turn on":
					grid[i][j] += 1
				case "turn off":
					if grid[i][j] > 0 {
						grid[i][j] -= 1
					}
				}

			}
		}

	}

	totalBrightness := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			totalBrightness += int(grid[i][j])
		}
	}

	return int32(totalBrightness), nil
}

func solutionD7P1(lines []string) (int32, error) {
	type record = struct {
		words     []string
		processed bool
	}
	tracking := []*record{}
	wireValues := map[string]uint16{}
	processedCounter := 0

	// cargamos instrucciones
	for _, line := range lines {
		tracking = append(tracking, &record{words: strings.Fields(line)})
	}

	// definimos valores de wires
	for _, r := range tracking {
		if len(r.words) == 3 {
			v, err := strconv.ParseInt(r.words[0], 10, 16)
			if err != nil {
				if errors.Is(err, strconv.ErrSyntax) {
					continue
				} else {
					return 0, err
				}
			}
			wireValues[r.words[len(r.words)-1]] = uint16(v)
			r.processed = true
			processedCounter += 1
		}

	}

	for processedCounter != len(tracking) {

		// recorremos el tracking hasta que todo este procesado
		for _, r := range tracking {
			if r.processed {
				continue
			}
			// set arguments
			operation := r.words[:len(r.words)-2]
			var operator string
			var operants []string
			resultDestiny := r.words[len(r.words)-1]

			if len(operation) == 1 {
				// ASSIGN
				v, ok := wireValues[operation[0]]
				if !ok {
					continue
				}
				wireValues[resultDestiny] = v
			} else if len(operation) == 2 {
				// NOT
				v, ok := wireValues[operation[1]]
				if !ok {
					continue
				}
				wireValues[resultDestiny] = ^v
			} else {
				operator = operation[1]
				operants = []string{operation[0], operation[2]}

				var leftOp uint16
				leftOp, ok := wireValues[operants[0]]
				if !ok {
					v, err := strconv.ParseInt(operants[0], 10, 16)
					if err != nil {
						continue
					}
					leftOp = uint16(v)
				}

				var rightOp uint16
				rightOp, ok = wireValues[operants[1]]
				if !ok {
					v, err := strconv.ParseInt(operants[1], 10, 16)
					if err != nil {
						continue
					}
					rightOp = uint16(v)
				}

				switch operator {
				case "AND":
					wireValues[resultDestiny] = leftOp & rightOp
				case "OR":
					wireValues[resultDestiny] = leftOp | rightOp
				case "RSHIFT":
					wireValues[resultDestiny] = leftOp >> rightOp
				case "LSHIFT":
					wireValues[resultDestiny] = leftOp << rightOp
				}

			}

			r.processed = true
			processedCounter += 1

		}
	}

	return int32(wireValues["a"]), nil
}

func solutionD7P2(lines []string) (int32, error) {
	const previousResponse = 16076

	type record = struct {
		words     []string
		processed bool
	}
	tracking := []*record{}
	wireValues := map[string]uint16{}
	processedCounter := 0

	// cargamos instrucciones
	for _, line := range lines {
		tracking = append(tracking, &record{words: strings.Fields(line)})
	}

	// definimos valores de wires
	for _, r := range tracking {
		if len(r.words) == 3 {
			v, err := strconv.ParseInt(r.words[0], 10, 16)
			if err != nil {
				if errors.Is(err, strconv.ErrSyntax) {
					continue
				} else {
					return 0, err
				}
			}
			wireName := r.words[len(r.words)-1]
			if wireName == "b" {
				wireValues[wireName] = uint16(previousResponse)
			} else {
				wireValues[wireName] = uint16(v)
			}
			r.processed = true
			processedCounter += 1
		}

	}

	for processedCounter != len(tracking) {

		// recorremos el tracking hasta que todo este procesado
		for _, r := range tracking {
			if r.processed {
				continue
			}
			// set arguments
			operation := r.words[:len(r.words)-2]
			var operator string
			var operants []string
			resultDestiny := r.words[len(r.words)-1]

			if len(operation) == 1 {
				// ASSIGN
				v, ok := wireValues[operation[0]]
				if !ok {
					continue
				}
				wireValues[resultDestiny] = v
			} else if len(operation) == 2 {
				// NOT
				v, ok := wireValues[operation[1]]
				if !ok {
					continue
				}
				wireValues[resultDestiny] = ^v
			} else {
				operator = operation[1]
				operants = []string{operation[0], operation[2]}

				var leftOp uint16
				leftOp, ok := wireValues[operants[0]]
				if !ok {
					v, err := strconv.ParseInt(operants[0], 10, 16)
					if err != nil {
						continue
					}
					leftOp = uint16(v)
				}

				var rightOp uint16
				rightOp, ok = wireValues[operants[1]]
				if !ok {
					v, err := strconv.ParseInt(operants[1], 10, 16)
					if err != nil {
						continue
					}
					rightOp = uint16(v)
				}

				switch operator {
				case "AND":
					wireValues[resultDestiny] = leftOp & rightOp
				case "OR":
					wireValues[resultDestiny] = leftOp | rightOp
				case "RSHIFT":
					wireValues[resultDestiny] = leftOp >> rightOp
				case "LSHIFT":
					wireValues[resultDestiny] = leftOp << rightOp
				}

			}

			r.processed = true
			processedCounter += 1

		}
	}

	return int32(wireValues["a"]), nil
}

func solutionD8P1(lines []string) (int32, error) {
	memoryChar := 0
	stringChar := 0

	for _, line := range lines {
		lineMemory := 0

		for idx := 0; idx < len(line); idx++ {
			if idx == 0 || idx == len(line)-1 {
				continue
			}

			if line[idx] == '\\' {
				switch line[idx+1] {
				case '\\', '"':
					idx += 1
				case 'x':
					idx += 3
				}
			}
			lineMemory += 1
		}

		stringChar += len(line)
		memoryChar += lineMemory
	}

	return int32(stringChar - memoryChar), nil
}

func solutionD8P2(lines []string) (int32, error) {
	originalChars := 0
	encodedChars := 0

	for _, line := range lines {
		encoded := ""

		encoded += `"`
		for idx := 0; idx < len(line); idx++ {
			if line[idx] == '\\' || line[idx] == '"' {
				switch line[idx] {
				case '\\':
					encoded += `\\`
				case '"':
					encoded += `\"`
				}
			} else {
				encoded += string(line[idx])
			}
		}
		encoded += `"`

		originalChars += len(line)
		encodedChars += len(encoded)
	}

	return int32(encodedChars - originalChars), nil
}
