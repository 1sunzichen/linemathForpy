import re

def test_merge_tag_matching(env, psm="test_psm", process_id=123):
    """
    测试merge_tag正则匹配逻辑
    模拟cov_tag_caller.py中的匹配规则
    """
    print(f"测试字符串: '{env}'")
    
    # 这里复制你要测试的匹配逻辑
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

# 测试用例
test_cases = [
    "merge_tag",           # 精确匹配
    "MERGE_TAG",           # 大写
    "Merge_Tag",           # 混合大小写
    "prefix_merge_tag_suffix",  # 包含子串
    "this has merge_tag in middle",  # 中间包含
    "no_merge_here",       # 不匹配
    "ppe_env",             # 匹配PPE
    "PPE_ENV",             # 大写PPE
    "boe_env",             # 匹配BOE
    "random_string",       # 默认PPE
    "",                    # 空字符串
    None                   # None值（会报错，测试边界情况）
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