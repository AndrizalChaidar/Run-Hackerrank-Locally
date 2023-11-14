package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type MapVal struct {
	idx    []uint64
	values []uint64
}

func (map_val *MapVal) Add_MapVal(i uint64, val uint64) {
	(map_val).idx = append(map_val.idx, i)
	(map_val).values = append(map_val.values, val)
}

type DataMap = map[string]*MapVal

func main() {

	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	nTemp, err := strconv.ParseUint(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)

	n := nTemp
	genesTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	healthTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")
	sTemp, err := strconv.ParseUint(strings.TrimSpace(readLine(reader)), 10, 64)
	s := sTemp
	checkError(err)

	data_map := DataMap{}
	data_set := make(map[string]bool)
	for i := uint64(0); i < n; i++ {
		genesItems := genesTemp[i]
		healthItemTemp, err := strconv.ParseUint(healthTemp[i], 10, 64)
		checkError(err)
		healthItem := healthItemTemp
		len_genes := len(genesItems)
		temp_str := ""
		for gene_index, r_gene := range genesItems {
			gene := string(r_gene)
			temp_str += gene
			data_set[temp_str] = true
			if int(gene_index) == len_genes-1 {
				data, ok := (data_map)[genesItems]
				if !ok {
					(data_map)[genesItems] = &MapVal{idx: []uint64{i}, values: []uint64{healthItem}}
				} else {

					data.Add_MapVal(i, healthItem)
				}
			}

		}
	}

	min := uint64(18446744073709551615)
	max := uint64(0)

	for sItr := uint64(0); sItr < s; sItr++ {
		firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

		firstTemp, err := strconv.ParseUint(firstMultipleInput[0], 10, 64)
		checkError(err)
		first := firstTemp

		lastTemp, err := strconv.ParseUint(firstMultipleInput[1], 10, 64)
		checkError(err)
		last := lastTemp

		dna := firstMultipleInput[2]
		res := uint64(0)
		add_res := func(idxs []uint64, values []uint64) {
			for i, idx := range idxs {
				if idx >= first && idx <= last {
					res += values[i]
				}
			}
		}
		for i, rn := range dna {
			char := string(rn)
			_, ok := (data_set)[char]
			if !ok {
				continue
			}
			data, ok := (data_map)[char]
			if ok {
				idx := data.idx
				values := data.values
				add_res(idx, values)
			}
			iSub := i
			dna_len := len(dna)
			for {
				iSub++
				if iSub > dna_len-1 {
					break
				}
				char += string(dna[iSub])
				_, ok := data_set[char]
				if !ok {
					break
				}
				subData, ok := data_map[char]
				if !ok {
					continue
				}
				idx := subData.idx
				values := subData.values
				add_res(idx, values)
				data = subData

			}

		}
		if res > max {
			max = res
		}
		if res < min {
			min = res
		}

	}

	fmt.Println(min, max)
}

func readLine(reader *bufio.Reader) string {
	str_byte, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}
	str := strings.TrimRight(string(str_byte), "\r\n")
	return str
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
