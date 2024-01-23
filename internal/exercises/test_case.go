package exercises

type TestCase interface {
	GetArguments() []interface{}
	GetExpected() interface{}
}

type TestCaseImpl struct {
	Arguments []interface{}
	Expected  interface{}
}

func NewTestCaseImpl(arguments []interface{}, expected interface{}) *TestCaseImpl {
	return &TestCaseImpl{
		Arguments: arguments,
		Expected:  expected,
	}
}

func (t *TestCaseImpl) GetArguments() []interface{} {
	return t.Arguments
}

func (t *TestCaseImpl) GetExpected() interface{} {
	return t.Expected
}
