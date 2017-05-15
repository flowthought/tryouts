import argparse

parser = argparse.ArgumentParser(description='Collect performance traces of PDF first page rendering under Edge')
parser.add_argument('-w', '--wait', default=15, help='Time to wait for PDF first page render to complete')
options = parser.parse_args()

print(type(options.wait))
