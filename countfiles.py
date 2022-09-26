import os 
import pathlib

fReadme = open("./Readme.md", "w")

def countLines(path):
    count = 0
    for f in os.listdir(path):
        fPath = path + "/" + f
        file_ext = pathlib.Path(fPath).suffix
        if (os.path.isfile(fPath) and file_ext in ['.go']):
            fCount = len(open(fPath).readlines())
            fReadme.write(str(f) + "(" + str(fCount) + "), ")
            count = count + fCount
    fReadme.write("S(" + str(count) + ")\n")
    return count


total = 0
fReadme.write("##Categories")
for f in os.listdir('./src-ig'):
    pathFull = "./src-ig/" + f
    fReadme.write("\n##" + str(f).upper() + "\n")
    if (os.path.isdir(pathFull)):
        total = total + countLines(pathFull)

print(total)


fReadme.write("\n#Total python code \n" + str(total))