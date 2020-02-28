# Assignment2

This is an API that aggregates information from https://git.gvk.idi.ntnu.no. This includes two parts, the development of an API for direct invocation, 
as well as an interface for the registration of Webhooks for invocation upon certain events.

For this assignment i have collaborated with other students from class, but I have implemented everything by myself. I felt like I learned a lot by doing that. 

## Endpoints

All the endpoints are added to http://10.212.137.62:8080

## Webhooks

syntax: /repocheck/v1/webhooks

### POST 

Make sure postman is on POST.

The POST request should happen on: 

http://10.212.137.62:8080/repocheck/v1/webhooks/ 

You have to use postman and write json in the Body part like this: 

{
  "event": "commits|languages|status",
  "url": "webhook url"
}

Use either commits, languages or status, and write your url to the thirdparty service that will recieve the notification when invoked. 

Now you should have registered a webhook. 

### GET

Make sure postman is on GET.
     
The GET request should happen on: 

http://10.212.137.62:8080/repocheck/v1/webhooks/

You can also retrieve only one webhook by copying the id from one of the ones you got from the get request (the id's are a bit complex): 

http://10.212.137.62:8080/repocheck/v1/webhooks/{webhookID}

### DELETE
Make sure postman is on DELETE.

The DELETE request should happen on: http://10.212.137.62:8080/repocheck/v1/webhooks/{webhookID}

The webhookID is the id of the webhook you want to delete. The id is easiest to retrieve by getting all webhooks and copy one of the id's.

## Commits

syntax: /repocheck/v1/commits{?limit=[0-9]+{&auth=<access-token>}}

NB! : Please be patient when doing the requests. It takes time to loop through all projects, so wait atleast a minute for each request before you exit. 

### GET

You can send a get request to:

- http://10.212.137.62:8080/repocheck/v1/commits/ 

It is important that you have the / after commits. This will return 5 repositories by default.

If you want to limit the results you send a request to: 

- http://10.212.137.62:8080/repocheck/v1/commits?limit=100

This will return 100 repositories for example, but you can use any number. 

You can also add your own token to the url to get all the repositories you have access to:

- http://10.212.137.62:8080/repocheck/v1/commits?limit=100&auth=YOUR_TOKEN

if you want to use limit as well

or:

- http://10.212.137.62:8080/repocheck/v1/commits?auth=YOUR_TOKEN

if you only want to use your token, and get the default 5 results. 

## Languages

NB! : Please be patient when doing the requests. It takes time to loop through all projects, so wait atleast a minute for each request before you exit.

### POST

syntax: /repocheck/v1/languages/ or /repocheck/v1/languages/{&auth=<access-token>}

NB! : Notice that you cannot query on limit here, it will return all langauges in this case (implementation choice). 

The post request should happen on: 

http://10.212.137.62:8080/repocheck/v1/languages/ 

or

http://10.212.137.62:8080/repocheck/v1/languages/?auth=YOUR_TOKEN

You have to use postman and send an array of the ID's of the projects, for example (these should be valid, but if you know specific project ID's, use those):  

- [605, 590, 589] 

in this format. (This is done in body in postman)

### GET

syntax: /repocheck/v1/languages{?limit=[0-9]+{&auth=<access-token>}}

You can send a get request to:

- http://10.212.137.62:8080/repocheck/v1/languages/ 

It is important that you have the / after languages. This will return 5 languages by default.

If you want to limit the results you send a request to: 

- http://10.212.137.62:8080/repocheck/v1/languages?limit=20

This will return 20 languages for example, but you can use any number. 

You can also add your own token to the url to get all languages from repositories you have access to:

- http://10.212.137.62:8080/repocheck/v1/languages?auth=YOUR_TOKEN

or:

- http://10.212.137.62:8080/repocheck/v1/languages?limit=20&auth=YOUR_TOKEN

 if you want to use limit as well.

## Status

syntax: /repocheck/v1/status

### GET

You can send a get request to:

- http://10.212.137.62:8080/repocheck/v1/status/

to get the health indication from the gitlab webiste and database, and the runtime. 

