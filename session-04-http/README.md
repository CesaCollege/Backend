# Session 04 - HTTP

## HTTP Protocol

We need to talk about how the web works when describing HTTP (Hypertext Transfer Protocol).

- HTTP is an application layer protocol and It is TCP/IP-based protocol.
- HTTP allows the web-applications to communicate and exchange data.
- It is Stateless and connection-less.

<aside>
💡 The server does not recognize that you had requested before. Although we have some mechanism for maintaining the user session.

</aside>

### Client

The client is a computer that requests the server to get or change some data.

### Server

The server is a computer responsible for responding to the requests of the clients.

## How does the web work?

1. On the client side, the application uses TCP/IP to find and establish the target server.
2. After establishing a connection, the client sends a request message to the server.
3. The server calculates the response and responds to the request.

## HTTP Message

The structure of an HTTP message consists of

- Start line
- Headers
- Body

## HTTP Request

The HTTP request looks like

```
METHOD URI Version

Headers

Body
```

For example,

```
GET /products/index.html HTTP/1.0

Host: www.mysebsite.com
Accept: text/html
Accept-language: en-us
...
```

### Start line in a Request

- **Methods**: We have multiple methods like POST, GET, UPDATE, DELETE, PATCH, etc. We are going to talk about them
  later.
- **URI**: It is a set of characters that locate the resource.
- **************Version**************: It specifies the HTTP version.

### Headers in a Request

Headers contain some readable information. They are like some key values information that we can feed in our request.

## HTTP Response

The HTTP response looks like

```
Version StatusCode

Headers

Data
```

For example,

```
HTTP/1.0 200:OK

Host: www.mysebsite.com
Accept: text/html
Accept-language: en-us

<The content of HTML File>
```

<aside>
💡 Let’s use the TCP/IP Swiss army knife `nc` to see these data.

</aside>

You need to establish a connection first with `nc`

```bash
nc localhost 8081
```

<aside>
💡 You can monitor the HTTP request of a website using the network inspector of your browser. Let’s see this page [`https://www.digikala.com/search/notebook-netbook-ultrabook/?sort=1`](https://www.digikala.com/search/notebook-netbook-ultrabook/?sort=1)

</aside>

You can use some clients to call the APIs. One of them is `curl`

```bash
curl -X GET localhost:8081/api/v1/ping
```

<aside>
💡 Let’s see a preview of the authentication API

</aside>

## HTTP2

- There is no need to change the code.
- It is faster and more efficient.

## Why use HTTPS?

As it was mentioned, the HTTP request and response are in plain text and anybody is able to read an HTTP communication.
HTTP uses encryption to prevent this issue.

## REST

**[REST Countries API**:](https://restcountries.com/) It provides information about countries all around the world,
including details such as population, area, languages spoken, currencies used, and more.

```bash
curl https://restcountries.com/v3.1/all

curl https://restcountries.com/v3.1/alpha/US

curl https://restcountries.com/v3.1/currency/dollar
```

<aside>
💡 You can use Insomnia for having a graphical interface.

</aside>

## Make a Request in Go

[See this code](simple-http/main.go)

## Homework

Read the document [https://restcountries.com/](https://restcountries.com/) and write a code. The input of your code is a
currency and a region. Your code should print the capital of the countries in that region that have that currency.

```bash
# Input an example for your script
asia rial

# The output of the script
["Tehran", "Muscat", "Sana'a"]
```

- Errors should be handled properly
- The script must have two HTTP requests.
- These two requests should be in separate Goroutines.
- After Gathering the responses to 2 requests, find the common countries in both results

## Preview

One of the Homework or Class activities will be scraping and storing the price of taxi service. Here is
mine! [http://91.107.180.78:9191/](http://91.107.180.78:9191/)