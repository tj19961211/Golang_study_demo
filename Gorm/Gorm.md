# Gorm

## 基本介绍

```go
// 链接数据库
db, err := gorm.Open("sqlite3", "test.db")  //gorm v1

// 创建数据库表
db.Create(&Product{Code: "L1212", Price: 1000})
db.Save(&product) // 如果主键有值更新，否则创建


//读取
var product Product
db.First(&product, 1) // 查询id为 1 的product
db.First(&product, "code=?", "L1212") // 查询code为L1212的product

// 更新product的price为2000
db.Model(&product).Update("Price", 2000)

db.Updates(&product)  // 只更新非零值字段


```