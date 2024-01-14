from pathlib import Path
import sys
import os
import subprocess
import re

if len(sys.argv) < 2:
    print(f'Usage: python {os.path.basename(__file__)} <path to out binary>')
    sys.exit(1)

bin_path = os.path.abspath(Path(sys.argv[1]))

help_output = subprocess.check_output([bin_path, '--help']).decode('utf-8')
usage = "```text\n" + help_output + "\n```"
version_output = subprocess.check_output([bin_path, '--version']).decode('utf-8')
if 'flag provided but not defined' in version_output:
    print(version_output)
    sys.exit(1)
version = version_output[version_output.find('version ') + len('version '):]
allowed_extensions = [re.search('\(([^)]+)', i).group(1).split(', ') for i in help_output.splitlines() if i.strip().startswith('<input-file>')][0]
supported_extensions = ''.join([f'\n* `.{i}`' for i in allowed_extensions])

readme = None
if Path('README.stub').exists():
    with open('README.stub') as f:
        readme_stub = f.read()
    readme = readme_stub
    readme = readme.replace('{VERSION}', version)
    readme = readme.replace('{SUPPORTED_EXTENSIONS}', supported_extensions)
    readme = readme.replace('{USAGE}', usage)
    readme = readme.strip()

if not readme:
    raise Exception('README.stub produces empty readme')

with open('README.md', 'w') as f:
    f.write(readme)
    print('[OK]')
