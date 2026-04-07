## Key Learnings

- 2 Pointer Technique
- Strings → immutable → convert to []byte
- No `.includes()` → use map lookup
- Always move pointers correctly in two-pointer problems
- Be careful with types (byte vs string)

---

## Approach-1 vs Approach-2

### Approach-1:

```go
for ; left<=right; {
	if !vowels[chars[left]]{
		left++
		continue
	} else if !vowels[chars[right]]{
		right--
		continue
	} else {
		chars[left], chars[right] = chars[right], chars[left]
		left++
		right--
	}
}
```

### Approach-2:

```go
for left < right {
	for left < right && !vowels[chars[left]] {
		left++
	}
	for left < right && !vowels[chars[right]] {
		right--
	}
	chars[left], chars[right] = chars[right], chars[left]
	left++
	right--
}
```



### 1. Cleanliness / Readability

- **Second version ✅**

  - Separates "move pointer" logic from "swap logic"
  - More readable and standard pattern
- **First version ❌**

  - Mixed conditions (`if / else if`)
  - Easier to mess up pointer movement

---

### 2. Bug Risk

- **First version ❌**

  - High chance of:
    - infinite loop
    - wrong pointer movement
    - missed increments
- **Second version ✅**

  - Safer pattern:
    ```
    find valid left
    find valid right
    swap
    move both
    ```

---

### 3. Performance (micro-level)

- Second version is slightly better in practice because:

  - Fewer conditional checks per iteration
  - Skips invalid chars in bulk (inner loops)
- Asymptotically same: **O(n)**

---

## Mental Model (Important Takeaway)

The second approach follows a standard 2-pointer template:

```go
for left < right {
    move left to valid position
    move right to valid position
    do work (swap)
    move both pointers
}
```


👉 This pattern is reusable for:

* Reverse vowels
* Valid palindrome
* Two sum (sorted)
* Container with most water

---

## Final Verdict

| Aspect      | First Version | Second Version |
| ----------- | ------------- | -------------- |
| Time        | O(n)          | O(n)           |
| Space       | O(n)          | O(n)           |
| Readability | ❌ Lower      | ✅ Higher      |
| Bug Risk    | ❌ High       | ✅ Low         |
| Pattern Use | ❌ Ad-hoc     | ✅ Standard    |

👉  **Second version is clearly better** , even though complexity is same.
