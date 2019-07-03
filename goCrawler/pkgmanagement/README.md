# 爬蟲練習
1. 訪問網頁https://golang.google.cn/pkg/
2. 抓取網頁下所有標準庫的名稱/連結(href標籤)/敘述
3. 劃分pkg的層級
4. 開一個mysql的table:`pkg_content`，把資料存進去
5. 使用gRPC創建一個管理pkg的Server，提供使用者輸入pkg回覆對應的pkg連結/文檔內容，以及是否有父級別pkg;有的話就帶上父級別的pkg
6. 使用net/http監聽一個端口，實現對外提供一個http api接口，接口可以接收pkg名稱，返回json格式的pkg信息
7. 接口需要使用jwt確認安全，並且該接口會應用到gRPC client去拿pkg資訊，並返回給呼叫者
8. 使用nginx_pass代理http api接口


# 環境設定
1. 架設mysql與phpadmin: docker-compose up -d
   1. 登入帳號(密碼): root/user(root/user)
   2. 從終端機登入mysql: docker exec -it mysql mysql -u root -p ... 密碼 root
2. 




# refer
- https://github.com/PuerkitoBio/goquery
- https://github.com/fuadajip/dockercompose-mysql-phpmyadmin
- https://stackoverflow.com/questions/27933866/use-goquery-to-find-a-class-whose-value-contains-whitespace
- https://studygolang.com/articles/4602