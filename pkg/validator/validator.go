package validator

type ValidatorApp interface {
	CheckValidVAT(string, string) (bool, error)
}
