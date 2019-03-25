
import argparse

def setup_parameters():
    """
    Set up the command line parameters
    """

    parser = argparse.ArgumentParser(description='Sample project')
    parser.add_argument('-l', '--loglevel', action='store', metavar='Log level',
                        dest='log_level', type=int,
                        help='Specify the log level')
    parser.add_argument('-v', '--version', action='store_true', dest='version',
                        help='Display version number')

    return parser.parse_args()

