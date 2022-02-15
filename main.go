package main

import(
   gmail "gitee.com/silas/gmail"
)
func main(){
    foo := gmail.New("username@@gmail.com","app specify") //自己的邮箱地址和密码
    foo.Send("title","email content","usermaie@mailitor.com") //邮件标题 邮件内容 需要发送到的邮箱地址
}