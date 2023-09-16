Given two strings ```ransomNote``` and ```magazine```, return ```true``` if ransomNote can be constructed by using the letters from magazine and ```false``` otherwise.

Each letter in _magazine_ can only be used once in _ransomNote_.



**Example 1**:
<pre>
<b>Input</b>: ransomNote = "a", magazine = "b"
<b>Output</b>: false
</pre>

**Example 2**:

<b>Input</b>: ransomNote = "aa", magazine = "ab"
<b>Output</b>: false

**Example 3**:

<b>Input</b>: ransomNote = "aa", magazine = "aab"
<b>Output</b>: true


**Constraints**:

- ```1 <= ransomNote.length, magazine.length <= 105```
- ```ransomNote and magazine consist of lowercase English letters.```
