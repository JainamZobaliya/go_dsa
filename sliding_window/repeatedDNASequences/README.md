## Key Takeaways: Repeated DNA Sequences

- Use **sliding window (fixed size = 10)**

  - Iterate from `0 → len(s) - 10`
- Use a **hash map (frequency counter)**

  - Track how many times each substring appears
- Add to result **only when count becomes 2**

  - Avoid duplicate entries in output
- **Time Complexity**

  - O(n) (each window processed once)
- **Space Complexity**

  - O(n) (storing substrings in map)
- **String slicing in Go**

  - `s[i:i+10]` is O(1) (no copy)
  - But hashing string for map → O(10)
- **Edge case**

  - If `len(s) < 10` → return empty result
- **Pattern learned**

  - Fixed-size sliding window + hashmap for frequency
- **Optimization insight (advanced)**

  - Replace string with **bit encoding (2 bits per char)**
  - Use rolling hash to reduce memory and improve speed
