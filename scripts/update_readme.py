"""This script updates the README.md file based on the output of the out binary."""

from pathlib import Path
import sys
import os
import subprocess
import re
import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


def error(msg=''):
    """Log an error message and exit with a non-zero status code."""
    logger.error(msg)
    sys.exit(1)


def ok(msg=''):
    """Log an informational message and exit with a zero status code."""
    logger.info(msg)
    sys.exit(0)


def main():
    """Update the README.md file based on the output of the out binary."""
    if len(sys.argv) < 2:
        error(
            f'Usage: python {os.path.basename(__file__)} <path to out binary>')

    bin_path = os.path.abspath(Path(sys.argv[1]))

    help_output = subprocess.check_output([bin_path, '--help']).decode('utf-8')
    version_output = subprocess.check_output(
        [bin_path, '--version']).decode('utf-8')

    if 'flag provided but not defined' in version_output:
        error(version_output)

    usage = "```text\n" + help_output + "\n```"
    version = version_output[version_output.find(
        'version ') + len('version '):]
    supported_extensions = [re.search(r'\(([^)]+)', i).group(1).split(', ')
                    for i in help_output.splitlines() if i.strip().startswith('<input-file>')][0]
    supported_extensions_str = ''.join(
        [f'\n* `.{i}`' for i in supported_extensions])

    readme = None
    if Path('README.stub').exists():
        with open('README.stub', encoding='utf-8') as f:
            readme_stub = f.read()

        readme = readme_stub
        readme = readme.replace('{VERSION}', version)
        readme = readme.replace(
            '{SUPPORTED_EXTENSIONS}', supported_extensions_str)
        readme = readme.replace('{USAGE}', usage)
        readme = readme.strip()

    if not readme:
        error('README.stub produces empty readme')

    with open('README.md', 'w', encoding='utf-8') as f:
        f.write(readme)
        ok('README.md updated')


if __name__ == '__main__':
    main()
