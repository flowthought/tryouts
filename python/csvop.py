from collections import OrderedDict

with open('baseline.csv') as csvfile:
    rdr = csv.DictReader(csvfile)
    records = []
    for row in rdr:
            records.append(row)

# Records is a list of ordered dict
