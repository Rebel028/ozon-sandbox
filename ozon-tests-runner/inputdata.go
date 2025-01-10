package ozon_tests_runner

type TestMetadata struct {
	Properties []InputDataValueInfo
}

type InputDataValueInfo struct {
	// starting line number for value
	LineNumber int
	// property name
	Name string
	// how many lines hold the data
	RangeLines int
	// data type
	DataType InputDataValueType
}
