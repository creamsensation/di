package di

type testNested interface {
	getValue() int
}

type testService struct {
	value         int
	nestedService testNested
}

type nestedService struct {
	value int
}

type testPtrService struct {
	value int
}

const (
	testValue       = 1
	testNestedValue = 2
)

func (s *testService) Provide(c Container) *testService {
	return &testService{
		value:         testValue,
		nestedService: Provide[nestedService](c),
	}
}

func (s *nestedService) Provide(c Container) *nestedService {
	return &nestedService{
		value: testNestedValue,
	}
}

func (s *testPtrService) Provide(c Container) *testPtrService {
	return &testPtrService{}
}

func (s *testPtrService) testMethod() bool {
	s.value = testValue
	return true
}

func (s *nestedService) getValue() int {
	return s.value
}
