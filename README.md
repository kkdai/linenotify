LineNotify template : A simple Golang Line Notify Bot template
==============

[![Join the chat at https://gitter.im/kkdai/linenotify](https://badges.gitter.im/kkdai/linenotify.svg)](https://gitter.im/kkdai/linenotify?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

 [![GoDoc](https://godoc.org/github.com/kkdai/linenotify.svg?status.svg)](https://godoc.org/github.com/kkdai/linenotify)  [![Build Status](https://travis-ci.org/kkdai/linenotify.svg?branch=master)](https://travis-ci.org/kkdai/linenotify.svg)

[![goreportcard.com](https://goreportcard.com/badge/github.com/kkdai/linenotify)](https://goreportcard.com/report/github.com/kkdai/linenotify)


## Just want to try it?

- [Go to my test Line Notify site](https://linenotify-app.herokuapp.com/auth)

![](img/notify1.png)

- Click this button and remember to login your Line account for authenication this notify.

![](img/notify2.png)

- Select one on one notification, click "Agreed and Connected"
- It will help you to add friend with "Line Notify".

![](img/notify3.png)

- Go to this link add arbitrary string after msg=xxxx ex: https://linenotify-app.herokuapp.com/notify?msg=test


![](img/notify4.png)



## Installation for developer

### 1. Got A Line Bot API devloper account

[Register your Line Notify Account](https://notify-bot.line.me/my/services/new)

- You need to fill all related info you need.
- For "Service Site" and "Callback URL", just fill arbitrary web site with "http://xxx.xxx.com"
- Remember you need "Client ID" and "Client Secret" for Heroku setup.

![](img/linenotify-1.png)



### 2. Just Deploy the same on Heroku

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

Remember your heroku ID. 

### 3.

<br><br>

## Reference

- [Line Notify Doc](https://notify-bot.line.me/doc/en/)
- [實作 Line Notify 通知服務 (1)](https://poychang.github.io/line-notify-1-basic/)
