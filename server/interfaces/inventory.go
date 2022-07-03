package interfaces

type InventoryService interface {
	GetSemesters(username interface{}) []int
}
