qeydiyyatdan kecirmek ucun

localhost:3000/account (POST)
POST body
{
  "firstName":"Sadagat",
  "lastName":"Asgarov",
  "password":"test"
}

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


istifadeci oz id ile oz idsini baxa biler
localhost:3000/account/(istifadecinin idsi)
id ile baxmaq tek tek istifadecilere
