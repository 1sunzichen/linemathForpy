
def extract_psm(raw_list):
    """Equivalent to the comprehension in your plugin_job.py"""
    return [
        item['psm'] if isinstance(item, dict) and 'psm' in item else str(item)
        for item in raw_list
    ]

# ---------- Test cases ----------
test_cases = [
    # Normal dict
    [{'psm': 'life.a', 'revision': 'r1'}, {'psm': 'life.b', 'revision': 'r2'}],
    # Missing psm field
    [{'revision': 'r3'}, {'psm': 'life.c'}],
    # Mixed dict / string / other types
    [{'psm': 'life.d'}, 'life.e', 123, {'psm': 'life.f'}],
    # Empty list
    [],
]

if __name__ == '__main__':
    for idx, case in enumerate(test_cases, 1):
        result = extract_psm(case)
        print(f'Case {idx}: {case} -> {result}')