package models

type Operation int8

const (
	Exit Operation = iota
	List
	Add
	Complete
	Remove
	Help
)


func (o Operation) String() string {
	return operationNames[o]
}

var operationNames = map[Operation]string{
	Exit:     "Exit",
	List:     "List",
	Add:      "Add",
	Complete: "Complete",
	Remove:   "Remove",
	Help:     "Help",
}
