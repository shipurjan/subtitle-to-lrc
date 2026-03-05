"""This script computes the supported extensions by listing the directories in the converter folder.
It excludes the 'shared' directory from the list of supported extensions."""

import glob
import os
from pathlib import Path
import sys
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
    """Compute the supported extensions by listing the directories in the converter folder."""
    supported_extensions = [
        os.path.basename(Path(dir)) for dir in glob.glob('./converter/*/')
        if os.path.basename(Path(dir)) != 'shared'
    ]
    logger.info(', '.join(supported_extensions))


if __name__ == '__main__':
    main()
