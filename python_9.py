import sqlite3


def main():
    #conn = sqlite3.connect('test.db')
    conn = sqlite3.connect(':memory:')
    cur = conn.cursor()

    cur.execute('''create table
    customers(id interger primary key,
    name text, city text)''')

    cur.execute('''insert into customers(id, name, city)
    values(1, 'weizheng', 'zh')''')
    cur.execute('''insert into customers(id, name, city)
    values(2, 'xiaohua', 'zh')''')

    cur.execute('''select * from customers''')
    print cur.fetchall()

    cur.execute('''update customers set name='xiaoming'
    where id=1''')

    cur.execute('''select * from customers''')
    print cur.fetchall()

    cur.execute('''delete from customers''')

    conn.close()


# import MySQLdb
#
# def main():
#     conn = MySQLdb.connect(host='127.0.0.1',
#                     port=3306,
#                     user='weizheng',
#                     passwd='weizheng',
#                     db='test1')
#
#     cur = conn.cursor()
#
#     cur.execute('''create table
#     customers(id int primary key,
#     name char(20),
#     city char(20))''')
#
#     cur.execute('''insert into customers(id, name, city)
#     values(1, 'weizheng', 'zh')''')
#     cur.execute('''insert into customers(id, name, city)
#     values(2, 'xiaohua', 'zh')''')
#
#     cur.execute('''select * from customers''')
#     print cur.fetchall()
#
#     cur.execute('''update customers set name='xiaoming'
#     where id=1''')
#
#     cur.execute('''select * from customers''')
#     print cur.fetchall()
#
#     cur.execute('''delete from customers''')
#
#     conn.close()


if __name__ == '__main__':
    main()




