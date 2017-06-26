# mgo-oid
A unique id generator that produces ids like mongodb object id. 

# ObjectId
Object id is a globally unique identifier for object in mongodb. It consists of 12 bytes, divided as follows:
* a 4-byte value representing the seconds since the Unix epoch
* a 3-byte machine identifier
* a 2-byte process id
* a 3-byte counter, starting with a random value
