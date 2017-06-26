# mgo-oid
A unique id generator that produces ids like mongodb object id. 

# ObjectId
Object id is a globally unique identifier for object in mongodb. It consists of 12 bytes, divided as follows:
```
 <table border="1">
     <caption>ObjectID layout</caption>
     <tr>
         <td>0</td><td>1</td><td>2</td><td>3</td><td>4</td><td>5</td><td>6</td><td>7</td><td>8</td><td>9</td><td>10</td><td>11</td>
     </tr>
     <tr>
         <td colspan="4">time</td><td colspan="3">machine</td> <td colspan="2">pid</td><td colspan="3">inc</td>
     </tr>
 </table>
``` 
