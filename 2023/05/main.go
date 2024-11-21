package main

import (
	"fmt"
	"math"
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

// Part One returns the answer to part one of the day's puzzle
func partOne() int {
	chunks := util.ReadFileAsByteChunks("input.txt")
	seeds := getSeedRanges(chunks[0], false)
	lowest := processSeeds(chunks, seeds)
	return lowest
}

// Part Two returns the answer to part two of the day's puzzle
func partTwo() int {
	chunks := util.ReadFileAsByteChunks("input.txt")
	seeds := getSeedRanges(chunks[0], true)
	lowest := processSeeds(chunks, seeds)
	return lowest
}

// getSeedRanges returns a list of segments from the seed chunk
func getSeedRanges(chunk []byte, series bool) SegmentList {
	seeds := SegmentList{}
	seedSpec, err := util.StringToIntArray(strings.SplitN(string(chunk), ": ", 2)[1])
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(seedSpec); i++ {
		// if !series: [20, 5, 10, 4] -> 20, 5, 10, 4
		// if series: [20, 5, 10, 4] -> 20, 21, 22, 23, 24, 10, 11, 12, 13
		seeds = append(seeds, Segment{
			start: seedSpec[i],
			end:   seedSpec[i] + ((seedSpec[i+util.BoolToInt(series)] - 1) * util.BoolToInt(series)),
		})
		if series {
			i++
		}
	}
	return seeds
}

// processSeeds returns the lowest value of the start of all the segments
func processSeeds(chunks [][]byte, seeds SegmentList) int {
	stages := getStages(chunks[1:])

	transformedSeeds := seeds
	for _, stage := range stages {
		transformedSeeds.ApplyStage(stage)
	}
	transformedSeeds.Merge()

	return transformedSeeds.Lowest()
}

func (segList *SegmentList) ApplyStage(stage TransformationStage) {
	seedBuffer := *segList
	workingSegs := SegmentList{}
	for len(seedBuffer) > 0 {
		seed := seedBuffer[0]
		seedBuffer = seedBuffer[1:]
		trnApplied := false
		for _, trnMap := range stage.transformations {
			if trnMap.source.start <= seed.end && trnMap.source.end >= seed.start {
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

func getStages(chunks [][]byte) []TransformationStage {
	stages := []TransformationStage{}
	for i := 0; i < len(chunks); i++ {
		stage := TransformationStage{}
		for _, line := range strings.Split(string(chunks[i]), "\n")[1:] {
			nums, err := util.StringToIntArray(line)
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

// applyTransformation applies the transformation to the seed and returns the transformed seed and any remainders
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

	transformedSeeds = append(transformedSeeds, Segment{
		start: util.MaxInt(transformation.source.start, seg.start) + transformation.alteration,
		end:   util.MinInt(transformation.source.end, seg.end) + transformation.alteration,
	})

	return transformedSeeds, remainder
}

func (segList *SegmentList) Sort() {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(*segList); i++ {
			for j := i + 1; j < len(*segList); j++ {
				if (*segList)[i].start > (*segList)[j].start {
					(*segList)[i], (*segList)[j] = (*segList)[j], (*segList)[i]
					swapped = true
				}
			}
		}
	}
}

func (segList *SegmentList) Merge() {
	segList.Sort()
	merged := SegmentList{}
	// merge any segments that overlap. It is possible for two (or more) segments to overlap
	// with a third segment, so we need to keep merging until there are no more overlaps
	for len(*segList) > 0 {
		seg := (*segList)[0]
		*segList = (*segList)[1:]
		for i := 0; i < len(*segList); i++ {
			// if the segment overlaps with the current segment
			if seg.start <= (*segList)[i].end+1 && seg.end >= (*segList)[i].start-1 {
				// merge the segments
				seg.start = util.MinInt(seg.start, (*segList)[i].start)
				seg.end = util.MaxInt(seg.end, (*segList)[i].end)
				// remove the overlapping segment
				*segList = append((*segList)[:i], (*segList)[i+1:]...)
			}
		}
		merged = append(merged, seg)
	}
	*segList = merged
}

func (segList *SegmentList) Lowest() int {
	lowest := math.MaxInt64
	for _, seg := range *segList {
		if seg.start < lowest {
			lowest = seg.start
		}
	}
	return lowest
}
