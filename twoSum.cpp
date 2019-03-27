class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> m; // map structure
        vector<int> res;
        for (int i = 0; i < nums.size(); ++i) { // restore all num into map
            m[nums[i]] = i; // value is index
        }
        for (int i = 0; i < nums.size(); ++i) {
            int t = target - nums[i];
            if (m.count(t) && m[t] != i) {
				// m.count(key)：由于map不包含重复的key，因此m.count(key)取值为0，或者1，表示是否包含。
                res.push_back(i);
                res.push_back(m[t]);
                break;
            }
        }
        return res;
    }
};