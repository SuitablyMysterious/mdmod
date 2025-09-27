# Take .mdmod files and validate their input to be valid

import pyyaml

def validate_file(file_path, encoding='utf-8'):
    with open(file_path, 'r', encoding) as file:
        try:
            content = pyyaml.safe_load(file_path)
            return content # change for prod to actually check the file is markdown 
        except Exception as e:
            return None