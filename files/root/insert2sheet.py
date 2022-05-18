import xlwt
import re

f = open("/home/jeff/Documents/tmp/shiryou.txt")
book = xlwt.Workbook()
sheet1 = book.add_sheet("筹款信息")

sheet1.write(0,0,"项目标题")
sheet1.write(0,1,"项目筹款额")
sheet1.write(0,2,"捐款人数")
sheet1.write(0,3,"目标筹款额")

lineNum = 1
while True:
	line = f.readline()
	if line =="":
		break
	sheet1.write(lineNum, 0, re.findall(r'项目标题：(.*) 项目筹款额', line))
	sheet1.write(lineNum, 1, re.findall(r'项目筹款额：(.*) 捐款人数', line))
	sheet1.write(lineNum, 2, re.findall(r'捐款人数：(.*) 目标筹款额', line))
	sheet1.write(lineNum, 3, re.findall(r'目标筹款额：(.*) ', line))
	lineNum += 1

book.save("筹款信息.xls")
f.close()