mysql_gen_data
==============

生成数据样本作为mysql测试数据

./gen_data -h

Usage of ./gen_data:

  -f="": row format %m[-n]d,%m[-n]s,%[m]n  #数据格式 设置m,n表示某数字区间或者字符串长度区间，d随机数字,n自增顺序数字,s字符串
  
  -n=10000: row number   #生成的数据记录条数


