package utils

import "github.com/djordjev/scURL/data"

func RemoveByKey(list []data.Pair, key string) []data.Pair {
	for index, elem := range list {
		if elem.Name == key {
			list[index] = list[len(list) - 1]
			return list[:len(list) - 1]
		}
	}

	return list
}
