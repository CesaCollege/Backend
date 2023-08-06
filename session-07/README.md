# API Design

This session consist of three sections. First we try to understand the HTTP handlers. Then, we are going to 
implement a very basic calculator. At the end we will implement a login and signup API.


# Create an HTTP Handler

## Handle

```go
type testStruct struct {
}

func (t testStruct) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello API"))
}

http.Handle("/test", &testStruct{})
```

## HandleFunc

```go
func HandlePingAPI(w http.ResponseWriter, r *http.Request) {
}

http.HandleFunc("/ping", HandlePingAPI)
```

# Response Writer

```go
// Access the header values and change them
w.Header().Add("Name", "Ali")

// Specifying the status code of the response
w.WriteHeader(http.StatusBadRequest)

// Writing the response value
_, _ = w.Write(data)
```

# HTTP Request

## Headers

To parse the request you can use `r.Header`

```go
// In this case, we look for the `Content-Lanugage` Header 
// and add a default value for it.
lang := r.Header.Get("Content-Language")
if lang == "" {
	lang = "en"
}
fmt.Println("language: ", lang)
```

## Method

```go
// In this code, we make sure that the request is called
// with Get method
fmt.Printf("%+v\n", r.Method)
if r.Method != http.MethodGet {
	w.WriteHeader(http.StatusMethodNotAllowed)
	return
}
```

## Query Params

```go
// Query Params
queryParams := r.URL.Query()
fmt.Printf("Query Params, %+v\n", queryParams)
```

## Body

```go
type signupRequest struct {
	Username    string `json:"username"`
	Name        string `json:"name"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	Gender      string `json:"gender"`
}

// Body
reqData, _ := io.ReadAll(r.Body)
var sr signupRequest
err := json.Unmarshal(reqData, &sr)
if err != nil {
	log.Println(err.Error())
}
fmt.Printf("%+v\n", sr)// Body
reqData, _ := io.ReadAll(r.Body)
var sr signupRequest
err := json.Unmarshal(reqData, &sr)
if err != nil {
	log.Println(err.Error())
}
fmt.Printf("%+v\n", sr)
```

# Class Activity

Now let’s create two APIs to perform a calculation.

## The Calculate API

The URL should be `/calculate?a=&b=`. The JSON response represents the `a <Operator> b`. You need to parse `Operator` header to find out what is the operation.

Here is an example,

```bash
# Request
curl -H 'Operator:-' -XGET localhost:8080/calculate?a=32&b=16

# Response
{
  "answer": 16
}
```

## The Average API

The URL should be `/average`. The array is passed in JSON format. The JSON response represents the average of these numbers.

Here is an example,

```bash
# Request
curl --data '{"numbers": [12.321, 5.33, 519.239, 32.44, 31]}' -XGET localhost:8080/average

# Response
{
  "answer": 
}
```