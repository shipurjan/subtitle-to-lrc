import glob
import os
from pathlib import Path
import sys


def error(msg=''):
    print(f'[ERROR] {msg}')
    sys.exit(1)


def ok(msg=''):
    print(f'[OK] {msg}')
    sys.exit(0)


def main():
    supported_extensions = [
        os.path.basename(Path(dir)) for dir in glob.glob('./converter/*/')
        if os.path.basename(Path(dir)) != 'shared'
    ]
    print(', '.join(supported_extensions))


if __name__ == '__main__':
    main()
