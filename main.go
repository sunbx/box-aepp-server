package main

import (
	"box/models"
	_ "box/routers"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/beego/toolbox"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"strconv"
	"time"
)

//引入数据模型
func init() {
	orm.Debug = true
	//注册驱动
	//orm.RegisterDriver("mysql", orm.DRMySQL)
	//注册默认数据库
	//host := beego.AppConfig.String("db::host")
	//port := beego.AppConfig.String("db::port")
	//dbname := beego.AppConfig.String("db::databaseName")
	//user := beego.AppConfig.String("db::userName")
	//pwd := beego.AppConfig.String("db::password")
	//dbconnect := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	//_ = orm.RegisterDataBase("default", "mysql", dbconnect /*"root:root@tcp(localhost:3306)/test?charset=utf8"*/) //密码为空格式
	//
	//// 注册数据库
	//models.RegisterArticleDB()

}

func main() {

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 7776000

	beego.BConfig.WebConfig.Session.SessionProvider = "file"
	beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"
	startTask()
	beego.Run()
}

//A flag that prevents repetitive execution of timed tasks
var isTask = true

//Perform timed tasks
func startTask() {
	fmt.Println("")
	//The execution time of a timed startTast, executed every 10 seconds
	tk := toolbox.NewTask("PriceFeedTask", "0/60 * * * * *", func() error {
		if isTask {
			isTask = false
			SynAeBlock()
			isTask = true
		} else {
		}
		return nil
	})
	toolbox.AddTask("AbcLockPoliceTask", tk)
	toolbox.StartTask()
	fmt.Println("#########  ABC LOCK TOOL START SUCCESS #########")
}

func SynAeBlock() {

	//var publicKey = beego.AppConfig.String("PublicKey")
	var signingKey = beego.AppConfig.String("SigningKey")

	police, _ := models.SigningKeyHexStringAccount(signingKey)

	height := models.ApiBlocksTop()
	//if int(height) > models.LastHeight {

	models.IsCheckIng = true

	startTime := time.Now().Unix()

	fmt.Println("Checking Block " + strconv.Itoa(int(height)))

	myResult, _, err := models.CallStaticContractFunction(police.Address, models.ABCLockContractV3, "get_mapping_accounts", []string{})

	if err != nil {
		models.IsCheckIng = false
		fmt.Println("Function static error : " + err.Error())
		return
	}
	switch myResult.(type) {
	case []interface{}:

		data := myResult.([]interface{})
		models.LockAccountSize = len(data)
		fmt.Println("Lock account size : " + strconv.Itoa(models.LockAccountSize))
		for i := 0; i < len(data); i++ {
			account := data[i].([]interface{})[1].(map[string]interface{})

			address := account["account"].(string)
			mappingCount := account["mapping_count"].(json.Number).String()

			apiGetAccount, err := models.ApiGetAccount(address)

			if err != nil {
				fmt.Println("ApiGetAccount error :" + err.Error())
				continue
			}
			accountCountFloat64, err := strconv.ParseFloat(apiGetAccount.Balance.String(), 64)
			mappingCountFloat64, err := strconv.ParseFloat(mappingCount, 64)

			if err != nil {
				fmt.Println("ParseFloat error :" + err.Error())
				continue
			}

			if mappingCountFloat64 > accountCountFloat64 {

				fmt.Println("Illegal account detected address :" + address)
				fmt.Println("Illegal account detected mappingCount :" + mappingCount)
				fmt.Println("Illegal account detected balance :" + apiGetAccount.Balance.String())

				myResult, _, err := models.CallStaticContractFunction(police.Address, models.ABCLockContractV3, "is_account_check_balance", []string{address})

				if err != nil {
					fmt.Println("is_account_check_balance error :" + err.Error())
					continue
				}

				switch myResult.(type) {
				case bool:
					isAccountCheckBalance := myResult.(bool)
					if isAccountCheckBalance {

						tx, err := models.CallContractFunction(police, models.ABCLockContractV3, "account_check_balance", []string{address})
						if err != nil {
							fmt.Println("account_check_balance error :" + err.Error())
							continue
						}
						jsons, _ := json.Marshal(tx)
						println(string(jsons))

					}
				}
			}
		}

	}
	endTime := time.Now().Unix()

	models.ConsumingTime = endTime - startTime

	fmt.Println("Time consuming : " + strconv.Itoa(int(models.ConsumingTime)) + "ms")

	models.LastHeight = int(height)
	models.IsCheckIng = false
	//}
}
