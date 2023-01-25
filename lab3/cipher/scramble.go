package cipher

import (
	"math/rand"
	"strconv"
)

/*
Task 4: Scrambling text

In this task the objective is to transform a given text such that all letters of each word
are randomly shuffled except for the first and last letter.

For example, given the word "scramble", the result could be "srmacble" or "sbcamrle",
or any other permutation as long as the first and last letters stay the same.

An entire sentence scrambled like this should still be readable:
"it deosn't mttaer in waht oredr the ltteers in a wrod are,
the olny iprmoetnt tihng is taht the frist and lsat ltteer be at the rghit pclae"
See https://www.mrc-cbu.cam.ac.uk/people/matt.davis/cmabridge/ for more
information and examples.

Implementation:
The task is to implement the scramble function, which takes a text in the form of a string and a seed.
A seed is given so the output from your solution should match the test cases if it is correct.
The seed should be applied at the start of the function.
Remember that the implementation should keep any punctuation and spacing intact, and all numbers should be untouched.

Shuffling the letters and applying the seed can be done using the math/rand package (https://golang.org/pkg/math/rand/).
Use the Shuffle function to ensure you reach the same values as given in the tests (scramble_test.go).
*/

func scramble(text string, seed int64) string {
	rand.Seed(seed)
	slicy := tokenize(text)
	var str string = ""
	for a := 0; a < len(slicy); a++ {

		_, err := strconv.Atoi(slicy[a])
		if err != nil {

			var shufle []byte
			for i := 1; i < len(slicy[a])-1; i++ {
				shufle = append(shufle, slicy[a][i])
			}

			if slicy[a] != " " && slicy[a] != "." && slicy[a] != "," && slicy[a] != "(" && slicy[a] != ")" {
				new_length := 0
				for d := 0; d < len(shufle); d++ {
					if string(shufle[d]) == "'" {
						new_length = d - 1
					}
				}
				if new_length != 0 {
					for l := 0; l < new_length-1; l++ {
						m := l + 1
						shufle[l], shufle[m] = shufle[m], shufle[l]
					}
				} else {
					rand.Shuffle(len(shufle), func(k, j int) {

						shufle[k], shufle[j] = shufle[j], shufle[k]
					})
				}

				var scrambled_list []byte
				scrambled_list = append(scrambled_list, slicy[a][0])
				for d := 0; d < len(shufle); d++ {
					scrambled_list = append(scrambled_list, shufle[d])
				}
				scrambled_list = append(scrambled_list, slicy[a][len(slicy[a])-1])

				str = str + string(scrambled_list)
			} else {
				switch slicy[a] {
				case " ":
					str += " "
				case ".":
					str += "."
				case "(":
					str += "("
				case ")":
					str += ")"
				case ",":
					str += ","
				default:

				}
			}
		} else {
			return text
		}
	}

	return str
}
