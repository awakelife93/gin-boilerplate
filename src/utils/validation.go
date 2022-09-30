package utils

// * If the Gorm model record is empty, Id is 0.
func IsEmptyRecord(id uint) bool {
	return id == 0
}
