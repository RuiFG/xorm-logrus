# Xorm-Logrus
######
xorm-logrus is a xorm logger implemented using logrus
# Installation
```shell script
go get github.com/RuiFG/xorm-logrus
```
# Simple Example
```go
import (
	"github.com/RuiFG/xorm-logrus"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"xorm.io/xorm"
)

func main() {
	DB, _ := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Asia%2FShanghai")
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	DB.SetLogger(xorm_logrus.NewLogrusLogger())
	DB.ShowSQL(true)
	ids := make([]int, 0)
	_ = DB.Table("asset").Cols("id").Find(&ids)
}
```

# License
This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.