# ini adalah Final Project 3 Golang Studi Independen
 yaitu membuat rest API bernama Kanban Board, yang merupakan sebuah aplikasi untuk project management, setiap user pada aplikasi ini akan dapat menambahkan task task mereka pada category category yang telah di sediakan oleh seorang admin. project ini mengunakan gorm, gin, dan juga database mysql.

## Untuk menggunakannya, bisa dengan aplikasi postman atau insomnia. 
Lalu menulisakan alamat atau url https://finalproject3-group5.up.railway.app/

## Enpoint User

Sebelum melakukan perintah CRUD, diharuskan untuk melakukan POST register terlebih dahulu pada url https://finalproject3-group5.up.railway.app/users/register
-Saat melakukan register, diperlukan full_name, email, password, dan role yang harus diisi dengan form,
selain dengan form register juga bisa diisi menggunakan syntax json pada bagian body. Misalnya : 
{ 
    "full_name" : "Fajrian Nugraha", 
    "email" : "member@gmail.com", 
    "password" : "member123" 
    "role" : member
}
(perlu diketahui juga bahawa untuk role disini akan terdefault menjadi member, dan untuk admin sudah otomatis terbuat ketika aplikasi di jalankan(seeding data))

#setelah melakukan register, user harus melakukan login pada url atau https://finalproject3-group5.up.railway.app/users/login di postman atau insomnia 

Untuk melakukan login cukup dengan mengisikan email dan password. Lalu token pada bagian response nya harus di-copy untuk melakukan perintah CRUD
#Untuk melakukan perintah pada user, seperti Update dan Delete user bisa dengan melakukan Request PUT pada url https://finalproject3-group5.up.railway.app/update-account
dan Request DELETE pada url https://finalproject3-group5.up.railway.app/delete-account dan tidak memerlukan userId

#Perintah yang bisa dilakukan, yaitu POST, PUT, GET, dan DELETE. Tetapi sebelum itu, pada bagian Header wajib ditambahkan "Authorization" dan memasukkan token yang telah didapat dari login tadi agar perintahnya berjalan.

## endpoint Categories

perlu diketahui bahwa untuk bisa melakukan request POST,PATCH dan, DELETE di endpoint categories ini hanya bisa dilakukan oleh admin saja tetapi untuk GET bisa oleh role admin dan member

untuk melakukan POST atau menginsert Categories, bisa dengan url POST https://finalproject3-group5.up.railway.app/categories/ lalu memasukkan data "type" nya
dan juga bisa input menggunakan syntax json pada bagian body, misalnya : 
{
    "type":"Olahraga"
}

untuk melakukan GET atau melihat Categories, bisa dengan url GET https://finalproject3-group5.up.railway.app/categories/

untuk melakukan PATCH atau mengupdate Categories, bisa dengan menambahkan categoryId dibelakang url, misalnya untuk mengubah photo dengan id 3, maka urlnya PATCH https://finalproject3-group5.up.railway.app/categories/3

untuk melakukan DELETE atau menghapus categories bisa dengan menambahkan categoryId dibelakang url, misalnya jika ingin menghapus photo dengan id 4, maka urlnya DELETE https://finalproject3-group5.up.railway.app/categories/4


## endpoint TASK

untuk melakukan POST atau menginsert task, bisa dengan url POST https://finalproject3-group5.up.railway.app/tasks/ lalu memasukkan data seperti "title", "description","status","category_id" dengan form, 
dan juga bisa input menggunakan syntax json pada bagian body, misalnya : 
{
    "title":"Bermain Bola", 
    "description": Main bola di lapangan bola, 
    "status": true, 
    "category_id": 1, 
}
perlu diketahui default dari field status adalah false

untuk melakukan GET atau melihat task, bisa dengan url GET https://finalproject3-group5.up.railway.app/tasks/

untuk melakukan PUT atau mengupdate task, bisa dengan menambahkan taskId dibelakang url, misalnya untuk mengubah task dengan id 3, maka urlnya PUT https://finalproject3-group5.up.railway.app/tasks/3

untuk melakukan PACTH atau mengupdate status task, bisa dengan menambahkan taskId dibelakang url, misalnya untuk mengubah status task dengan id 3, maka urlnya PATCH https://finalproject3-group5.up.railway.app/tasks/update-status/3

untuk melakukan PACTH atau mengupdate category task, bisa dengan menambahkan taskId dibelakang url, misalnya untuk mengubah category task dengan id 3, maka urlnya PATCH https://finalproject3-group5.up.railway.app/tasks/update-category/3

untuk melakukan DELETE atau menghapus task bisa dengan menambahkan taskId dibelakang url, misalnya jika ingin menghapus task dengan id 4, maka urlnya DELETE https://finalproject3-group5.up.railway.app/tasks/4

# Kanban Board

Repository ini telah di Deploy melalui Railway.app `https://finalproject3-group5.up.railway.app/`

## How To Run

We use Makefile to run the project.

```bash
make run
```

## Our Team

- Pande Putu Devo Punda Maheswara
- Hanif Fadillah Amrynudin
- I Putu Agus Arya Wiguna
