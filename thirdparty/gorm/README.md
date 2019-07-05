# 使用GORM連結MySql
1. 使用`docker-compose up -d`快速架設一個mysql伺服器
2. 使用sql指令創造新的usr&db，並且把db權限賦予該usr
```sql
create user 'jim' IDENTIFIED by 'password';
create database `demo_db`;
grant all privileges on demo_db.* to 'jim';
```


# refer
- golang的orm操作手冊
https://gorm.io/docs/
- gorm tutorial
https://tutorialedge.net/golang/golang-orm-tutorial/
- 使用orm連接到mysql
https://github.com/jinzhu/gorm/issues/403
- mysql創造新的database遇到的坑
https://stackoverflow.com/questions/44916136/error-1064-42000-when-creating-database-in-mysql
- 討論是否每次使用gorm.Open以後要在程序結束以前使用gorm.Close()
https://github.com/jinzhu/gorm/issues/1427
- 如何查看gorm具體執行的sql語句
https://github.com/jinzhu/gorm/issues/1544
- golang在做(關聯)對多的情況
https://blog.csdn.net/rocky0503/article/details/80915157

# 中文參考
- https://www.bookstack.cn/read/gorm-cn-doc/crud.md
- https://segmentfault.com/a/1190000013216540