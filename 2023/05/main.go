package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/matthewchivers/advent-of-code/util"
)

type Segment struct{ start, end int }

type SegmentList []Segment

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

// Part One reads the initial seed numbers from the input file and converts them through a sequence
// of transformations (seed-to-soil, soil-to-fertiliser, etc.) based on the given almanac rules.
// It finds the lowest final location number after all transformations are applied.
func partOne() int {
	chunks := util.ReadFileAsByteChunks("input.txt")
	seeds := getSeedRanges(chunks[0], false)
	lowest := processSeeds(chunks, seeds)
	return lowest
}

// In Part Two, the input represents ranges of seeds rather than individual numbers.
// Transforms these ranges them through similar transformations to Part One, and then
// finds the lowest resulting location number.
func partTwo() int {
	chunks := util.ReadFileAsByteChunks("input.txt")
	seeds := getSeedRanges(chunks[0], true)
	lowest := processSeeds(chunks, seeds)
	return lowest
}

// getSeedRanges returns a list of segments from the seed chunk
// Depending on the `series` flag, either generates a direct mapping of the seed values
// or generates a contiguous range from each seed value.
func getSeedRanges(chunk []byte, series bool) SegmentList {
	seeds := SegmentList{}
	seedSpec, err := util.StringToIntArray(strings.SplitN(string(chunk), ": ", 2)[1])
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(seedSpec); i++ {
		// If `series` is false, treat each value as an individual seed.
		// If `series` is true, treat each value pair as defining a range: the first value is the start,
		// and the second value is the number of values in the range.
		start := seedSpec[i]
		end := start
		if series && i+1 < len(seedSpec) {
			rangeLength := seedSpec[i+1]
			end = start + rangeLength - 1
			i++ // Skip the next value since it's part of the range definition
		}
		seeds = append(seeds, Segment{start: start, end: end})
	}
	return seeds
}

// processSeeds takes the initial seed segments and applies multiple transformation stages to them.
// It returns the lowest value from the transformed segments.
// Each transformation stage is parsed and applied to propagate the segments through each mapping level.
func processSeeds(chunks [][]byte, seeds SegmentList) int {
	stages := getStages(chunks[1:])

	transformedSeeds := seeds
	for _, stage := range stages {
		transformedSeeds.ApplyStage(stage)
	}
	transformedSeeds.Merge()

	return transformedSeeds.Lowest()
}

// ApplyStage applies a series of transformations from a TransformationStage to the SegmentList.
// Each transformation modifies specific segments that overlap with its source range.
func (segList *SegmentList) ApplyStage(stage TransformationStage) {
	seedBuffer := *segList
	workingSegs := SegmentList{}
	for len(seedBuffer) > 0 {
		seed := seedBuffer[0]
		seedBuffer = seedBuffer[1:]
		trnApplied := false
		for _, trnMap := range stage.transformations {
			if trnMap.source.start <= seed.end && trnMap.source.end >= seed.start {
				// Apply the transformation if the seed overlaps with the transformation source range
				transformedSeg, leftOvers := seed.applyTransformation(trnMap)
				seedBuffer = append(seedBuffer, leftOvers...)
				workingSegs = append(workingSegs, transformedSeg...)
				trnApplied = true
				break
			}

		}
		if !trnApplied {
			workingSegs = append(workingSegs, seed)
		}
	}
	*segList = workingSegs
}

// getStages parses the transformation chunks and creates a list of TransformationStage instances
// This optimized version uses a more readable and efficient loop to parse each line of transformation.
func getStages(chunks [][]byte) []TransformationStage {
	stages := make([]TransformationStage, len(chunks))
	for i, chunk := range chunks {
		lines := strings.Split(string(chunk), "\n")[1:]
		transformations := make([]Transformation, 0, len(lines))
		for _, line := range lines {
			if line == "" {
				continue // Skip empty lines
			}
			nums, err := util.StringToIntArray(line)
			if err != nil {
				panic(err)
			}
			transformations = append(transformations, Transformation{
				source: Segment{
					start: nums[1],
					end:   nums[1] + nums[2] - 1,
				},
				alteration: nums[0] - nums[1],
			})
		}
		stages[i] = TransformationStage{transformations: transformations}
	}
	return stages
}

// applyTransformation applies the transformation to the seed and returns the transformed seed and any remainders.
// It splits the segment into parts that remain unchanged and the part that undergoes transformation.
func (seg *Segment) applyTransformation(transformation Transformation) (SegmentList, SegmentList) {
	remainder, transformedSeeds := SegmentList{}, SegmentList{}

	if transformation.source.start > seg.start {
		// everything before the transformed segment
		remainder = append(remainder, Segment{
			start: seg.start,
			end:   transformation.source.start - 1,
		})
	}
	if transformation.source.end < seg.end {
		// everything after the transformed segment
		remainder = append(remainder, Segment{
			start: transformation.source.end + 1,
			end:   seg.end,
		})
	}

	// The transformed part of the segment
	transformedSeeds = append(transformedSeeds, Segment{
		start: util.MaxInt(transformation.source.start, seg.start) + transformation.alteration,
		end:   util.MinInt(transformation.source.end, seg.end) + transformation.alteration,
	})

	return transformedSeeds, remainder
}

// Sort sorts the SegmentList in ascending order based on the start value of each segment.
func (segList *SegmentList) Sort() {
	sort.Slice(*segList, func(i, j int) bool {
		return (*segList)[i].start < (*segList)[j].start
	})
}

// Merge combines overlapping or adjacent segments in the SegmentList into a single segment.
// This function has been optimized to improve efficiency.
func (segList *SegmentList) Merge() {
	segList.Sort()
	merged := SegmentList{}

	for _, seg := range *segList {
		if len(merged) == 0 {
			merged = append(merged, seg)
			continue
		}

		last := &merged[len(merged)-1]
		// If the current segment overlaps or is adjacent to the last segment in merged, merge them.
		if seg.start <= last.end+1 {
			last.end = util.MaxInt(last.end, seg.end)
		} else {
			merged = append(merged, seg)
		}
	}

	*segList = merged
}

// Lowest returns the lowest start value from all the segments in the SegmentList.
func (segList *SegmentList) Lowest() int {
	lowest := math.MaxInt64
	for _, seg := range *segList {
		if seg.start < lowest {
			lowest = seg.start
		}
	}
	return lowest
}
