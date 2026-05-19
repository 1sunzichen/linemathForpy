import requests
import json

def call_psm_api():
    try:
        params = {
            'psm': 'life.hermes.msg_base',
            'event': {},
            'context': {},
            'body': {}
        }
        
        response = requests.get(
            'https://life-qa-ftf-api.byted.org/api/qa_infra/diff/psm/infov2',
            params=params,
            timeout=5
        )
        
        # Process nested data structure
        result = {
            "code": 200,
            "data": {
                "code": 0,
                "data": {
                    "has_more": False,
                    "psm_info": response.json().get('data', {}).get('psm_info', [{}])[0],
                    "total": 1
                },
                "msg": ""
            }
        }
        return json.dumps(result, indent=2, ensure_ascii=False)

    except Exception as e:
        return json.dumps({
            "code": 500,
            "message": f"API调用失败: {str(e)}"
        }, ensure_ascii=False)

if __name__ == "__main__":
    print(call_psm_api())