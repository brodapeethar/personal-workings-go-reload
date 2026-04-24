package tree

func Transformer(s string) string {
	content := Hexadecimal(s)
	content = Binary(content)
	content = Capitalize(content)
	content = Punctuate(content)
	content = SingleQuote(content)
	//content = Vowel(content)

	return content
}
