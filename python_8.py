import exceptions


# def main():
# a = 1 / 0
# print 'end'

# def main():
# for i in dir(exceptions):
# print i


# def main():
# try:
#         a = 1 / 0
#     except ZeroDivisionError:
#         print "you can't divide by zero"

# def main():
#     try:
#         d = one / 0
#     except (NameError, ZeroDivisionError):
#         print "a local or global name is not found"
#         print "you can't divide by zero"

# def main():
#     try:
#         d = one / 0
#     except (NameError, ZeroDivisionError) as e:
#         print e.message, e.args

#
# def main():
#     try:
#         d = one / 0
#     except NameError as e:
#         print e.message, e.args
#     except ZeroDivisionError:
#         print "you can't divide by zero"


# def main():

#         d = 1 / 2
#     except (NameEoor, ZeroDivisionError) as e:
#         print "you can't divide by zero"
#         print e
#     else:
#         print d

# def main():
#     raise Exception('test')

# def main():
#     try:
#         raise Exception('test')
#     except Exception as e:
#         print e

# class MyError(Exception):
#     def __init__(self, value):
# 	    self.value = value


# def main():
#     try:
#         raise MyError(2*2)
#     except MyError as e:
#         print 'My exception occurred, value:', e.value

# def main():
#     str = raw_input('please input your name:')
#     print "your name is: ", str

import os

# def main():
#     try:
#         students = []
#         f = open('/Users/weizheng/Downloads/name.txt')
#     except IOError as e:
#         print 'file not found', e
#     else:
#         for line in f:
#             students.append(line)
#
#         print students
#
#         pos = f.seek(0, 0)
#
#         str = f.read()
#
#         f.close()
#
#         print str
#
#         fo = open('/Users/weizheng/Downloads/students.txt', 'w')
#
#         fo.write(str)
#
#         fo.close()

def main():
    print os.getcwd()

    dirlist = os.listdir('/usr')
    for file in dirlist:
        if os.path.isdir('/usr/' + file):
            print file
        elif os.path.isfile('/usr/' + file):
            print file
        else:
            continue
    os.mkdir('/Users/weizheng/Downloads/test')
    os.rename('/Users/weizheng/Downloads/test', '/Users/weizheng/Downloads/test1')
    os.rmdir('/Users/weizheng/Downloads/test1')

if __name__ == '__main__':
    main()
