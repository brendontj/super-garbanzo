package sorter

import "sort"

func SortStringToIntMapIntoASlice(in map[string]int) [][]any {
	keyValues := make([][]interface{}, 0, len(in))
	for k, v := range in {
		keyValues = append(keyValues, []interface{}{k, v})
	}
	sort.Slice(keyValues, func(i, j int) bool {
		if keyValues[i][1].(int) == keyValues[j][1].(int) {
			return keyValues[i][0].(string) < keyValues[j][0].(string)
		}
		return keyValues[i][1].(int) > keyValues[j][1].(int)
	})
	return keyValues
}
