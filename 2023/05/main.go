package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/matthewchivers/advent-of-code/utils"
)

type Segment struct {
	start, end int
}

type Transformation struct {
	source     Segment
	alteration int
}

type TransformationStage struct {
	transformations []Transformation
}

func main() {
	fmt.Println("Hello, advent of code 2023 - Day 5!")
	fmt.Println("Part one:", partOne())
	fmt.Println("Part two:", partTwo())
}

// Part One returns the answer to part one of the day's puzzle
func partOne() int {
	chunks := utils.ReadFileAsByteChunks("input.txt")
	seeds := getIndividualSeeds(chunks)
	lowest := getLowestLocation(chunks, seeds)
	return lowest
}

// Part Two returns the answer to part two of the day's puzzle
func partTwo() int {
	chunks := utils.ReadFileAsByteChunks("input.txt")
	seeds := getSeedRanges(chunks[0])
	lowest := getLowestLocation(chunks, seeds)
	return lowest
}

func getLowestLocation(chunks [][]byte, seeds []Segment) int {
	stages := getStages(chunks[1:])

	transformedSeeds := seeds
	transformedSeeds = applyStages(stages, transformedSeeds)

	transformedSeeds = mergeSegmentLists(transformedSeeds)
	lowest := getLowestSegment(transformedSeeds)
	return lowest
}

func getIndividualSeeds(chunks [][]byte) []Segment {
	seeds := []Segment{}
	seedSpec, err := utils.StringToIntArray(strings.SplitN(string(chunks[0]), ": ", 2)[1])
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(seedSpec); i++ {
		seg := Segment{
			start: seedSpec[i],
			end:   seedSpec[i],
		}
		seeds = append(seeds, seg)
	}
	return seeds
}

func getLowestSegment(transformedSeeds []Segment) int {
	lowest := math.MaxInt64
	for _, seed := range transformedSeeds {
		if seed.start < lowest {
			lowest = seed.start
		}
	}
	return lowest
}

func applyStages(stages []TransformationStage, transformedSeeds []Segment) []Segment {
	for _, stage := range stages {

		seedBuffer := mergeSegmentLists(transformedSeeds)
		transformedSeeds = []Segment{}
		for len(seedBuffer) > 0 {
			seed := seedBuffer[0]
			seedBuffer = seedBuffer[1:]
			trnApplied := false
			for _, trnMap := range stage.transformations {
				if trnMap.source.start <= seed.end && trnMap.source.end >= seed.start {
					transformedSeg, leftOvers := applyTransformation(trnMap, seed)
					seedBuffer = append(seedBuffer, leftOvers...)
					transformedSeeds = append(transformedSeeds, transformedSeg...)
					trnApplied = true
					break
				}

			}
			if !trnApplied {
				transformedSeeds = append(transformedSeeds, seed)
			}
		}
	}
	return transformedSeeds
}

func getStages(chunks [][]byte) []TransformationStage {
	stages := []TransformationStage{}
	for i := 0; i < len(chunks); i++ {
		stage := TransformationStage{}
		for _, line := range strings.Split(string(chunks[i]), "\n")[1:] {
			nums, err := utils.StringToIntArray(line)
			if err != nil {
				panic(err)
			}
			stage.transformations = append(stage.transformations, Transformation{

				source: Segment{
					start: nums[1],
					end:   nums[1] + nums[2] - 1,
				},
				alteration: nums[0] - nums[1],
			})
		}
		stages = append(stages, stage)
	}
	return stages
}

func applyTransformation(transformation Transformation, seed Segment) ([]Segment, []Segment) {
	workingSeeds := []Segment{}
	transformedSeeds := []Segment{}
	if transformation.source.start > seed.start {
		preSeed := Segment{
			start: seed.start,
			end:   transformation.source.start - 1,
		}
		workingSeeds = append(workingSeeds, preSeed)
	}
	transformedSeed := Segment{
		start: utils.MaxInt(transformation.source.start, seed.start) + transformation.alteration,
		end:   utils.MinInt(transformation.source.end, seed.end) + transformation.alteration,
	}
	transformedSeeds = append(transformedSeeds, transformedSeed)

	if transformation.source.end < seed.end {
		postSeed := Segment{
			start: transformation.source.end + 1,
			end:   seed.end,
		}
		workingSeeds = append(workingSeeds, postSeed)
	}
	return transformedSeeds, workingSeeds
}

func getSeedRanges(chunks []byte) []Segment {
	seeds := []Segment{}
	seedSpec, err := utils.StringToIntArray(strings.SplitN(string(chunks), ": ", 2)[1])
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(seedSpec); i += 2 {
		seg := Segment{
			start: seedSpec[i],
			end:   seedSpec[i] + seedSpec[i+1] - 1,
		}
		seeds = append(seeds, seg)
	}
	return seeds
}

func mergeSegmentLists(segs []Segment) []Segment {
	workingSegs := segs
	workingSegs = sortSegments(workingSegs)
	merged := []Segment{}
	// merge any segments that overlap. It is possible for two (or more) segments to overlap
	// with a third segment, so we need to keep merging until there are no more overlaps
	for len(workingSegs) > 0 {
		seg := workingSegs[0]
		workingSegs = workingSegs[1:]
		for i := 0; i < len(workingSegs); i++ {
			// if the segment overlaps with the current segment
			if seg.start <= workingSegs[i].end+1 && seg.end >= workingSegs[i].start-1 {
				// merge the segments
				seg.start = utils.MinInt(seg.start, workingSegs[i].start)
				seg.end = utils.MaxInt(seg.end, workingSegs[i].end)
				// remove the overlapping segment
				workingSegs = append(workingSegs[:i], workingSegs[i+1:]...)
			}
		}
		merged = append(merged, seg)
	}

	return merged
}

func sortSegments(segs []Segment) []Segment {
	sorted := segs
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(sorted); i++ {
			for j := i + 1; j < len(sorted); j++ {
				if sorted[i].start > sorted[j].start {
					sorted[i], sorted[j] = sorted[j], sorted[i]
					swapped = true
				}
			}
		}
	}
	return sorted
}
