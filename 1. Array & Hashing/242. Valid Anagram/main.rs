use std::collections::HashMap;

impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        if s.len() != t.len() {
            return false;
        }
        let mut count_map: HashMap<char, i32>= HashMap::new();
        for c in s.chars() {
            *count_map.entry(c).or_insert(0) += 1;
        }

        for c in t.chars() {
            let count = count_map.entry(c).or_insert(0);
            *count -= 1;
            if *count < 0 {
                return false;
            }
        }
    return true
    }
}
