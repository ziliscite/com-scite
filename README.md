## Introduction
This project will be dedicated my 19-year-old self

## Architecture
Rough explanation of what each service should be doing

- |Comic|
Host online comic information metadata, such as title, author, ratings, popularity, etc. 
cover url will contain the link to an image file from an object storage.

- |Chapter|
Will host a CDN to each chapter of a comic. It will contain the relation between chapters 
and comic as well as the pages (another url to an object storage)

### Restriction
For get operations, it's a free game, everyone can read anything for free

For stuff such as giving a rating, adding comic to a bookmark, etc, however, can only be done by an authenticated user

For things like CREATE, UPDATE, and DELETE, it can only be done by an authorized user (admin)

## Data
The database schema. Won't be too complex since the learning objective is to figure out 
how microservices, event-driven, as well as, AWS (S3, CloudFront, and Route 53) work in tandem. 
I'll probably not gonna deploy it whatsoever, perhaps later, after I figure out kubernetes.

### Comic
|Comic Table|  
-> |Popularity Table| (one-to-one) -> A very heavy update operations  
-> |Ratings Table| (many-to-many) -> Update heavy operations ?
-> |User Bookmark Join Table| (many-to-many) ? Questionable, both might be another service altogether
-> |Cover Url Table| (one-to-one) ? Or just embed url into the main table? Or is it going to be some kind of cover versioning (one-to-many)?
-> |Status Enum Table| I guess the instance of this table can be embedded into the comic table
-> |Type Enum Table| This as well

|Genre Table|
-> |Comic-Genre Join Table| (many-to-many)

### Chapter
|Chapter Table|
-> |Chapter Page Table| (one-to-many) -> A chapter consists of multiple pages

## Sketch

Since chapter didn't hold reference to comic, to maintain integrity...

Event-Driven:  
Producer (Comic Service):

When a comic is created or updated, the Comic service sends an event (e.g., ComicCreated or ComicUpdated) to a message queue.

Consumer (Chapter Service):  
The Chapter service listens for the ComicCreated or ComicUpdated events. Upon receiving the event, the Chapter service updates its local data or performs necessary actions based on the comic_id.
This method would ensure that the Chapter service stays in sync with the Comic service without needing a direct foreign key relationship.

