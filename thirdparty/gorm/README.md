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
- gorm在做(關聯)對多的情況
https://blog.csdn.net/rocky0503/article/details/80915157
- goorm做transaction
https://motion-express.com/blog/gorm:-a-simple-guide-on-crud

# 中文參考
- https://www.bookstack.cn/read/gorm-cn-doc/crud.md
- https://segmentfault.com/a/1190000013216540


# 後記: gorm單元測試的坑...
一開始使用go-mocket(sql-mock)包一層的測試框架，參考
- https://github.com/DATA-DOG/go-sqlmock/issues/118#issuecomment-386692428
發現還是不是很好操作，有些呼叫到底層sql-mock的錯誤，不太好排查...於是打算從sql-mock重新刻一個...但是sql-mock也是會遇到create時發生一些錯誤，於是參考到
- https://github.com/jinzhu/gorm/issues/711#issuecomment-167469666
決定先轉用sqlite3當作測試db(os: 反正就是測試完畢刪除一個檔案哩...)