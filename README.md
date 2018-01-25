# ari

This package is for calculating the [Automated Readability Index (ARI)](https://en.wikipedia.org/wiki/Automated_readability_index) of English text.

## usage

```go
package main

import "github.com/karlmcguire/ari"

func main() {
	var score int = ari.ARI("The quick brown fox jumped over the lazy dog.")

	println(score) // "4"
}
```

## reference

| Score | Age | Grade Level |
|:-----:|:---:|:------------|
| 1 | 5-6 | Kindergarten |
| 2 | 6-7 | First Grade |
| 3 | 7-8 | Second Grade |
| 4 | 8-9 | Third Grade |
| 5 | 9-10 | Fourth Grade |
| 6 | 10-11 | Fifth Grade |
| 7 | 11-12 | Sixth Grade |
| 8 | 12-13 | Seventh Grade |
| 9 | 13-14 | Eighth Grade |
| 10 | 14-15 | Ninth Grade |
| 11 | 15-16 | Tenth Grade |
| 12 | 16-17 | Eleventh Grade |
| 13 | 17-18 | Twelfth Grade |
| 14 | 18-22 | College |













