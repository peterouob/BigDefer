# BigDefer

### 需求？

1. 用戶操作(登入，註冊，重置密碼，刪除，在公共地方發送消息)
2. 後台檢查(檢查token後->看到所有人，可以對他們進行增刪改查，黑名單策略，文章、留言審查，監控使用者即時發送訊息)
3. 用戶發文->文章具備Title和內容，可以點讚、評論(文章ID為用戶登入token)
4. 文章留言為樹狀結構，會有留言下的留言
5. 1對1私訊功能，對相同group發送訊息(group中的所有人都會接受到，groupID為創造者token)

### 工具？
後端->Golang

前端->Vue3/Vite

API->foxAPI

Docker -> mysql,redis,mongo

### 前端
#### 儲存：Pinia
#### 請求：axios
#### 路由：vue-router
#### 組建：element-plus

### 後端

#### framework : GIN 
- https://gin-gonic.com/zh-cn/docs/
#### framework : JWT
- https://github.com/dgrijalva/jwt-go

#### Config : Viper
- https://github.com/spf13/viper
#### Database : Mysql->user/文章/留言 
- GORM：https://gorm.io/
- 留言：https://blog.csdn.net/qq_42887507/article/details/119960591

#### Database : redis->token
- Redis：https://redis.uptrace.dev/zh/
- 黑名單：https://juejin.cn/post/7022579299349692452
#### Database : mongo->聊天(websocket)/後續可轉MQTT
- Mongo：https://github.com/mongodb/mongo-go-driver
- Websocket：https://github.com/gorilla/websocket
- 聊天：https://medium.com/@lalit.garghate/create-websocket-in-golang-to-take-data-from-mongodb-5e90651611a6
#### 消息對列：Kafaka

### 文件包
- router -> 管理路由規則
- utils ->全域使用包
- static ->前端以及文欓，目前暫無前端
- model -> 管理模塊(使用者，資料庫)
- service -> 和model進行溝通，放GIN方法，http使用模組集合
- docker -> 管理dockerCompose，dockerfile
- test -> 測試各種方法

## 遇到的問題
### 使用docker建mysql環境導致權限不足
- 解決:https://ithelp.ithome.com.tw/questions/10200426
- 新使用者 newuser/password
```SQL
mysql -uroot -p
CREATE USER 'newuser'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON * . * TO 'newuser'@'%';
```
### Gorm的Create不斷造成Goroutine阻塞
- 解決:使用gorm/gen
- https://gorm.io/gen/
- 原因？
  -  後序處理連表方便，而且具備使用經驗
- 解決2，聲明db時不要用短聲明，因為後面未必使用而導致錯誤(純用gorm的情況)
- 解決3，初始化mysql時，如果使用的是gen的話要加上dal.Default(db)
### PORT被佔用
- lsof -i:<要查找的port>，接著kill他的編號

## 舉報問題思路
https://blog.csdn.net/qq_45627684/article/details/108576906
1. 發出檢舉請求給伺服端->伺服端回應以接受
2. 伺服端查看檢舉類行(不同處罰)
3. 檢查言論
