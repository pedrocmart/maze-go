# maze-go

- Install
  * Go
  * Docker
  * docker-compose

- Building:
```bash
  docker compose build
  docker compose up
```
- API endpoints inside the container:
```bash
POST localhost:5000/v1/level
GET localhost:5000/v1/level
GET localhost:5000/v1/level/1
GET localhost:5000/v1/game/1
```

- Generate server:
    * swagger generate server -f ./doc/swagger.yaml -t api --exclude-main -A maze-go


## Tasks
- Task #1 - Creating Levels - Average time spent: ~3h
    - Initial config (swagger specs, APIs, Docker-compose)

- Task #2 - Validation - Average time spent: ~2h
    - Required validations:
        - Rectangular maps
        - Maps length
        - Map spaces [0-4]
    - Extra validations:
        - Map has one and only one start point
        - All the columns have the same size
        - Map has exit

- Task #3 - Minimum Survivable Path - Average time spent: ~5h
    - Here I tried to perform a BFS algorithm, traversing through the map, looking for the shortest path. Since we have more variables in our game, player's life and traps, I had to adapt to deduct player's life according to the trap. What I couldn't yet is to define what will be the path for the player: if he's going to survive or to find the shortest path. At the moment, the algorithm is always looking for the shortest one.
    
    - Complexity Time: O(M*N), where M is the number of arrays in the first dimension and N is the number of arrays in the second dimension.


- Extras:
    - Docker-compose
    - Unit Tests
    - APIs:
        - Get All Levels
        - Get Level by ID
        - Game by ID 

- Demo:
[![DEMO](https://img.youtube.com/vi/sCwWA4j0cp4/0.jpg)](https://youtu.be/sCwWA4j0cp4)
