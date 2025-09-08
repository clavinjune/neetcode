use std::collections::HashMap;

impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut sum_map = HashMap::with_capacity(nums.len());
        for (i, &num) in nums.iter().enumerate() {
            if let Some(&j) = sum_map.get(&(target - num)) {
                return vec![j as i32, i as i32];
            }
            sum_map.insert(num, i);
        }
        vec![]
    }
}
