package helpers

import "strconv"

func GetLimit(s string) int32 {
	limit, err := strconv.Atoi(s)
	if err != nil {
		limit= 10
	}

	return int32(limit)
}

func GetPage(p string) int32 {
	page, err := strconv.Atoi(p)
	if err != nil {
		page= 1
	}

	return int32(page)
}