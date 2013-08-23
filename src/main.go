package main
import(
    "fmt"
    "net/http"
    "log"
    "controller"
)

type MethodMux struct {
    methodMap map[string]func(http.ResponseWriter, *http.Request)
}

func (mm *methodMux) ServeHTTP(response http.ResponseWriter, request *http.Request) {
    if function, ok := mm.methods[request.Method]; ok {
        function(response, request);
    } else {
        http.NotFoundHandler().ServeHTTP(response, request)
    }
}

func main() {
    http.HandleFunc("/user", methodMux{
        map[string]func(http.ResponseWriter, *http.Request) {
            "POST" userAdd
        }
    })
    http.HandleFunc("/user/", Add)
    s := &http.Server {
        Addr: ":8080",
        ReadTimeout: 30 * time.Second,
        WriteTimeout: 30 * time.Second,
        MaxHeaderBytes: 1<<20,
    }
    log.Fatal(s.ListenAndServe())
}
