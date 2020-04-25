package main
import ("github.com/gin-gonic/gin"
    "net/http"
    "time"
    "os"
    "strconv"
    )

const fileName = "data/eatTime.txt"

func main() {
	r := gin.Default()

    r.GET("/api/lastEatTime", func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin","*")
        c.JSON(http.StatusOK, gin.H{
            "timeStamp": loadLastTime(),
            "now": time.Now().Unix(),
        })
    })

    r.GET("/api/eat", func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin","*")
        err := writeLastTime()
        if err != nil { panic(err) }
        c.JSON(http.StatusOK, gin.H{})
    })
	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func loadLastTime() int64 {
    file, err := os.Open(fileName)
    if err != nil {
        if !os.IsExist(err) {
            return time.Now().Unix()
        }
        panic(err)
    }
    defer file.Close()

    stat, err := file.Stat()
    if err != nil {
        panic(err)
    }

    buffer := make([]byte, 10)
    n, err := file.ReadAt(buffer, stat.Size() - 10)
    if err != nil { panic(err) }
    t, err := strconv.ParseInt(string(buffer[:n]), 10, 64)
    if err != nil { panic(err) }
    return t
}

func writeLastTime() error {
    file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        panic(err)
    }
    defer file.Close()

    unixTimeString := strconv.FormatInt(time.Now().Unix(), 10)
    if _, err := file.Write([]byte(unixTimeString)); err != nil {
        return err
    }
    return nil
}
