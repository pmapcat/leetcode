// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 14:53 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@
package main

import (
	"log"
)

// Problem: https://leetcode.com/problems/concatenated-words/

// Input: ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]

// check("cats") =>
//   getPrefixes("cats") => [cat,cats]
//     getPrefixes("cat") => [cat]
//     getPrefixes("cats") => [cat,cats]

// Compound words are always longer than simple words that they are part of
// sorted(x)
// Solution(x) returns ([simple words of x], [compound words of x])
// Solution(x) = Solution(x-1) + if(x in simple words) => do nothing
//                                                     => check against simple words,
//                                                          if match, then to [concatenaded words]
//                                                          otherwise      to [simple words]
// No point in doing that, although the idea that compound words are always longer wants to be exploited

//
//  Tricky usecase for greedy check: ["catsdog","cat","sdog","dog"], if match "dog", then the "catsdog" is not going to be matched
// [catsdogscat,cats,cat,dog]
// getAllWordsThatStartWith(cats) => [cats,catsdogscat]
// check(catsdog) => [catsdog,cat]

// (cats, catsdogscat)
//["catsdogscat", "cats", "cat", "dog","catsdogcat"]
// check(catsdogcat) = [cats] + check(dogcat)
//   check(dogcat)  = [dog] + check(cat)
//   check(cat)     = true

// check(cats) = (cat + check(s)) == false || (cats + check())

// check(catsdogscat) = [cats] + check(dogscat)
//   check(dogscat)  =  [dog] + check(scat)
//   check(scat)     =  false

// check("cats") = false
// check("cat")  = false
// check("dog")  = false
// check("catsdogscat") = cat + check(sdogscat) || cats + check(dogscat) || catsdogscat + check()
// check(sdogscat) = []
// check(cats)     = []
// check(catsdogcat) = [cats + dogcat + ca]tg

// solution(X,isFirstLevel?)
//   noPrefix && isFirstLevel  => fatal("Cannot happen")
//   noPrefix && notFirstLevel => return false
//   for prefix in X:
//     isFirstLevel:
//       prefix is empty => fatal("cannot happen")
//       prefix != X  =>
//         if solution(X-prefix,false) =>  return true
//     notFirstLevel:
//       prefix is empty => return true
//       prefix == X => return true (match happened)
//       prefix != X => if solution(X-prefix,false) =>  return true
//   return false

// words that start with X: for each
func wordsThatStartWith(input string, words map[string]bool) []string {
	result := []string{}
	_, ok := words[input]
	if ok {
		result = append(result, input)
	}
	for i := 0; i < len(input); i++ {
		_, ok := words[input[:i]]
		if ok {
			result = append(result, input[:i])
		}
	}
	return result
}

func findAllConcatenatedWordsInADict(words []string) []string {
	// store words in an array
	wordStore := map[string]bool{}
	for _, word := range words {
		wordStore[word] = true
	}

	// checkFnMemo := map[string]bool{}
	var checkFn func(string) bool
	checkFn = func(word string) bool {
		// checked, ok := checkFnMemo[word]
		// if ok {
		// 	return checked
		// }

		if len(word) == 0 {
			return true
		}
		prefixes := wordsThatStartWith(word, wordStore)
		if len(prefixes) == 1 {
			return true
		}
		if len(prefixes) == 0 {
			return false
		}
		for _, prefix := range prefixes {
			if checkFn(word[len(prefix):]) {
				log.Println(prefix, "|", word)
				// checkFnMemo[word] = true
				return true
			}
		}
		// checkFnMemo[word] = false
		return false
	}

	result := []string{}
	for _, word := range words {
		for _, check := range wordsThatStartWith(word, wordStore) {
			// skip self reference
			if check == word {
				continue
			}
			if checkFn(check) {
				result = append(result, word)
				break
			}
		}
	}
	return result
}
