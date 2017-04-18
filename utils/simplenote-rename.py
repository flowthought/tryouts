#!/usr/bin/python
import re
import os

notes = [n for n in os.listdir('.') if n.endswith('.txt')]
for note in notes:
    with open(note, 'r') as f:
        name = f.readline()
        name = re.sub(' ', '_', name.casefold())
        name = re.sub('[\W]', '', name) + '.markdown'
        
        os.rename(note, name)
