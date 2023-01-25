![Frame 1](https://user-images.githubusercontent.com/65754609/213882280-367ceb95-600c-4998-b4cb-c1fd914a89b4.png)
![Frame 2](https://user-images.githubusercontent.com/65754609/213882292-b5c2f635-a86e-4f38-b88e-962679120926.png)

# Details about Project
Wiki: https://github.com/JewishLewish/PolygonDB/wiki

# Usage
Adjust databases/name-of-server/config.json
```json
{
    "path":"database.json",
    "key":"Better_Password"
}
```

Database.json example
```json
{
	"rows": [
		{
			"age": 5,
			"name": "A"
		},
		{
			"age": 20,
			"name": "B"
		},
		{
			"age": 30,
			"name": "C"
		}
	]
}
```
Python code for accessing the server
```python
import json
from websocket import create_connection


ws = create_connection("ws://localhost:8000/ws")

ws.send(json.dumps(
    {
        'password': 'Secret_Password', 
        'dbname': 'CatoDB',
        'location' :'rows.0.name',
        'action' : 'retrieve'
    }
))
print(json.loads(ws.recv()))  # "A"

ws.send(json.dumps(
    {
        'password': 'Secret_Password', 
        'dbname': 'CatoDB',
        'location' :'rows.0.age',
        'action' : 'record',
        'value':'5'
    }
))
print(json.loads(ws.recv())) # {Status: Success}

ws.send(json.dumps(
    {
        'password': 'Secret_Password', 
        'dbname': 'CatoDB',
        'location' :'rows',
        'action' : 'search',
        'value':'age:30'
    }
))
print(json.loads(ws.recv())) # {'Index': 2, 'Value': {'age': 30, 'name': 'C'}}
```

# Companies that uses PolygonDB

<div style="display: flex; justify-content: center;">
	<img src="https://discordapp.com/api/guilds/1024761808407498893/widget.png?style=banner2" alt="Discord Banner 2"/>
	<img src="https://discordapp.com/api/guilds/1046141941387116565/widget.png?style=banner2" alt="Discord Banner 2"/>
</div>

<img src="https://discordapp.com/api/guilds/692451473698586704/widget.png?style=banner2" alt="Discord Banner 2"/>

Sidenote, if you want your server to be on here then open a pull request. 

# Modules (Makes Polygon Easier to Use)
Javascript - https://github.com/NekaouMike/PloyConJS

Feel free to make your own module for Polygon, I will put your github on here. 
