package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var allDirSizes = make([]int, 100)

func main() {
	file, err := os.Open("day7/input")
	if err != nil {
		log.Fatalf("Error reading file")
	}
	defer file.Close()

	var line string
	scanner := bufio.NewScanner(file)
	scanner.Scan() // skip root directory we will already initialize

	rootDirectory := newDirectory("/", nil)
	currentDirectory := rootDirectory

	for scanner.Scan() {
		line = scanner.Text()
		switch line[:4] {
		case "$ cd":
			if line == "$ cd .." { // GO BACK DIRECTORY
				currentDirectory = currentDirectory.parent
			} else { // SWITCH DIRECTORY
				directoryToSwitchTo := strings.Split(line, " ")[2] // last argument of cd command
				for _, directory := range currentDirectory.directories {
					if directory.name == directoryToSwitchTo {
						currentDirectory = directory
						break
					}
				}
			}
		case "$ ls":
		case "dir ":
			newDirectoryName := line[4:]
			currentDirectory.directories = append(currentDirectory.directories, newDirectory(newDirectoryName, currentDirectory))
		default:
			sizeAndFilename := strings.Split(line, " ")
			size, _ := strconv.Atoi(sizeAndFilename[0])
			filename := sizeAndFilename[1]

			currentDirectory.files = append(currentDirectory.files, newFile(filename, size))
		}
	}
	rootDirSize := rootDirectory.size()
	spaceLeft := 70000000 - rootDirSize
	spaceNeeded := 30000000 - spaceLeft

	closestDirToSpaceNeeded := math.MaxInt64
	totalUnder100kDirectories := 0
	for _, dirSize := range allDirSizes {
		if dirSize <= 100000 {
			totalUnder100kDirectories += dirSize
		} else if dirSize >= spaceNeeded {
			if dirSize-spaceNeeded < closestDirToSpaceNeeded-spaceNeeded {
				closestDirToSpaceNeeded = dirSize
			}
		}
	}

	fmt.Printf("Day 7 - Part 1: %d\n", totalUnder100kDirectories)
	fmt.Printf("Day 7 - Part 2: %d\n", closestDirToSpaceNeeded)
}

type Directory struct {
	name        string
	parent      *Directory
	files       []*File
	directories []*Directory
}

func (dir *Directory) size() (dirSize int) {
	for _, file := range dir.files {
		dirSize += file.size
	}
	for _, subDir := range dir.directories {
		subDirSize := subDir.size()
		allDirSizes = append(allDirSizes, subDirSize)
		dirSize += subDirSize
	}
	return dirSize
}

func newDirectory(name string, parent *Directory) *Directory {
	return &Directory{name: name, parent: parent}
}

type File struct {
	name string
	size int
}

func newFile(name string, size int) *File {
	return &File{
		name: name,
		size: size,
	}
}
