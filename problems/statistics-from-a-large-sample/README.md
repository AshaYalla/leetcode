<!--|This file generated by command(leetcode description); DO NOT EDIT.    |-->
<!--+----------------------------------------------------------------------+-->
<!--|@author    Openset <openset.wang@gmail.com>                           |-->
<!--|@link      https://github.com/openset                                 |-->
<!--|@home      https://github.com/openset/leetcode                        |-->
<!--+----------------------------------------------------------------------+-->

[< Previous](https://github.com/openset/leetcode/tree/master/problems/shortest-common-supersequence "Shortest Common Supersequence ")
　　　　　　　　　　　　　　　　
[Next >](https://github.com/openset/leetcode/tree/master/problems/car-pooling "Car Pooling")

## 1093. Statistics from a Large Sample (Medium)

<p>We sampled integers between <code>0</code> and <code>255</code>, and stored the results in an array <code>count</code>:&nbsp; <code>count[k]</code> is the number of integers we sampled equal to <code>k</code>.</p>

<p>Return the minimum, maximum, mean, median, and mode of the sample respectively, as an array of <strong>floating point numbers</strong>.&nbsp; The mode is guaranteed to be unique.</p>

<p><em>(Recall that the median of a sample is:</em></p>

<ul>
	<li><em>The middle element, if the elements of the sample were sorted and the number of elements is odd;</em></li>
	<li><em>The average of the middle two elements, if the elements of the sample were sorted and the number of elements is even.)</em></li>
</ul>

<p>&nbsp;</p>
<p><strong>Example 1:</strong></p>
<pre><strong>Input:</strong> count = [0,1,3,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
<strong>Output:</strong> [1.00000,3.00000,2.37500,2.50000,3.00000]
</pre><p><strong>Example 2:</strong></p>
<pre><strong>Input:</strong> count = [0,4,3,2,2,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]
<strong>Output:</strong> [1.00000,4.00000,2.18182,2.00000,1.00000]
</pre>
<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ol>
	<li><code>count.length == 256</code></li>
	<li><code>1 &lt;= sum(count) &lt;= 10^9</code></li>
	<li>The mode of the sample that count represents is unique.</li>
	<li>Answers within <code>10^-5</code> of the true value will be accepted as correct.</li>
</ol>

### Related Topics
  [[Math](https://github.com/openset/leetcode/tree/master/tag/math/README.md)]
  [[Two Pointers](https://github.com/openset/leetcode/tree/master/tag/two-pointers/README.md)]

### Hints
<details>
<summary>Hint 1</summary>
The hard part is the median.  Write a helper function which finds the k-th element from the sample.
</details>