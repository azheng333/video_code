# coding=utf-8


from urllib import urlopen
from BeautifulSoup import BeautifulSoup
import re


webpage = urlopen('http://bbs.tianya.cn/').read()

divBegin = webpage.find('<div class="wrap-bd-in bbs-list-box show clearfix"')

article = webpage[divBegin:(divBegin + 3000)]

allTitle = re.findall('<a href=".*" target="_blank">(.*)</a>', article)
allLink = re.findall('<a href="(.*)" target="_blank"', article)

for i in range(len(allTitle)):
    # print 'titile: ', allTitle[i]  #打印每个标题
    # print 'href: ', allLink[i]  #打印每个标题对应的链接
    # print '-' * 100  #打印分割线
    articlePage = urlopen(allLink[i]).read()
    soup = BeautifulSoup(articlePage)
    paragrapList = soup.findAll('div', {"class": "bbs-content clearfix"})
    for v in paragrapList:
        print v.text[:100]
        print '-' * 100




