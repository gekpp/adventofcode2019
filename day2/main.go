package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	stdin := os.Stdin
	r := bufio.NewReader(stdin)
	// r := bufio.NewReader(strings.NewReader("1,1,1,4,99,5,6,0,99"))
	data := readInput(r)

	d1 := executeProgram(data, 12, 2)
	fmt.Printf("Cell[0] = %d\n", d1[0])

loop:
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			res := executeProgram(data, int64(noun), int64(verb))
			// fmt.Printf("%d, %d => %d\n", noun, verb, res[0])
			if res[0] == 19690720 {
				fmt.Println("100 * noun + verb = ", 100*noun+verb)
				break loop
			}
		}
	}
}

func executeProgram(data []int64, noun, verb int64) []int64 {
	cp := make([]int64, len(data))
	copy(cp, data)
	data = cp
	data[1] = noun
	data[2] = verb
loop:
	for i := 0; i < len(data); i += 4 {
		instr := data[i]
		// fmt.Printf("opcode = %d. State = %d\n", instr, data)
		switch instr {
		case 1:
			data = insert(data, data[i+3], data[data[i+1]]+data[data[i+2]])
		case 2:
			data = insert(data, data[i+3], data[data[i+1]]*data[data[i+2]])
		case 99:
			break loop
		default:
			panic("unknown opcode")
		}
	}
	return data
}

func readInput(r *bufio.Reader) []int64 {
	res := make([]int64, 0)
	for {
		v, ok := readOne(r)
		if !ok {
			return res
		}
		res = append(res, v)
	}
}

func readOne(r *bufio.Reader) (int64, bool) {
	bytes, err := r.ReadString(',')
	if err != nil && err != io.EOF {
		panic(err)
	}

	if len(bytes) == 0 {
		return 0, false
	}

	t, _ := strconv.ParseInt(string(bytes[:len(bytes)-1]), 10, 64)
	return t, true
}

func insert(data []int64, pos, val int64) []int64 {
	var res []int64
	if int(pos) >= len(data) {
		res = make([]int64, pos+1)
		copy(res, data)
	} else {
		res = data
	}
	res[pos] = val
	return res
}
