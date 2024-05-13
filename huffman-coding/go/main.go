package main

import (
	"fmt"
	"os"
	"slices"
)

type FrequencyTable struct {
	Character rune
	Frequency int
}

type HuffmanNode struct {
	Character rune
	Frequency int
	Code      int
	Left      *HuffmanNode
	Right     *HuffmanNode
}

func main() {
	for i := 1; i <= 6; i++ {
		bookFilename := fmt.Sprintf("book%d.txt", i)
		freqFilename := fmt.Sprintf("book%d.freq", i)
		codeFilename := fmt.Sprintf("book%d.code", i)

		// create frequency table
		frequencyTable := calculate_frequency_table(bookFilename)

		// persist frequency table
		persist_frequency_table(freqFilename, frequencyTable)

		// sort frequency table
		sort_frequency_table(frequencyTable)

		// calculate huffman tree
		rootNode := calculate_huffman_tree(frequencyTable)

		// generate huffman codes
		codes := make(map[rune]string)
		generate_huffman_codes(rootNode, "", codes)

		// persist huffman codes
		persist_huffman_codes(codeFilename, codes)
	}
}

func calculate_frequency_table(filename string) []FrequencyTable {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		fmt.Printf("%s cannot be opened", filename)
		os.Exit(1)
	}

	rawByte := make([]byte, 1)
	freq := make([]FrequencyTable, 0)

	for {
		_, err := file.Read(rawByte)
		if err != nil {
			if err.Error() == "EOF" {
				break
			} else {
				fmt.Println(err.Error())
				os.Exit(1)
			}
		}
		key := rune(rawByte[0])
		idx := slices.IndexFunc(freq, func(f FrequencyTable) bool { return f.Character == key })
		if idx < 0 {
			freq = append(freq, FrequencyTable{
				Character: key,
				Frequency: 1,
			})
		} else {
			freq[idx].Frequency++
		}
	}

	file.Close()

	return freq
}

func persist_frequency_table(filename string, table []FrequencyTable) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("freq.txt cannot be opened to write")
		os.Exit(1)
	}

	for _, f := range table {
		line := fmt.Sprintf("'%c'(%d): %d\r\n", f.Character, f.Character, f.Frequency)
		file.WriteString(line)
	}

	file.Close()
}

func sort_frequency_table(table []FrequencyTable) {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j].Frequency > table[j+1].Frequency {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}

func calculate_huffman_tree(table []FrequencyTable) *HuffmanNode {
	nodes := make([]*HuffmanNode, 0)
	for _, f := range table {
		nodes = append(nodes, &HuffmanNode{
			Character: f.Character,
			Frequency: f.Frequency,
		})
	}

	for len(nodes) > 1 {
		newNode := &HuffmanNode{
			Character: '$',
			Frequency: nodes[0].Frequency + nodes[1].Frequency,
			Left:      nodes[0],
			Right:     nodes[1],
		}

		nodes = nodes[2:]

		nodes = append(nodes, newNode)

		sort_nodes(nodes)
	}

	return nodes[0]
}

func sort_nodes(nodes []*HuffmanNode) {
	for i := 0; i < len(nodes)-1; i++ {
		for j := 0; j < len(nodes)-i-1; j++ {
			if nodes[j].Frequency > nodes[j+1].Frequency {
				nodes[j], nodes[j+1] = nodes[j+1], nodes[j]
			}
		}
	}
}

func generate_huffman_codes(rootNode *HuffmanNode, code string, codes map[rune]string) {
	if rootNode.Left != nil {
		generate_huffman_codes(rootNode.Left, code+"0", codes)
	}

	if rootNode.Right != nil {
		generate_huffman_codes(rootNode.Right, code+"1", codes)
	}

	if rootNode.Left == nil && rootNode.Right == nil {
		codes[rootNode.Character] = code
	}
}

func persist_huffman_codes(filename string, codes map[rune]string) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("freq.txt cannot be opened to write")
		os.Exit(1)
	}

	for char, code := range codes {
		line := fmt.Sprintf("'%c'(%d): %s\r\n", char, char, code)
		file.WriteString(line)
	}

	file.Close()
}
