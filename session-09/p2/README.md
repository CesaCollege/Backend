# practice for docker-compose (nginx-golang-postgres)

1. write a docker file for golang app in the api folder
2. write a docker-compose.yml with these specifications :
    1. a service called api for golang app
    2. a service called db for postgres database (user = geek , password=1234 , dbname = temp)
    3. a service for nginx including config file in nginx folder as default config file (/etc/nginx/conf.d/default.conf)

<aside>
💡 don't forget to define volumes and networks

</aside>

<aside>
💡 hint → you can use bind mount for nginx configuration file

</aside>