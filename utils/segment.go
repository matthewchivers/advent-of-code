package utils

import "math"

type Segment struct {
	Start, End int
}

type SegmentList []Segment

func (segList *SegmentList) Sort() {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(*segList); i++ {
			for j := i + 1; j < len(*segList); j++ {
				if (*segList)[i].Start > (*segList)[j].Start {
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
			if seg.Start <= (*segList)[i].End+1 && seg.End >= (*segList)[i].Start-1 {
				// merge the segments
				seg.Start = MinInt(seg.Start, (*segList)[i].Start)
				seg.End = MaxInt(seg.End, (*segList)[i].End)
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
		if seg.Start < lowest {
			lowest = seg.Start
		}
	}
	return lowest
}
