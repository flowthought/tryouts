name_file = open("hero-names.txt", "r")
url_file = open("guide-urls.txt", "r")
out_file = open("join.md", "w")

names = []
for name in name_file:
    names.append(name)

urls = []
for url in url_file:
    urls.append(url)

for i in range(len(names)):
    out_file.write(names[i] + urls[i] + '\n')
