package main

import (
	"log"
	"regexp"
	"strconv"

	aoc "github.com/matthewchivers/advent-of-code/utils"
)

var (
	root     = directory{}
	lines    = aoc.ReadFileAsLines("input.txt")
	dirSizes map[string]int
)

type directory struct {
	directories map[string]*directory
	files       map[string]int
	parent      *directory
	name        string
}

func main() {
	log.Println("Part One:", partOne())
	log.Println("Part Two:", partTwo())
}

// calculates the total size of all directories and files whose size is less than 100000 and returns the result.
func partOne() int {
	initialise()
	size := addUnder100()
	return size
}

// calculates the smallest directory that can be deleted to free up 30000000 bytes of space and returns the size of that directory.
func partTwo() int {
	initialise()
	totalSizeOccupied := dirSizes["/"]
	totalSizeAvailable := 70000000 - totalSizeOccupied
	totalSizeRequired := 30000000
	delta := totalSizeRequired - totalSizeAvailable
	smallestDir := ""
	for k := range dirSizes {
		if dirSizes[k] >= delta && (smallestDir == "" || dirSizes[k] < dirSizes[smallestDir]) {
			smallestDir = k
		}
	}
	return dirSizes[smallestDir]
}

// populates the directories and files and stores the size of each directory in a map.
func initialise() {
	if root.directories == nil {
		populateDirectories()
	}
	if dirSizes == nil {
		traverseDirectories()
	}
}

// adds the size of all directories and files whose size is less than 100000 and returns the result.
func addUnder100() int {
	size := 0
	for k := range dirSizes {
		if dirSizes[k] < 100000 {
			size += dirSizes[k]
		}
	}
	return size
}

// helper function that calls the recursive function to traverse the directory tree.
func traverseDirectories() {
	dirSizes = make(map[string]int)
	traverseDirectory(&root, "")
}

// recursively traverses the directory tree and stores the size of each directory in a map.
func traverseDirectory(dir *directory, path string) int {
	dirSizeTotal := 0
	if dir.name != "/" {
		path += dir.name + "/"
	} else {
		path = "/"
	}
	for _, subdir := range dir.directories {
		dirSizeTotal += traverseDirectory(subdir, path)
	}
	for _, fileSize := range dir.files {
		dirSizeTotal += fileSize
	}
	dirSizes[path] = dirSizeTotal
	return dirSizeTotal
}

// populates the directories and files from the input.
func populateDirectories() {
	root.directories = make(map[string]*directory)
	root.files = make(map[string]int)
	currentDir := &directory{}
	for _, line := range lines {
		switch line[0] {
		case '$': // command
			command := line[2:4]
			switch command {
			case "cd":
				arg := line[5:]
				switch arg {
				case "..":
					currentDir = currentDir.parent
				case "/":
					currentDir = &root
					currentDir.name = "/"
				default:
					dir := currentDir.directories[arg]
					currentDir = dir
				}
			}
		case 'd': // directory
			newDirName := line[4:]
			dir := directory{
				directories: make(map[string]*directory),
				files:       make(map[string]int),
				parent:      currentDir,
				name:        newDirName,
			}
			currentDir.directories[newDirName] = &dir
		default: // file
			re := regexp.MustCompile(`(\d+) ([a-z0-9\.]+)`)
			matches := re.FindStringSubmatch(line)
			size, _ := strconv.Atoi(matches[1])
			name := matches[2]
			currentDir.files[name] = size
		}
	}
}
