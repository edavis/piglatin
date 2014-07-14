package piglatin

const (
	pigLatinSuffix string = "ay"
	pigLatinVowelSuffix string = "day"
)

func Translate(in string) string {
	first := in[0]
	switch first {
	case 'a', 'e', 'i', 'o', 'u', 'y':
		return in + pigLatinVowelSuffix
	default:
		return in[1:] + in[0:1] + pigLatinSuffix
	}
}
