package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, HTML_PAGE)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":47687", nil)
}

const (
    HTML_PAGE = `
<!DOCTYPE html>
<html>

<h1 style="color:chartreuse"> <big><big><big> Welcome to Poseidon </big></big></big> </h1>
<body style="background-color:black;">

<!-- rabbit -->
<p>
<a href="http://localhost:15672">
<img src="http://sciencenordic.com/sites/default/files/imagecache/620x/rabbit_0.jpg" alt="rabbit" width="400" height="267" border="0">
</a>
</p>

<!-- mongo -->
<p>
<a href="http://localhost:28017">
<img src="http://www.todayifoundout.com/wp-content/uploads/2011/08/honey-badger.jpg" alt="mongo" width="400" height="267" border="0">
</a>
</p>

<!-- docs -->
<p>
<a href="http://localhost:2222">
<img src="https://upload.wikimedia.org/wikipedia/commons/3/35/Portsmouth_Historic_Dockyard.jpg" alt="docs" width="400" height="267" border="0">
</a>
</p>

<!-- monitor -->
<p>
<a href="http://localhost:1111">
<img src="http://mylittlepony.hasbro.com/images/spring2016/ponies/char_pinkiepie.png" alt="monitor" width="400" height="267" border="0">
</a>
</p>

<!-- switch -->
<p>
<a href="http://localhost:80">
<img src="https://farm2.static.flickr.com/1152/4731530618_29eec62bf1_b.jpg" alt="monitor" width="400" height="267" border="0">
</a>
</p>

</body>
</html>
`
)