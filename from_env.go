package yacl

type EnvParams struct {
	Prefix    string
	Delimiter string
}

func FromEnv[T any](params *EnvParams, config *T) {

}
