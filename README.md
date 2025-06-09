# MyChat

基于go+vue实现的聊天室+仿微信项目

## 项目概述
1. 简介：LishengChat 是一个前后端分离的即时通讯项目，具备后台管理、单聊群聊、联系人管理、多种消息（文本 / 文件 / 视频）处理、离线消息处理以及音视频通话等功能，旨在打造类似微信的聊天体验。
2. 技术栈：
   + 前端：Vue3、Vue Router、Vuex、WebSocket、Element - UI 等。
   + 后端：Go、Gin、GORM、GoRedis、WebSocket、Kafka、WebRTC、Zap 日志库等。


## 功能特性

#### 即时通讯服务
1. **多样化聊天模式**
    - **单聊与群聊体验**：系统提供一对一私密聊天和群组聊天两种模式，无论你是想与好友单独交流，还是和一群志同道合的人畅聊，都能轻松实现。所有消息都能实时推送，让你不错过任何重要信息。
    - **联系人灵活管理**：你可以自由地添加、删除联系人，还能将不友好的用户拉黑。同时，对于好友申请，你可以方便地进行处理，确保你的社交圈纯净有序。
2. **丰富的消息类型**：支持发送和接收多种类型的消息，包括文本、文件、音视频等。无论是分享日常趣事、重要文件，还是一段精彩的视频，都能轻松实现。
3. **可靠的离线消息处理**：即使你处于离线状态，消息也不会丢失。当你再次上线时，系统会自动将离线期间的消息推送给你，让你无缝衔接聊天内容。

#### 音视频通话功能
基于 WebRTC 技术，实现了 1 对 1 音视频通话功能。你可以自由发起通话，对方也能根据自己的情况选择拒绝、接收或挂断通话，让沟通更加直观和便捷。

#### 后台管理能力
系统配备了专门的后台管理界面，靓号用户可以通过该界面进行人员管控等维护操作，确保系统的正常运行和用户管理的高效性。

#### 安全验证机制
1. **SMS 短信验证**：在登录和注册过程中，采用 SMS 短信验证方式，增加账号的安全性，确保只有本人能够使用该账号。**但是！目前短信验证还在测试阶段，只能在阿里云短信验证添加测试手机号才能用短信验证。**

#### 数据持久化与性能优化
1. **数据库操作**：使用 GORM 进行后台 MySQL 数据库操作，确保数据能够持久化存储，即使系统出现故障，数据也不会丢失。
2. **日志记录与监控**：引入 Zap 日志库，对系统的运行情况进行详细记录。这些日志信息有助于开发人员快速排查问题，监控系统性能，确保系统的稳定运行。
3. **消息队列处理**：采用 Kafka 作为消息队列，能够高效地处理大量消息，确保消息的快速传输和处理，避免消息积压。
4. **Redis 缓存加速**：使用 GoRedis 进行缓存操作，将经常访问的数据存储在缓存中，减少数据库的访问压力，提高系统的响应速度和性能。

#### 实时消息推送
利用 WebSocket 技术实现实时消息推送，无论何时何地，只要有新消息，系统都会立即将其推送给用户，保证消息的及时性和准确性。 

## 项目结构

### 后端

```
lisheng-chat-server/
├── api/
│   └── v1/
│       └── chatroom_controller.go
│       └── controller.go
│       └── group_info_controller.go
│       └── message_controller.go
│       └── session_controller.go
│       └── user_contact_controller.go
│       └── user_info_controller.go
│       └── ws_controller.go
├── cmd/
│   └── lisheng-chat-server/
│       └── main.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── dao/
│   │   └── gorm.go
│   ├── dto/
│   │   ├── request/
│   │   │   └── ......
│   │   └── respond/
│   │   │   └── ......
│   ├── https_server/
│   │   └── https_server.go
│   ├── model/
│   │   ├── contact_apply.go
│   │   ├── group_info.go
│   │   ├── message.go
│   │   ├── session.go
│   │   ├── user_contact.go
│   │   └── user_info.go
│   └── service/
│       ├── chat/
│       │   ├── client.go
│       │   ├── kafka_server.go
│       │   └── server.go
│       ├── gorm/
│       │   ├── chatroom_service.go
│       │   ├── group_info_service.go
│       │   ├── message_service.go
│       │   ├── session_service.go
│       │   ├── user_contact_service.go
│       │   └── user_info_service.go
│       ├── kafka/
│       │   └── kafka_service.go
│       ├── redis/
│       │   └── redis_service.go
│       └── sms/
│           ├── local/
│           │   └── user_info_service_local.go
│           └── auth_code_service.go
├── logs/
│   └── test.log
├── pkg/
│   ├── constants/
│   │   └── constants.go
│   ├── enum/
│   │   ├── contact/
│   │   ├── contact_apply/
│   │   ├── group_info/
│   │   ├── message/
│   │   ├── session/
│   │   └── user_info/
│   ├── ssl/
│   │   ├── xxx.pem
│   │   ├── xxx-key.pem
│   │   └── tls_handler.go
│   ├── util/
│   │   └── random/
│   │       └── random_int.go
│   └── zlog/
│       └── logger.go
├── configs/
│   ├── config.toml
│   └── config_local.toml
├── static/
│   ├── avatars/
│   │   └── ......
│   └── files/
│   │   └── ......
├── web/
│   └── (前端项目结构)
├── .gitignore
├── go.mod
├── go.sum
└── README.md

```

### 前端

```
web/chat-server/
├── src/
│   ├── assets/
│   │   ├── cert/
│   │   │   ├── xxx.pem
│   │   │   ├── xxx-key.pem
│   │   │   └── mkcert.exe
│   │   ├── css/
│   │   │   └── chat.css
│   │   ├── img/
│   │   │   └── chat_server_background.jpg
│   │   ├── js/
│   │   │   ├── random.js
│   │   │   └── valid.js
│   ├── components/
│   │   ├── ContactListModal.vue
│   │   ├── DeleteGroupModal.vue
│   │   ├── DeleteUserModal.vue
│   │   ├── DisableGroupModal.vue
│   │   ├── DisableUserModal.vue
│   │   ├── Modal.vue
│   │   ├── NavigationModal.vue
│   │   ├── SetAdminModal.vue
│   │   ├── SmallModal.vue
│   │   └── VideoModal.vue
│   ├── router/
│   │   └── index.js
│   ├── store/
│   │   └── index.js
│   ├── views/
│   │   ├── access/
│   │   │   ├── Login.vue
│   │   │   ├── Register.vue
│   │   │   └── SmsLogin.vue
│   │   ├── chat/
│   │   │   ├── contact/
│   │   │   │   ├── ContactChat.vue
│   │   │   │   └── ContactList.vue
│   │   │   ├── session/
│   │   │   │   └── SessionList.vue
│   │   │   └── user/
│   │   │       └── OwnInfo.vue
│   │   ├── manager/
│   │       └── Manager.vue
│   ├── App.vue
│   └── main.js
├── .gitignore
├── package.json
├── README.md
└── vue.config.js
```

## 项目演示
### 登录注册

https://github.com/user-attachments/assets/8c2ef83d-4d3f-4dc3-8676-788ef6a14137

### 添加好友

https://github.com/user-attachments/assets/a46999ad-cb5c-4ec8-960c-c66b336aea07

### 发送接收消息

https://github.com/user-attachments/assets/52d8ec2e-2894-4597-a1a8-65e5cae86441

### 其他操作

https://github.com/user-attachments/assets/ed5f9d93-2544-43bf-b9d6-5e7756161d92

### 音视频通话

https://github.com/user-attachments/assets/6d037f8d-7529-439f-a373-7fdcf4dbd67b
