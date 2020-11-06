#!/usr/bin/python3

import pymysql
import tushare as ts
import logging

LOG_FORMAT = "%(asctime)s - %(levelname)s - %(message)s"
logging.basicConfig(filename='create.log', level=logging.DEBUG, format=LOG_FORMAT)

pro = ts.pro_api()

data = pro.stock_basic(exchange='', list_status='L', fields='symbol')

db = pymysql.connect('localhost','root','123456','stock_pricevolume' )

cursor = db.cursor()

for symbol in data.symbol:
  cursor.execute("DROP TABLE IF EXISTS %s" % ('`'+symbol+'`'))
  sql = """CREATE TABLE %s (
          trade_date INT PRIMARY KEY NOT NULL,
          open decimal(10,2),
          high decimal(10,2),
          low decimal(10,2),
          close decimal(10,2),
          pre_close decimal(10,2),
          vol decimal(10,2),
          amount double
          )""" % ('`'+symbol+'`')
  cursor.execute(sql)
  logging.info("success create table: " + symbol)
db.close()