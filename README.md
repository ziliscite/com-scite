### Rough Architecture
|Gateway| -> rate limit & single point of entry for the client
|Auth| -> authentication service to store user and authenticate through basic auth
|Token| -> to activate user once they are registered
|Mailer| -> will be called through rabbitmq asynchronously once user has registered (sending activation token) or activated their account

- Register
     create user   create    send token
|Gateway| <-> |Auth| -> |Token| --> |Mailer|
        respond

---

Second design 

- Register
     create user  create token
|Gateway| <-> |Auth| <-> |Token|
        respond | send token back to auth
                --> |Mailer|
        send token + user identity to mailer

---

- Login
    get and authenticate
|Gateway| <-> |Auth|
      return jwt

- Activate
    validate token  activate user
|Gateway| <-> |Token| <-> |Auth|
          resp   |   resp with user    
                 |--> |Mailer|
         send congrats email

### Final Say
These collections of microservices are quite coupled with each other; however, 
I'll make sure that |Auth| service will be quite alright to be used on its own if you decided
to not add "email activation" stuff
