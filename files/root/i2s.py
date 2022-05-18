import xlwt
import re
import glob

book = xlwt.Workbook()
sheet1 = book.add_sheet("筹款信息")

sheet1.write(0,0,"项目标题")
sheet1.write(0,1,"项目筹款额")
sheet1.write(0,2,"捐款人数")
sheet1.write(0,3,"目标筹款额")

f = glob.glob(r"C:\Users\86152\Desktop\p\*.txt")
lineNum = 1
for txt in f:
	w = open(txt, 'r', encoding="gb2312").read()

	sheet1.write(lineNum, 0, re.findall(r"\['(.*) \| 微博", w)[0])
	#sheet1.write(lineNum, 1, re.findall(r"已筹款\(元\) ([0-9]*(\.[0-9]*)?)", w))
	sheet1.write(lineNum, 1, re.findall(r"已筹款\(元\) (.*)\n", w)[0])
	sheet1.write(lineNum, 2, re.findall(r"捐款人次 (.*)\n", w)[0])
	sheet1.write(lineNum, 3, re.findall(r"目标\(元\)'\] (.*)\n", w)[0])
	lineNum += 1

book.save("筹款信息.xls")
