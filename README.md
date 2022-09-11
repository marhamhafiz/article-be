### PREPARE
After pull this app use the article.sql and import it to your MYSQL database.
Next, change 'root:password' in posts.go line 30 with your user:password in Mysql (root=> your database user, password=> your password database)

### RUN THIS APP WITH
first
``` bash
$ go mod tidy
```

then
``` bash
$ go run main.go
```

### AFTER THIS APP RUN
you can user postman article collection in this folder => 'Article.postman_collection.json' and import it to your postman

for default this app run in localhost:10000

