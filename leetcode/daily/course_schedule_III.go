package daily

/*
	There are n different online courses numbered from 1 to n. You are given an array courses where courses[i] = [durationi, lastDayi] indicate that the ith course should be taken continuously for durationi days and must be finished before or on lastDayi.

	You will start on the 1st day and you cannot take two or more courses simultaneously.

	Return the maximum number of courses that you can take.
 */

func scheduleCourse(courses [][]int) int {
	var count, duration int
	maxLimitCourse := []int{0, 0}
	for _, slice := range courses {
		if slice[0] < slice[1] {
			if slice[1] > duration + slice[0] {
				duration += slice[0]
				count++
				if maxLimitCourse[1] > slice[1] {
					maxLimitCourse = slice
				}
			} else {
				duration -= maxLimitCourse[0]
				count--
				//tmp := maxLimitCourse

				duration += slice[0]
				count++
			}
		}
	}

	return count
}
