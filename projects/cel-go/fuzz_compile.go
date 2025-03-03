package cel

func FuzzCompile(data []byte) int {
	env, err := NewEnv()
	if err != nil {
		panic("impossible to create env")
	}
	ast, issues := env.Compile(string(data))
	if issues != nil && issues.Err() != nil {
		return 0
	}
	_, err = env.Program(ast)
	if err != nil {
		panic("impossible to create prog from ast")
	}

	return 1
}
