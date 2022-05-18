
class LNode:
    def __init__(self,elem):
        self.elem = elem
        self.next = None
        self.head = None
class LinkList:
    def __init__(self):
        self.head=None
        self.changdu=0

        self.tail = None
    def insert(self,l):#首位插入
        l.next = self.head
        self.head.head = l
        self.head = l
        if self.changdu == 0:
            self.tail = l
        
        self.changdu = self.changdu + 1
    def isNoneLink(self):#判断是否是空列表
        if self.changdu ==0:
            print("这个列表是空列表")
        else:
            print("这个列表不是空列表，长度为{}".format(self.changdu))

    
    def drop1st(self): #首位元素删除操作
        self.head = None
        self.head.next.head = None
        self.changdu = self.changdu - 1

        if self.changdu == 0:
            self.tail = None

    def append(self, l): # 尾端添加元素
        self.tail.next = l
        l.head=self.tail
        self.tail=l

    def droptail(self):
        

    
        
l1 = LinkList()
q1 = LNode(12)
q2 = LNode(13)
q3 = LNode(14)
l1.isNoneLink()
l1.insert(q1)
l1.isNoneLink()
l1.insert(q2)
l1.isNoneLink()
l1.insert(q3)
l1.isNoneLink()

q0 = LNode(4)
l1.insert(q0)
print(l1.changdu)



