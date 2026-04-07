## Key Learnings / Takeaways

- Use **two-pointer technique**

  - `tp` → iterate over `t`
  - `sp` → track progress in `s`
- **Subsequence ≠ Substring**

  - Subsequence → order matters, **not necessarily continuous**
  - Substring → must be continuous
- **Core idea**

  - Move `tp` always
  - Move `sp` **only when characters match**
  - Never reset `sp`
- **Safe access pattern**

  - Always guard: `sp < len(s)` before accessing `s[sp]`
- **Edge cases**

  - `s == ""` → always `true`
  - `len(s) > len(t)` → always `false`
- **Time & Space**

  - Time: `O(n)` where `n = len(t)`
  - Space: `O(1)`

---

## Substring and Subsequence

**Substring****:** A **substring** is a **continuous (contiguous)** part of a string.

    👉 Characters must appear**together without gaps** .

Examples:

String: `"abcde"`

* `"abc"` ✅ (continuous)
* `"bcd"` ✅
* `"de"` ✅
* `"ace"` ❌ (not continuous, skipped characters)

**Subsequenc****e: **A **subsequence** is a sequence derived from a string by **removing some or no characters** **without changing the order** of the remaining characters.

    👉 Characters**do NOT need to be continuous** , but  **order must be preserved** .

Examples:

String: `"abcde"`

* `"abc"` ✅
* `"ace"` ✅ (skipped `b` and `d`)
* `"ae"` ✅
* `"aec"` ❌ (order changed)

### Key Difference

| Feature        | Substring      | Subsequence     |
| -------------- | -------------- | --------------- |
| Continuity     | ✅ Required    | ❌ Not required |
| Order          | ✅ Preserved   | ✅ Preserved    |
| Skipping chars | ❌ Not allowed | ✅ Allowed      |

### One-line Intuition

* **Substring** → "take a slice"
* **Subsequence** → "delete some characters"

---

## Approach-1 vs Approach-2

### Approach-1 (Explicit Edge Handling)

```go
if len(s) == 0 { return true }
if len(s) > len(t) { return false }

sp := 0
for tp := 0; tp < len(t); tp++ {
    if s[sp] == t[tp] {
        sp++
    }
    if sp == len(s) {
        return true
    }
}
return false
```

#### Pros

* Early exit for edge cases
* Slightly optimized for trivial inputs
* Clear intent

#### Cons

* Relies on **manual correctness guarantees**
* Can break if modified carelessly
* Needs extra boilerplate

---

### Approach-2 (Guarded / Safe Pattern)

```go
sp := 0
for tp := 0; tp < len(t); tp++ {
    if sp < len(s) && s[sp] == t[tp] {
        sp++
    }
}
return sp == len(s)
```

#### Pros

* **Safer (no out-of-bounds risk)**
* Handles edge cases implicitly
* Cleaner and more reusable
* Preferred in interviews / production

#### Cons

* Slight extra condition check (negligible) -> `sp < len(s)`

---

## Quick Comparison

| Aspect          | Approach-1         | Approach-2     |
| --------------- | ------------------ | -------------- |
| Safety          | ❌ Relies on logic | ✅ Always safe |
| Readability     | 🤝 Good            | 🤝 Good        |
| Edge handling   | Explicit           | Implicit       |
| Maintainability | ❌ Lower           | ✅ Higher      |
| Performance     | O(n)               | O(n)           |

---

## Final Recommendation

* Prefer **Approach-2 (guarded pattern)**
* Think in terms of:
  * "Scan `t`, consume `s` when possible"

```

```
