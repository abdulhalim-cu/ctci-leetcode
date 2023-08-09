## URLify a given string (Replace spaces with %20)

Write a method to replace all the spaces in a string with ‘```%20```’. You may assume that the string has sufficient space at the end to hold the additional characters and that you are given the “```true```” length of the string.
Examples:
```
Input: "Mr John Smith", 13
Output: Mr%20John%20Smith
```
```
Input: "Mr John Smith   ", 13
Output: Mr%20John%20Smith
```
A simple solution is to create an auxiliary string and copy characters one by one. Whenever space is encountered, place %20 in place of it.
