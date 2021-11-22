package main
import (
   "fmt"
   "github.com/spf13/viper"
   "os"
)
 
func main() {
   path, err := os.Getwd()                          //获取项目的执行路径
   if err != nil {
      panic(err)
   }
     
   config := viper.New()                            //实体化名为 config
   config.AddConfigPath(path)                       //设置读取的文件路径
   config.SetConfigName("config")                   //设置读取的文件名
   config.SetConfigType("yaml")                     //设置文件的类型
 
   if err := config.ReadInConfig(); err != nil {    //尝试进行配置读取
      panic(err)
   }
 
   //打印文件读取出来的内容:
   fmt.Println(config.Get("database.host"))
   fmt.Println(config.Get("database.user"))
   fmt.Println(config.Get("database.dbname"))
   fmt.Println(config.Get("database.pwd"))
}
