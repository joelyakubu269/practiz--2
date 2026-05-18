package main

func StackAll(blocks [][]string) []string {
	if len(blocks) == 0 {
		return []string{}
	}
	var res = []string{}
	for i := 0; i < len(blocks); i++ {
		res = StackTwo(res, blocks[i])

	}
	return res
}
