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
	sInterface, ok := conf.([]string)
	if ok {
		return sInterface
	}

	aInterface, ok := conf.([]interface{})
	if !ok {
		return []string{}
	}

	aString := make([]string, len(aInterface))

	for i, v := range aInterface {
		str, strOk := v.(string)
		if !strOk {
			continue
		}

		aString[i] = str
	}

	return aString
}
