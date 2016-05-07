# coding=utf-8

import re

string = 'hello world hello'
# match = re.match('hello', string)
# match = re.match('world', string)
# match = re.search('world', str)
match = re.match('(\w+) (\w+) (\w+)', string)
if match is not None:
    print match.group()
    print match.group(1)
    print match.group(2)
    print match.group(3)
else:
    print 'None'

matchList = re.findall('hello', string)
print matchList


import os

print os.system('date')
f = os.popen('date')
today = f.read()
print 'today is ', today

import subprocess

print subprocess.call('date')

print subprocess.call(['ls', '-l'])
print subprocess.call('ls -l', shell=True)

# import sys
#
# p = subprocess.Popen('ping -c 5 www.baidu.com',
#                      shell=True, stdout=subprocess.PIPE)
# output, err = p.communicate()
#
# while True:
#     if p.poll() is not None:
#         break
#
#     if p.stdout is not None:
#         out = p.stdout.readline()
#         if out == '':
#             break
#         else:
#             sys.stdout.write(out)
#             sys.stdout.flush()







