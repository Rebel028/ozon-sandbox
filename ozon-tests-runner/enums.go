package ozon_tests_runner

type InputDataValueType string

var InputDataValueTypes = createRegistry()

const single = "single"
const multi = "multi"
const json = "json"

func createRegistry() *inputDataValueTypeRegistry {
	var integer InputDataValueType = "int"
	var json InputDataValueType = "json"
	var intArray InputDataValueType = "intArray"
	var str InputDataValueType = "string"

	return &inputDataValueTypeRegistry{
		Integer:  &integer,
		Json:     &json,
		IntArray: &intArray,
		String:   &str,
	}
}

type inputDataValueTypeRegistry struct {
	Integer *InputDataValueType
	Json    *InputDataValueType
	// array of integers separated with space
	IntArray *InputDataValueType
	String   *InputDataValueType
}
