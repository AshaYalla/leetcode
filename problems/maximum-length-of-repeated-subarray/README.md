<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    Openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

## 718. Maximum Length of Repeated Subarray (Medium)

<p>Given two integer arrays <code>A</code> and <code>B</code>, return the maximum length of an subarray that appears in both arrays.</p>

<p><b>Example 1:</b><br />
<pre>
<b>Input:</b>
A: [1,2,3,2,1]
B: [3,2,1,4,7]
<b>Output:</b> 3
<b>Explanation:</b> 
The repeated subarray with maximum length is [3, 2, 1].
</pre>
</p>

<p><b>Note:</b><br>
<ol>
<li>1 <= len(A), len(B) <= 1000</li>
<li>0 <= A[i], B[i] < 100</li>
</ol>
</p>

### Related Topics
[[Array](https://github.com/openset/leetcode/tree/master/tag/array/README.md)]
[[Hash Table](https://github.com/openset/leetcode/tree/master/tag/hash-table/README.md)]
[[Binary Search](https://github.com/openset/leetcode/tree/master/tag/binary-search/README.md)]
[[Dynamic Programming](https://github.com/openset/leetcode/tree/master/tag/dynamic-programming/README.md)]

### Similar Questions
  1. [Minimum Size Subarray Sum](https://github.com/openset/leetcode/tree/master/problems/minimum-size-subarray-sum) (Medium)

### Hints
  1. Use dynamic programming.  dp[i][j] will be the answer for inputs A[i:], B[j:].