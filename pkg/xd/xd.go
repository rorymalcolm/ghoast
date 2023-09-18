package xd

//chained example struct that is exported

type Chained struct {
	//exported field
	ExportedField string
	//unexported field
}

func New() Chained {
	return Chained{
		ExportedField: "Hello, World!",
	}
}

func (c Chained) ThisIsExported() Chained {
	return Chained{
		ExportedField: "Hello, World!",
	}
}
