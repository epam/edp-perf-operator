package common

import "sort"

func GetStringP(val string) *string {
	return &val
}

func SortArray(array []string) {
	sort.Slice(array, func(i, j int) bool {
		return array[i] < array[j]
	})
}

func ConvertToStringArray(conf interface{}) []string {
	aInterface := conf.([]interface{})
	aString := make([]string, len(aInterface))
	for i, v := range aInterface {
		aString[i] = v.(string)
	}
	return aString
}
