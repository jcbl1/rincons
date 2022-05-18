class dot:
    def __init__(self,label):
        self.neighbor=[]
        self.label=label

class graph:
    def __init__(self):
        self.dots=[]
        self.root=None
    def addOneDot(self,dot1):
        if dot1 not in self.dots:
            self.dots.append(dot1)
    def addDots(self,dotx):
        for i in dotx:
            if i not in self.dots:
                self.dots.append(i)
    def addEdge(self,dot1,dotx):
        for i in dotx:
            if i not in dot1.neighbor:
                dot1.neighbor.append(i)
            if dot1 not in i.neighbor:
                i.neighbor.append(dot1)
    def calDegree(self):
        deg=[]
        for i in self.dots:
            deg.append(len(i.neighbor))
        print(deg)
        return(deg)
    def gdbl(self,dotx):
        bl=[dotx]
        kk = 0
        while len(bl)<len(self.dots):
            for i in bl[kk].neighbor:
                if i not in bl:
                    bl.append(i)
            kk+=1
        for i in bl:
            print(i.label)
        return(bl)
    def sdbl(self,dotx,bl):
        if dotx not in bl:
            bl.append(dotx)
            for i in dotx.neighbor:
                self.sdbl(i,bl)





# 操作1区
dot1=dot('1')
dot2=dot('2')
dot3=dot('3')
dot4=dot('4')
dot5=dot('5')

g1=graph()

g1.addDots([dot1,dot2,dot3,dot4,dot5])

g1.addEdge(dot1,[dot2,dot3,dot4])
g1.addEdge(dot2,[dot3,dot5])
g1.addEdge(dot3,[dot4])
g1.addEdge(dot4,[dot5])

dott1 = dot('1')
dott2 = dot('2')
dott3 = dot('3')
dott4 = dot('4')
dott5 = dot('5')
dott6 = dot('6')
dott7 = dot('7')
dott8 = dot('8')

g2 = graph()

g2.addDots([dott1,dott2,dott3,dott4,dott5,dott6,dott7,dott8])

g2.addEdge(dott1,[dott2,dott3,dott4])
g2.addEdge(dott2,[dott5])
g2.addEdge(dott3,[dott4,dott6])
g2.addEdge(dott4,[dott5,dott7])
g2.addEdge(dott5,[dott8])
g2.addEdge(dott6,[dott8])

# 操作1区，以上

g1.calDegree()
g1.gdbl(dot1)

print("计算g2所有点的度")
g2.calDegree()
print("广度遍历")
g2.gdbl(dott1)
print("深度遍历")
bl=[]
g2.sdbl(dott1,bl)
for i in bl:
    print(i.label)


print("深度遍历生成树")
d1 = '1'
d2 = '2'
d3 = '3'
d4 = '4'
d5 = '5'
d6 = '6'
d7 = '7'
d8 = '8'
t = graph()
dots_ = [d1,d2,d3,d4,d5,d6,d7,d8]
t.addDots(dots_)
for i in range(len(bl)):
    if i < len(bl) - 1:
        if bl[i] in bl[i+1].neighbor:
            t.addEdge(dots_[i],[dots_[i+1]])
        else:
            for j in range(i):
                if bl[j] in bl[i+1].neighbor:
                    t.addEdge(dots_[j],[dots_[i]])
    else:
        for j in range(i):
            if bl[j] in bl[i+1].neighbor:
                t.addEdge(dots_[j],[dots_[i]])

