## 🧩 Problem: Remove Duplicates from Sorted Array

Given a **sorted array**, remove duplicates **in-place** such that:

- Each element appears **only once**
- Return the number of unique elements (`k`)
- First `k` elements of the array should contain the result

👉 Order must be preserved
👉 No extra array allowed (O(1) space)

---

## 🧠 Simple Intuition

> Move all **unique elements to the front**, ignore the rest

---

## ⚙️ Approach-1 (Deleting elements ❌)

### Idea

- Use two pointers
- When duplicate found → **remove it using `append`**

### Problems

- ❌ Deleting shifts elements → costly
- ❌ Time complexity becomes **O(n²)**
- ❌ Modifying slice during iteration → bugs / index issues

---

## ⚙️ Approach-2 (Overwriting ✅)

### Idea

- Use two pointers:

  - `p1` → last unique element index
  - `p2` → scan array
- When new element found:

  - Move `p1`
  - Copy value → `nums[p1] = nums[p2]`

### Why it works

- No deletion
- Just overwrite duplicates

---

## ⚡ Key Differences

| Aspect     | Approach-1 ❌   | Approach-2 ✅      |
| ---------- | --------------- | ------------------ |
| Strategy   | Delete elements | Overwrite elements |
| Time       | O(n²)          | O(n)               |
| Space      | O(1)            | O(1)               |
| Safety     | Error-prone     | Safe               |
| Efficiency | Poor            | Optimal            |

---

## 🏁 Final Takeaway

> ❌ Don’t remove elements
> ✅ Just shift unique elements forward
