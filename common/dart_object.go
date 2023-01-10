package common

type DartObject interface {
	ObjID() string
	ObjName() string
	ObjType() string
}
