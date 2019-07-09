# 這是一個爬蟲系統
1. 透過crawler去抓取指定的網頁內容
2. 透過grpc server做傳遞訊息
```
crawler --|
          | --grpc_server-- grpc_client(net/http listen server) -- nginx
mysql ----|
```
3. 