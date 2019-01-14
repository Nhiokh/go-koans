title: Go(lang) 101
author:
  name: Florent Baldino & Samuel Abiassi
  twitter: SamAbiassi
output: basic.html
controls: true

--

# Go(lang) 101
## Un langage simple, flexible et performant

--

#### 1 - Historique et utilisateurs
#### 2 - Le langage
#### 3 - Les outils
#### 4 - Parallélisme et concurrence

--

### Inventé chez Google (2012) par :
* **Ken Thompson** - UNIX, UTF-8, Plan 9, B
* **Rob Pike** - UNIX, UTF-8, Plan 9, Inferno, Newsqueak
* **Robert Griesemer** - Object Oberon

--

### Pourquoi un nouveau langage ?
* API toujours plus complexes
* Compilations trop lentes
* Évolution vers le multi-cœurs

--

### Qui utilise Go ?
* Google, Youtube, Twitter...
* Docker, Kubernetes...
* Algolia, Leboncoin, MolotovTv...
* https://github.com/golang/go/wiki/GoUsers

--

# Le langage

--

### Syntaxe similaire à Typescript

```go
package main

import "fmt"

func add(x int, y int) int {
  return x + y
}

func main() {
  var res = add(6,2)
  fmt.Println(res)
}
```

--

### Librairie standard conséquente

--

### Pas d'héritage, on compose

--

### Code organisé par package

--

### Gestion des erreurs

--

# Les outils

--

### go fmt

--

### go get

--

### go doc

--

### go test

--

### go run

--

### go build

--

# Parallélisme et concurrence

--
### Simple tâche

![Simple tâche](https://i.imgur.com/DLJ2KZo.jpg)

--

### Modèle parallèle

![Parallèle](https://i.imgur.com/qyg029I.jpg)

--

### Modèle concurrent

![Modèle concurrent](https://imgur.com/rlVtj9I.jpg)

--

### Modèle concurrent parallélisé

![Concurrent et parallèle](https://imgur.com/gAyELMb.jpg)

--

### Goroutines

`myFunc ()`

devient asynchrone en ajoutant simplement

`go myFunc()`

--

### Channels

Communication entre les goroutines via pointeurs (allocations mémoires)

`ch := make(chan int)`

Deux actions possible :
* `send: ch <- val`
* `receive: x = <-ch`

--

### Exemple de pipeline

Counter (0,1,2,3...) -> Squarer (0,1,4,6...) -> Printer

--

### Plus concret ?

Handler -> Broadcaster -> Logger

Modèle parallélisable !
