package utils

func IsEmptyRecord(id uint) bool {
	const emptyIdValue = 0
	return id == emptyIdValue
}
