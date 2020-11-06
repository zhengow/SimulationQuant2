#!/usr/bin/python3

import pymysql
import tushare as ts
import logging

LOG_FORMAT = "%(asctime)s - %(levelname)s - %(message)s"
logging.basicConfig(filename='insert.log', level=logging.INFO, format=LOG_FORMAT)

pro = ts.pro_api()

data = pro.stock_basic(exchange='', list_status='L', fields='ts_code')

db = pymysql.connect('localhost','root','123456','stock_pricevolume' )

cursor = db.cursor()

for symbol in data.ts_code:
  table = symbol[:6]
  logging.info("start insert table: " + table)
  df = pro.daily(ts_code=symbol, start_date='20100101', end_date='20201105')
  for ii in range(len(df)):
    trade_date = int(df.iloc[ii].trade_date)
    ops = df.iloc[ii].open
    high = df.iloc[ii].high
    low = df.iloc[ii].low
    close = df.iloc[ii].close
    pre_close = df.iloc[ii].pre_close
    vol = df.iloc[ii].vol
    amount = df.iloc[ii].amount
    sql = """INSERT INTO %s
         VALUES (%d, %f, %f, %f, %f, %f, %f, %f)""" % ('`'+table+'`', trade_date, ops,
                                                    high, low, close, pre_close, vol, amount)
    try:
      cursor.execute(sql)
      db.commit()
    except Exception as e:
      logging.error("Fail to insert table: " + table)
      logging.error(e)
      db.rollback()

db.close()