
def extract_psm(raw_list):
    """等价于你在 plugin_job.py 里的那一句推导"""
    return [
        item['psm'] if isinstance(item, dict) and 'psm' in item else str(item)
        for item in raw_list
    ]

# ---------- 测试用例 ----------
test_cases = [
    # 正常 dict
    [{'psm': 'life.a', 'revision': 'r1'}, {'psm': 'life.b', 'revision': 'r2'}],
    # 缺 psm 字段
    [{'revision': 'r3'}, {'psm': 'life.c'}],
    # 混 dict / 字符串 / 其它类型
    [{'psm': 'life.d'}, 'life.e', 123, {'psm': 'life.f'}],
    # 空列表
    [],
]

if __name__ == '__main__':
    for idx, case in enumerate(test_cases, 1):
        result = extract_psm(case)
        print(f'Case {idx}: {case} -> {result}')