# Read a csv into memory
from collections import OrderedDict

with open('baseline.csv') as csvfile:
    rdr = csv.DictReader(csvfile)
    records = []
    for row in rdr:
            records.append(row)

# Records is a list of ordered dict

x = [1, 2, 3]
y = [4, 5, 6]
sum_x_y = [i + j for (i, j) in zip(x, y)]
