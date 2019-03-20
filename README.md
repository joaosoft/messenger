# messenger
[![Build Status](https://travis-ci.org/joaosoft/messenger.svg?branch=master)](https://travis-ci.org/joaosoft/messenger) | [![codecov](https://codecov.io/gh/joaosoft/messenger/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/messenger) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/messenger)](https://goreportcard.com/report/github.com/joaosoft/messenger) | [![GoDoc](https://godoc.org/github.com/joaosoft/messenger?status.svg)](https://godoc.org/github.com/joaosoft/messenger)

A simple messenger with socket notification. (This implementation should use for example rabbitmq to keep messages when shutdown)

## Support for 
> Http

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## Dependency Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`

## Configuration
```
{
  "messenger": {
    "host": "localhost:8001",
    "token_key": "banana",
    "socket": {
      "client": {
        "server_address": "localhost:9001/api/v1",
        "log": {
          "level": "info"
        }
      },
      "server": {
        "address": ":9001",
        "log": {
          "level": "info"
        }
      }
    },
    "dbr": {
      "db": {
        "driver": "postgres",
        "datasource": "postgres://user:password@localhost:7000/postgres?sslmode=disable&search_path=messenger"
      }
    },
    "log": {
      "level": "info"
    },
    "migration": {
      "path": {
        "database": "schema/db/postgres"
      },
      "db": {
        "driver": "postgres",
        "datasource": "postgres://user:password@localhost:7000/postgres?sslmode=disable&search_path=messenger"
      },
      "log": {
        "level": "info"
      }
    }
  },
  "manager": {
    "log": {
      "level": "info"
    }
  },
  "socket": {
    "client": {
      "server_address": "localhost:9001/api/v1",
      "log": {
        "level": "info"
      }
    }
  }
}
```

>### Go
```
go get github.com/joaosoft/messenger
```

## Usage 
This examples are available in the project at [messenger/examples](https://github.com/joaosoft/messenger/tree/master/examples)

> Server example
```go
func main() {
	m, err := messenger.NewMessenger()
	if err != nil {
		panic(err)
	}

	if err := m.Start(); err != nil {
		panic(err)
	}
}
```

> Client example
```
func main() {
	var user string
	flag.StringVar(&user, "listen", "", "The user identifier")
	flag.Parse()

	if user == "" {
		panic("invalid user identifier")
	}

	client, err := socket.NewClient()
	if err != nil {
		panic(err)
	}

	if err := client.Start(); err != nil {
		panic(err)
	}

	if err := client.Subscribe(user, "messenger::in-box"); err != nil {
		panic(err)
	}

	client.Listen(user, "messenger-in-box", func(message []byte) error {
		fmt.Printf("\nreceived on listener the message %s", string(message))
		return nil
	})

	// Unsubscribe
	//if err := client.Unsubscribe("topic_1", "channel_1"); err != nil {
	//	panic(err)
	//}

	client.Wait()
}
```

### How to run this?
1. go run main/main.go
2. go run main/client.go -listen user_one
3. go run main/client.go -listen user_two
4. Make the next http requests to send messages
   
    4.1. from user_one to user_two
    
    Method: ```POST``` 
    
    Route: ```http://localhost:8001/api/v1/messenger/message/users/user_two```
    
    Header: user: user_one
    
    Body:
    ```
    {
        "message": "hello my friend"
    }
    ```

    4.2 from user_two to user_one
    
    Method: ```POST``` 
    
    Route: ```http://localhost:8001/api/v1/messenger/message/users/user_one```
    
    Header: user: user_two
    
    Body:
    ```
    {
        "message": "hello my friend"
    }
    ```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
