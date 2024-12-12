package day9

import (
	"bufio"
	"os"
	"strconv"
)

func GetCorrectDiskChecksumFromFile(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	diskMap := []rune{}

	// input should be one line
	for scanner.Scan() {
		line := scanner.Text()
		diskMap = []rune(line)
	}
	diskBlocks, err := writeDiskBlocksFromDiskMap(diskMap)
	if err != nil {
		return 0, err
	}
	diskBlocks = compressDiskBlocks(diskBlocks)
	checksum := getChecksumFromCompressedDiskBlocks(diskBlocks)
	return checksum, nil

}

func writeDiskBlocksFromDiskMap(diskMap []rune) ([]string, error) {
	diskBlocks := []string{}
	blockIndex := 0
	for index, digit := range []rune(diskMap) {
		digitInt, err := strconv.Atoi(string(digit))
		if err != nil {
			return nil, err
		}
		runeToAppend := "."
		if index%2 == 0 {
			runeToAppend = strconv.Itoa(blockIndex)
			blockIndex++
		}
		for i := 0; i < digitInt; i++ {
			diskBlocks = append(diskBlocks, runeToAppend)
		}
	}
	return diskBlocks, nil
}

func compressDiskBlocks(diskBlocks []string) []string {
	fileRanges := map[int][2]int{}
	freeRanges := []struct{ start, size int }{}

	start := -1
	currFileId := -1
	for index, element := range diskBlocks {
		if element == "." {
			if start == -1 {
				start = index
			}
		} else {
			if aux, _ := strconv.Atoi(element); currFileId != aux {
				if start != -1 {
					freeRanges = append(freeRanges, struct{ start, size int }{start, index - start})
					start = -1
				}
				currFileId, _ = strconv.Atoi(element)
				fileRanges[currFileId] = [2]int{index, index}
			} else {
				fileRanges[currFileId] = [2]int{fileRanges[currFileId][0], index}
			}
		}
	}
	if start != -1 {
		freeRanges = append(freeRanges, struct{ start, size int }{start, len(diskBlocks) - start})
	}

	for id := len(fileRanges) - 1; id >= 0; id-- {
		rangeStart, rangeEnd := fileRanges[id][0], fileRanges[id][1]
		size := rangeEnd - rangeStart + 1

		for i, free := range freeRanges {
			if free.size >= size && free.start < rangeStart {
				for j := 0; j < size; j++ {
					diskBlocks[free.start+j] = strconv.Itoa(id)
					diskBlocks[rangeStart+j] = "."
				}

				freeRanges[i].start += size
				freeRanges[i].size -= size
				if freeRanges[i].size == 0 {
					freeRanges = append(freeRanges[:i], freeRanges[i+1:]...)
				}
				break
			}
		}
	}

	return diskBlocks
}

func getNextFreeSpaceOfNSize(diskBlocks []string, reqSize int) int {
	for i := 0; i <= len(diskBlocks)-reqSize; i++ {
		freeSpace := true
		for j := 0; j < reqSize; j++ {
			if diskBlocks[i+j] != "." {
				freeSpace = false
				break
			}
		}
		if freeSpace {
			return i
		}
	}
	return -1
}

func getChecksumFromCompressedDiskBlocks(diskBlocks []string) int {
	checksum := 0
	for index, block := range diskBlocks {
		if block != "." {
			blockNum, _ := strconv.Atoi(string(block))
			checksum += blockNum * index
		}
	}
	return checksum
}
