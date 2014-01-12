### Restful API application example in Go, based on Martini framework.
Some features from the box:

- DRY Models (right now, only MongoDB supported)
- Controllers
- Auth (ready to use user model)
- Schemes
- Named connections manager
- Simple config

### Prerequisite.
This application uses MongoDB, so if you don't already have it, install it.

### Getting started.
```
go get github.com/3d0c/skeleton
```
Now it's ready to use. Default configuration supposes that MongoDB listen at least on localhost:27017 and dosn't require auth.

**Let's create a user**:  

```sh
curl -XPOST -i  -H "Content-Type:application/json" \
       -d '{"login": "testuser", "password":"xxx"}' \
       http://localhost:8000/users
```

Expected response:

```
{ 
	"id" 	 :  "52d00ab0fcb05d57cd000001",
	"login"   :  "testuser",
	"profile" : {
		"first_name" : "",
		"last_name" : ""
	}
}
```

**Getting user profile (login)**:

```sh
curl -XGET -i \
       -u "testuser:xxx" \
       http://localhost:8000/user
```

Expected response: ```200 OK``` and

```
{ 
	"id" 	 :  "52d00ab0fcb05d57cd000001",
	"login"   :  "testuser",
	"profile" : {
		"first_name" : "",
		"last_name" : ""
	}
}
```

or ```401 Unauthorized``` if you gave wrong credentials.

**Updating user**.

```sh
curl -XPUT -i  -H "Content-Type:application/json" \
       -d '{"profile": {"first_name": "xxx"}}' \
       http://localhost:8000/user
```

Expected response: ```200 OK``` and

```
{ 
    "id"      :  "52d00ab0fcb05d57cd000001",
    "login"   :  "testuser",
    "profile" : {
        "first_name" : "xxx",
        "last_name"  : ""
    }
}
```

**All other**:

Here is one sample object "Posts", wich implements common REST behaviour:

```
GET		/posts		avail url opts: &offset=(int), &limit=(int)
GET		/posts/:id
POST	/posts		-d '{"title": "", "body": ""}'
PUT		/posts/:id
DELETE	/posts/:id
```


#### Add something.
For example, let's add a comments to the blog.

