import re

def test_merge_tag_matching(env, psm="test_psm", process_id=123):
    """
    Test merge_tag regex matching logic
    Simulate the matching rules in cov_tag_caller.py
    """
    print(f"测试字符串: '{env}'")
    
    # Copy the matching logic you want to test here
    if re.search(r'(?i)merge_tag', env) is not None:
        result = "ATE_" + "#" + psm + "#" + str(process_id)
        print(f"✅ 匹配成功: {result}")
        return result
    elif re.search(r'(?i)ppe', env) is None:
        result = "ATE_BOE" + "#" + psm + "#" + str(process_id)
        print(f"⚠️  匹配BOE: {result}")
        return result
    else:
        result = "ATE_PPE" + "#" + psm + "#" + str(process_id)
        print(f"❌ 匹配PPE: {result}")
        return result

# Test cases
test_cases = [
    "merge_tag",           # Exact match
    "MERGE_TAG",           # Uppercase
    "Merge_Tag",           # Mixed case
    "prefix_merge_tag_suffix",  # Contains substring
    "this has merge_tag in middle",  # Contains in middle
    "no_merge_here",       # No match
    "ppe_env",             # Match PPE
    "PPE_ENV",             # Uppercase PPE
    "boe_env",             # Match BOE
    "random_string",       # Default PPE
    "",                    # Empty string
    None                   # None value (will error, testing edge case)
]

print("=" * 60)
print("测试 merge_tag 正则匹配")
print("=" * 60)

for i, test_env in enumerate(test_cases, 1):
    print(f"\n测试 {i}:")
    try:
        test_merge_tag_matching(str(test_env) if test_env is not None else "")
    except Exception as e:
        print(f"❌ 错误: {e}")

print("\n" + "=" * 60)
print("测试完成")