package main

import "fmt"

/*
 * Topological Sort using BFS
 */
func findOrder(numCourses int, prerequisites [][]int) (ret []int) {
	graph, degree := make([][]int, numCourses), make([]int, numCourses)

	// initialize graph and degree
	for _, v := range prerequisites {
		src, dest := v[1], v[0]

		graph[src] = append(graph[src], dest)
		degree[dest]++
	}

	var queue []int

	// initialize queue
	for i := 0; i < numCourses; i++ {
		if degree[i] == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		var node int
		node, queue = queue[0], queue[1:]
		ret = append(ret, node)

		for _, v := range graph[node] {
			degree[v]--
			if degree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if len(ret) == numCourses {
		return ret
	}
	return make([]int, 0)
}

// Leetcode #210: Course Schedule II
// https://leetcode.com/problems/course-schedule-ii/
func main() {
	var input = [][]int{
		{1, 0}, {2, 0}, {3, 1}, {3, 2},
	}
	fmt.Printf("%v", findOrder(4, input))
}
