Qeydiyyatdan kecmek ucun
(POST)
token verir onu headerse yazmaq qalir
ve login olan zaman o token ile yoxlanilir

http://localhost:3000/account
```
{
  "firstName":"Sadagat",
  "lastName":"Asgarov",
  "password":"password"
}
```

Butun istifadecileri siralamaq ucun
(GET)

[http://localhost:3000/account ](url)
```
{
  "id": 0,
  "firstName": "Sadagat",
  "lastName": "Asgarov",
  "number": 656314,
  "balance": 0,
  "createdAt": "2023-11-25T04:21:26.137130896Z"
}
```

Qeydiyyatdan kecmis istifadecinin giris etmesi
number (POST)

[http://localhost:3000/account](url) 

 ```
 {
    "number":,
    "password":"password"

  }
```

Proyekti ise salmaq ucun docker-compose.yml fayli
```
version: "3.8"
services:
  db:
    image: postgres
    container_name: local_pgdb
    restart: always
    ports:
      - "5432:5432"
    environment:
      POSTGRES_USER: user-name
      POSTGRES_PASSWORD: strong-password
    volumes:
      - local_pgdata:/var/lib/postgresql/data


  my-json:
    image: sadagatasgarov/my-json-app:latest
    container_name: myjson
    restart: always
    ports:
      - "3000:3000"
    environment:
      HOSTAPP: local_pgdb
      USERAPP: user-name
      PASSAPP: strong-password
      DB: postgres

volumes:
  local_pgdata:
```


