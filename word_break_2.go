// @@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@
// @ Copyright (c) Michael Leahcim                                                      @
// @ You can find additional information regarding licensing of this work in LICENSE.md @
// @ You must not remove this notice, or any other, from this software.                 @
// @ All rights reserved.                                                               @
// @@@@@@ At 2019-02-28 14:51 <thereisnodotcollective@gmail.com> @@@@@@@@@@@@@@@@@@@@@@@@

package main

// Working on this problem:
// https://leetcode.com/problems/word-break-ii/

// s = "catsanddog"
// wordDict = ["cat", "cats", "and", "sand", "dog"]

// Let's define a function `solution` that, given the input,
// returns the amount of existing solutions.

// solution(catsanddog) = [cats and dog, cat sand dog]
// solution(catsanddo) =  []
// solution(catsandd ) =  []
// solution(catsand  ) =  [cat sand, cats and]
// solution(catsan   ) =  []
// solution(catsa    ) =  []
// solution(cats     ) =  [cats]
// solution(cat      ) =  [cat]
// solution(ca       ) =  []
// solution(c        ) =  []
// solution(         ) =  []

// Let's define match function: match(suffix) => wordDict[suffix] or Nothing
// Let's recursively define solution function:
// solution(x) =   [each + " " + suffix for each in solution([len(x):-len(suffix)])]  when len(x)-len(suffix) > 0 and match(suffix)
//                 [suffix]                                                                len(x)-len(suffix) = 0 and match(suffix)
//                 []