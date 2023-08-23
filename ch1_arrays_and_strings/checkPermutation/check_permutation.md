### Problem: Check Permutation

You are given two strings, ```s1``` and ```s2```. Your task is to implement a function ```IsPermutation(s1, s2)``` that determines whether s1 is a permutation of s2. The function should return ```true``` if s1 and s2 are permutations of each other, and ```false``` otherwise.

A permutation of a string is another string that contains the same characters but possibly in a different order. 
For example, ```abc``` and ```bca``` are permutations of each other, while ```abc``` and ```def``` are not.

Write a Go function IsPermutation(s1, s2 string) bool to solve this problem.
```
func IsPermutation(s1, s2 string) bool {
}
```
Input:

Two strings s1 and s2
Output:

Return true if s1 is a permutation of s2, otherwise return false.
Example:
```
if IsPermutation("abc", "bca") != true {
// Handle incorrect implementation
}
```
```
if IsPermutation("abc", "def") != false {
// Handle incorrect implementation
}
```

Note:

- You should consider the characters' frequencies in the strings when checking for permutations.
- The comparison should be case-sensitive.