### Restful API application example in Go, based on Martini framework.
Some features out of the box:

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
~ go get github.com/3d0c/skeleton
```
Now it's ready to use. Default configuration supposes that MongoDB listen at least on localhost:27017 and dosn't require auth.

```
~ cd $GOPATH/src/github.com/3d0c/skeleton
~ go run app.go
2014/01/12 15:03:38 app.go:104: Waiting for connections...
```

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
       -u "testuser:xxx" \
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

Here is one sample object "Posts", which implements common REST behaviour:

```
GET		/posts		avail url opts: &offset=(int), &limit=(int)
GET		/posts/:id
POST	/posts		-d '{"title": "", "body": ""}'
PUT		/posts/:id
DELETE	/posts/:id
```


#### Add something.
For example, let's add comments to the blog posts.  

**1. Starting from Scheme**

Add CommentScheme to the ```models/schemes.go```

```go
type CommentScheme struct {
	Id     bson.ObjectId `bson:"_id,omitempty" json:"id"      binding:"-"`
	Uid    bson.ObjectId `bson:"uid,omitempty" json:"uid"     binding:"-"`
	PostId bson.ObjectId `bson:"pid,omitempty" json:"post_id" binding:"require"`
	Body   string        `bson:",omitempty"    json:"body"`
}

```

We could add Validation method right here, e.g.:

```go
func (this CommentScheme) Validate(errors *binding.Errors, req *http.Request) {
	if len(this.Body) == 0 {
		errors.Fields["body"] = "Body can't be empty."
	}
}

```

**2. Creating Model**

Create ```models/commentsmodel.go``` 

```go
package models

import (
	"github.com/3d0c/martini-contrib/model"
)

type Comments struct {
	*model.Model
}

func (*Comments) Construct(args ...interface{}) *Comments {
	this := &Comments{
		Model: model.New(CommentScheme{}),
	}

	return this
}

```

That's all. Our model already has Create(), Find(), Update(), Delete() methods, which work with a CommentScheme.

**3. Routing and Controller**

Now, inside app.go add route to the create method:

```go
route.Post("/comments",
	models.Construct(models.Users{}, credentials),
	models.Construct(models.Posts{}),
	models.Construct(models.Comments{}),
	binding.Bind(models.CommentScheme{}),
	controllers.CheckAuth,
	controllers.CommentsCreate,
)

```

It's pretty straightforward — initialize all models, bind values, check auth and run CommentsCreate

Creating ```controllers/comments.go```

```go
package controllers

import (
	"github.com/3d0c/martini-contrib/encoder"
	"github.com/3d0c/skeleton/models"
	// "github.com/codegangsta/martini"
	// "labix.org/v2/mgo/bson"
	"net/http"
)

func CommentsCreate(u models.Users, post models.Posts, comment models.Comments, params models.CommentScheme, enc encoder.Encoder) (int, []byte) {
	var result interface{}

	// we've got here a post model, so we could check
	// is there a post with 'id', given by user.

	params.Uid = u.Object.Id

	if result = comment.Create(&params); result == nil {
		return http.StatusInternalServerError, []byte{}
	}

	return http.StatusOK, encoder.Must(enc.Encode(result))
}

```

That's all. Let's check:

```sh
curl -XPOST -i  -H "Content-Type:application/json" \
     -u "testuser:xxx" \
     —d '{"post_id":"52cea9a1fcb05d3338000001", "body":"Clever comment"}' \
     http://localhost:8000/comments
```

Sure, you have to provide valid post_id from previously created post.

Expected response:

```
{
	"id"      : "52d2808ffcb05d6fe2000001",
	"uid"     : "52d26dcefcb05d6e4f000002",
	"post_id" : "52cea9a1fcb05d3338000001",
	"body"    : "Clever comment"
}
```


#### Models relations.
Coming soon…
