# API文档
## Basic Response Body

序号  |参数|类型|规则|
---- | -----| -----|-----|
1|code|number|正常响应200|
2|message|string|{...}|
3|data|object|实例或null|
4|status|number|0失败，1成功|

## 1.注册

- URL：/v1/api/register  

- method:POST

- Request Headers

        "Content-Type": "application/json"

- Request Body

      {  
            "username": "aaa",  
    
            "password": "123456"    
      }
- Response Body

        {
           "code": 200,
           
           "data": null,
           
           "message": "注册成功"
        }

        {
          "code": 422,
          
          "data": null,
          
          "message": "密码不能少于6位"
         }
## 2.登录
- URL:v1/api/login

- method:POST

- Request Headers

        "Content-Type": "application/json"

- Request Body

        {  
            "username": "aaa",  
    
            "password": "123456"    
        }
- Response Body

        {
        "code": 200,
        "data": {
        "token":      "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjEsImV4cCI6MTU4NzAyOTU1NCwiaWF0IjoxNTg2NDI0NzU0LCJpc3MiOiJoYWNrd2VlayIsInN1YiI6InVzZXIgdG9rZW4ifQ.U35akfANrcQqx2zI25lQAsJFopQMbYawHYKNAVTdGjY"
        },
        "message": "注册成功"
        }

     
       {
      "message": "密码错误",
      "status": 0
        }
        
        或
        
       {
        "code": 422,
        "message": "密码不能少于6位"
        }
## 3.获取用户信息

- URL：/v1/api/info 

- method:GET

- Request Headers

        "Content-Type": "application/json"

- Request Body

      {  
            "username": "aaa",  
    
            "password": "123456"    
      }
- Response Body
        
       {
          "code": 200,
          "data": {
          "user": {
            "name": "aaa"
            }
         }
       }
     
        {
         "message": "权限不足",
         "status": 0
        }

## 4.发表故事

- URL：/v1/api/post 

- method:POST

- Request Headers

        "Content-Type": "application/json"

- Request Body

      {  
            "story": "aaa",      
      }
- Response Body

      {
        "code": 200,
        "data": {
        "imagurl": "ABC",
        "name": "aaa",
        "tag": "ABC",
        "text": "ABC",
        "title": "ABC"
        }，
        "message": "发表成功"
        }
        
        
        {
        "code": "422",
        "message": "内容不能为空"
        }
        
    


 
        
        
        
        
        
