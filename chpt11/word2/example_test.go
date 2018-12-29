package word

import "fmt"

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("text"))
	// Output:
	// true
	// false
}
