import argparse

parser = argparse.ArgumentParser(description='Random testing')
parser.add_argument('-w', '--wait', help='Help for --w option')
options = parser.parse_args()

print(options.wait)
