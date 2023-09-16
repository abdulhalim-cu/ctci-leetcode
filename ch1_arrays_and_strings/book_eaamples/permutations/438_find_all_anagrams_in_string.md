> Cracking the coding interview book example:
> Given a smaller string s and a bigger string b, design an algorithm to find all permutations of the shorter
    string withing the longer one. Print the location of each permutation 


Given two strings ```s``` and ```p```, return an array of all the start indices of ```p's``` anagrams in ```s```. You may return the answer in any order.

An **Anagram** is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

<b>Example 1</b>:

<pre>
<b>Inputs:</b> s = "cbaebabacd", p = "abc"
<b>Output:</b> [0,6]
<b>Explanation:</b>
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
</pre>

<b>Example 2</b>:

<pre>
<b>Inputs:</b> s = "abab", p = "ab"
<b>Output:</b> [0,1,2]
<b>Explanation:</b>
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
</pre>

<b>Constraints</b>:

- ```1 <= s.length, p.length <= 3 * 104```
- ```s and p consist of lowercase English letters```