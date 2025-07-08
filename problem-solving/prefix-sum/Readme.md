## ✅ What is Prefix Sum?

The **Prefix Sum** of an array is a technique where we precompute cumulative sums at each index, allowing us to quickly answer **range sum queries** or **range-based transformations** in constant time.

---

## 📌 Example

Let’s say:
```
arr = [2, 4, 6, 8]
```
Then the prefix sum array will be:
```
prefix_sum = [2, 6, 12, 20]
```
🧠 How to Calculate Prefix Sum
To calculate the prefix sum at index i, we use:
```
prefix_sum[i] = prefix_sum[i - 1] + arr[i]
```
▶️ Step-by-Step Example
Let’s say i = 2:
```
prefix_sum[2] = prefix_sum[1] + arr[2]
             = 6 + 6
             = 12
```

✅ This approach is useful for solving subarray sum problems, sliding window computations, and many dynamic programming challenges efficiently.

## 🛠️ When to Use Prefix Sum?
- Range sum queries

- Frequent subarray sum calculations

- Difference array transformations

- Matrix region sum (2D prefix sums)

- Counting frequencies or ranges efficiently

- Sliding window optimizations
  
