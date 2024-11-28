package logic

type TestLogic interface {
	Readiness() (int, map[string]string)
	Error() (int, string)
}

type testLogicImpl struct {}

func NewTestLogic() TestLogic {
	return testLogicImpl{}
}

func (_ testLogicImpl) Readiness() (int, map[string]string) {
	return 200, map[string]string{"status": "ok"}
}

func (_ testLogicImpl) Error() (int, string) {
	return 500, "An error occurred"
}