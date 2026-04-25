package piscine

func Compact(ptr *[]string) int {
	var result []string

	for _, v := range *ptr {
		if v != "" {
			result = append(result, v)
		}
	}
	*ptr = result
	return len(result)
}
