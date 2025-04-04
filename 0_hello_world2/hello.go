package helloworld2

const helloConstant = "Hello, "

func Hello(name string) string {
	if name == "" {
		return helloConstant + "World!"
	}
	return helloConstant + name
}