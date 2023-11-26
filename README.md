qeydiyyatdan kecirmek ucun

localhost:3000/account (POST)
POST body
{
  "firstName":"Sadagat",
  "lastName":"Asgarov",
  "password":"password"
}
`{
  "firstName":"Sadagat",
  "lastName":"Asgarov",
  "password":"password"
}`
Butun istifadecileri siralamaq ucun
localhost:3000/account (GET)
{
  "id": 0,
  "firstName": "Sadagat",
  "lastName": "Asgarov",
  "number": 656314,
  "balance": 0,
  "createdAt": "2023-11-25T04:21:26.137130896Z"
}

localhost:3000/account (POST)

  {
    "number":,
    "password":"password"

  }

istifadeci oz id ile oz idsini baxa biler
localhost:3000/account/(istifadecinin idsi)
id ile baxmaq tek tek istifadecilere


Proyekti ise salmaq ucun docker-compose.yml fayli

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

