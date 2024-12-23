# lem-in

This project is a simulation of ant movement through a network of rooms connected by tunnels, where the goal is to find the quickest route from the start to the end room for a given number of ants. The program ensures efficient pathfinding while handling errors and displaying the movement of ants in an organized manner.

for more details check this [Link](https://github.com/01-edu/public/tree/master/subjects/lem-in)


## Usage

### Installation 

1. Clone the repository:
   ```bash
     git clone https://github.com/heeemzaaa/lem-in
   
2. Navigate to the project directory:
   ```bash
     cd lem-in/

### How to use

Run the program with an extra argumrent contains a file like this 

 `go run main.go <filename.txt>`

### Example 

- Run program:
    ```bash
        go run main.go test.txt
        5
        ##start
        start 2 2
        r1 3 1
        r2 3 3
        r3 4 1
        ##end
        end 5 2

        start-r1
        start-r2
        r1-r3
        r2-end
        r3-end


        L2-r1 L1-r2 
        L2-r3 L4-r1 L1-end L3-r2 
        L2-end L4-r3 L3-end L5-r2 
        L4-end L5-end 

- Get steps count:
    ```bash
        go run main.go test.txt | grep '^L' | wc -l

## Graphs in file

- The file looks like this 
    ```bash
        5
        ##start
        start 2 2
        r1 3 1
        r2 3 3
        r3 4 1
        ##end
        end 5 2

        start-r1
        start-r2
        r1-r3
        r2-end
        r3-end

## How program works

The program simulates an ant farm, where ants move from a start room to an end room through various paths. It reads the farm's configuration from a file, extracts room links, and identifies bad rooms with excessive links. The program finds all valid paths from the start to the end room using BFS, filters out invalid paths, and assigns ants to the shortest paths. Ants are then moved along these paths, ensuring no room conflicts and simulating their progress until all ants reach the end room.


## Authors

- **Hassan El Ouazizi**     [GitHub Profile](https://github.com/helouazizi)
- **Hamza Elkhawlani**      [GitHub Profile](https://github.com/heeemzaaa)
- **Saad Salah Tounsadi**   [GitHub Profile](https://github.com/0Emperor)
- **Ismail Haji**           [GitHub Profile](https://github.com/hajji-Ismail)



