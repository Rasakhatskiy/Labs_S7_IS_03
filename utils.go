package main

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func removeDuplicateSpecies(sliceList []Species) []Species {
	allKeys := make(map[string]bool)
	list := []Species{}
	for _, item := range sliceList {
		if _, value := allKeys[item.Gene]; !value {
			allKeys[item.Gene] = true
			list = append(list, item)
		}
	}
	return list
}
