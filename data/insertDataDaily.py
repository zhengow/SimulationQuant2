#!/usr/bin/python3
import pymysql
import tushare as ts
import logging
import time
import datetime

def insert():
  data = pro.stock_basic(exchange='', list_status='L', fields='ts_code')
  today = time.strftime("%Y%m%d", time.localtime()) 

  db = pymysql.connect('localhost','root','123456','stock_pricevolume' )

  cursor = db.cursor()

  symbols = []
  for symbol in data.ts_code:
    symbols.append(symbol)

  for i in range(0,len(symbols),100):
    codes = ''
    for symbol in symbols[i:i+100]:
      codes += symbol + ','
    codes = codes[:-1]

    df = pro.daily(ts_code=codes, start_date=today, end_date=today)
    
    for ii in range(len(df)):
      table = df.ts_code[ii][:6]
      logging.info("start insert table: " + table)
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

if __name__ == "__main__":
  LOG_FORMAT = "%(asctime)s - %(levelname)s - %(message)s"
  logging.basicConfig(filename='insertDaily.log', level=logging.INFO, format=LOG_FORMAT)
  pro = ts.pro_api()

  while True:
    d_time = datetime.datetime.strptime(str(datetime.datetime.now().date()) + '20:00', '%Y-%m-%d%H:%M')
    d_time1 = datetime.datetime.strptime(str(datetime.datetime.now().date()) + '21:00', '%Y-%m-%d%H:%M')
    n_time = datetime.datetime.now()
    if n_time > d_time and n_time < d_time1:
      insert()
    else:
      time.sleep(3600)

