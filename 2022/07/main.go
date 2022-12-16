package main

import (
	"log"
	"regexp"
	"strconv"

	aoc "github.com/matthewchivers/advent-of-code/utils/go"
)

var (
	root     = directory{}
	lines    = aoc.ReadLinesAsString("input.txt")
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

func partOne() int {
	populateDirectories()
	_ = traverseDirectories()
	size := addUnder100()
	return size
}

func partTwo() int {
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

func addUnder100() int {
	size := 0
	for k := range dirSizes {
		if dirSizes[k] < 100000 {
			size += dirSizes[k]
		}
	}
	return size
}

func traverseDirectories() int {
	dirSizes = make(map[string]int)
	return traverseDirectory(&root, "")
}

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
