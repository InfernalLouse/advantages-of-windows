package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "2.6" )

func AwcbL5GgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 274
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func JjQIvMrMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dQZ6xCQPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RKGrpDWWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ocSyWIXAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 52
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D7fmUYVRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func T1OHVZOlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func X925s963Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VBIDrXQ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qQy6HIlzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZOhc9hrZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OQqm4i3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 120
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jN2YWlFTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 189
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BmjITtJLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gW4Li8qmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func y8pkVBXWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 11jBd5b0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P3JfJj2AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oew7lTf6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 108
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CXlN3BiTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xgCbeKMrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uHd9KxtwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8actbOQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ncFMCWU7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 107
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func WMUDArq5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lEb1DxVXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 31
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l0TCcmEtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 54
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xuRTtEIKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 24
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DG7gdy7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jdBI0bZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NEOvVK8VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 131
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ync4ZulgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 173
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func baPhhFq6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LhHrTJhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 105
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func czQva7UQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8oe55H32Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 297
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OlKFCOAZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 157
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Dn8hlkMZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TUVyByzZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GH6i740DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 139
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n51S8Q3lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func eAP9OcncWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 220
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func FVKGZ2lrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AJuPd7LqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 184
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kHuj6x2jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9ga20yw9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 104
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func S6WVAkTdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 215
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cLAmd1FuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 172
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b02OEmrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EH7Wx5dbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 73
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fj65NeFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 26
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GttrPljGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 264
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K6Xlg70nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 272
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TX1GzFilWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HSpZmMKxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 249
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zNGgV7LBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MhSfmBoUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rMa27oSYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OTPXvk5sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Qa1xJ8WbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lAzJPeAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lIePC81OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 07Vs4tDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 216
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Fd85XzyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Nh7Gy5kfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GGPLSRxlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BWbrcJEqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PSqAGpJyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 56
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XC7sdqGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func udot3svKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func reQhe5oRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x2lqrp0gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 69
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oFwdCg51Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 130
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pC4cBsJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HSbKjFUQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nOyGBdYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 299
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w0wkD93iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7iGmlUcVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 288
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 8HIcnWK2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gG2AuuwEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 196
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dXcbJKpPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 148
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hMpqVl6dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iTPznlINWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 14
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fSEgQmMRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func g75pExhCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 262
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CohQi9jaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func W9tqRmYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 232
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lWjdLyjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bKSc38FHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 278
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zEK3rrz6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x5PByGqzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yZjktAKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LyagfFCGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kKKuj658Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 65hJeT2JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 40
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xOE6C2j7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 121
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gLxuaaJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 211
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 007ilEAeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 58
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func c3JU05xpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 185
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func M0gNReVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 281
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yfpa5WhVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func IfzArMAgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 11
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func irXMXPbPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 233
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func rfA1QaJ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 128
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SazovPJwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 287
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ChIiTOFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7IbZm7KeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 253
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func s3ky3gLgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RLq5Tp9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 199
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GVO5xpFyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 268
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5yLUdkvsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func D0DUvSGZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CNRSiKx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 241
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func U4mWSnG7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vdXSmatiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 227
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6H2pxpShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PJOOtPRWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Sqmfbo77Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func qUGgB03YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ucQf9iR4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 152
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6JRoc82zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 214
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vqM62UgyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 222
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 53xYnXqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 169
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VlRO7L9GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 226
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jAnmBB5NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func x7f0KbKtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 198
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2gOc4qTpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func BJSdkNjmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mOeoHA0BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 217
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h2DFcNQaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 159
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func l0OQJc2kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 291
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func q4otZBAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func sNHuI4GWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 122
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ubN37CIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 126
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Ftg3BOX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func O4WazNptWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func e0kmw0Z2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 193
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9RxFCFtSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 129
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func N0jiaoVSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fvizjnrqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gUX4FePOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 90
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4Xhp13UgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 19
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9DJG72ZxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NhV4Bi23Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pYG1k91jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jImJx7IUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func aZF6iRagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Wg1v2yCLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 208
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TGdAt1uKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 153
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func bpB6eN71Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7mXCfrtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2gwvz2DwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 82
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3fNID3I8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 175
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LdLOqa3HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 76
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fCYy8AsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vzj69jdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 103
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tFFRMBKbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 210
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func r7fjR21vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func kNsJC03fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func GBuoOXR1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 260
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mslm8ScIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 75
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Lz7kJOOwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 83
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pt8Bgr4XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 298
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5KA3N3SDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 236
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5tTnGE4kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 282
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yx14uU9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QQd8SveqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 194
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nzJBWG6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 64
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SxU3AzuUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0A8uK7kYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 256
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HDTTRRqsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 242
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func iuqCrg8nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 284
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func o1FJRGalWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 180
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RN5lMwDzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ugK6WaREWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6K3LyXERWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 224
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m88pzNQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func I2yxQKyiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AN53wbRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 44
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ajBTUuXwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 243
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xpc4fay9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 167
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vGJEU5D7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7Otxt8QlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pTUosTSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0QkG4SgKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 6rcCYqRUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 166
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 17ZO32gcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZPQbBSHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 127
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hlk0mG6ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EoyHsUVMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 218
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Euin0zOeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 271
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EyPXBmo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 149
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yPmdf5y0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 100
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L8Pr1vLPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 96
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func We2OnF9uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nasMTR7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 101
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pGlMD9KiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 23
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7CdhmZIjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XZcfYuB6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 102
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hohjDdb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 161
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 71l90oZIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EEBqr0vQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 41
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tUsCh9ZaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PZMjkgm6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 43
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7rhqTYvWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dDfFHXQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 162
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Zk54kwiDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vc4GWa7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 195
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MBEe7hf0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OkLPKQXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 168
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TEJby93GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 1EQj5fo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func hPEcVYHFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 221
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0X1QN5PwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 48
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func t2bQIzj7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 10
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b48rSragWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 47
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func DCGMU42UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 86
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHEHqRnPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HQdJmj97Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 134
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ee2fezvGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 250
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func CQGnyHPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 45
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RsLnHBH4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AZ8gVSg0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 254
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func HRo709I6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 124
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func w89qpgW5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 138
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func oY3rk4BRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 202
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func xZdYeofuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 147
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func EHyAHpJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func m1i4KLYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 225
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L57mqi7NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 94
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func P8Q0tJKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YlI47VEoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2WgnrwljWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 144
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func B25Qje7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 170
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func K3lvGvNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 203
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ATJldH8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 33
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mJqOYa04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3b6QzynEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 295
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UU7YZR5MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 296
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func MyOhfBhYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 212
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func E2YBvoPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Vc6UU4caWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 30
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func pytySZ4oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 294
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RM2Kc4SwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 66
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func weZlnwGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 276
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func VjP3w689Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 247
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func yLNBJTNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 93
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RQwhOQIhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 80
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7jdXgnZfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 140
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func cbMnIRVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 141
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nSkzxXk4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func czmmLjDKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func fLPsjNRiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 238
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 06gckMe2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3DSA5IrnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 207
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func C7dLpFpIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 62
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func n0vfQb0EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 133
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 5pJllSDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 60
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func k2UK2Z8MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 237
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func RWXAXOonWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 16
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func L49hrL2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 177
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func XDnlWuvxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 34
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9zB8ya9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 176
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mKqIrQS3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 259
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func QYuiDt1DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func auT3IFI5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 240
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ZCWrSN84Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 143
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func SPenlODPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 187
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OfxFPKDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 46
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 0eQlWJIHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 245
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func UrxuHiBqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 106
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func b6f7c22CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 117
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 85aMwT91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 154
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LbevwHz8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 155
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func KhuDb0OdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func jTE5Y4SoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 28
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func tALInpEfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func nWGFuHhQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 113
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 9EQGXFSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 114
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func YRzMlxd4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 279
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 4xGNe8xdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 77
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vNooh2sPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 160
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func OadGgVdOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func mHqDXMLHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 229
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func AEmAxjyHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 132
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ORxGGJHxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 112
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func gIzqdDgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 53
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func dbRnkCIEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 37
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 3sURagHdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 164
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 7cGPajUEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 85
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vE6JIVbrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 190
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func vXDkHrqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 156
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func h0y5o9qjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 286
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func flpStCFJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 275
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Oan6z0kHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 42
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ezFuD6FNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func Hnw8BBUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 55
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func LB0fMlDBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 12
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ksij7SAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 15
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func z3xpLCg7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 21
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func uK81jnMOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 252
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func zFO8kXWVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 67
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 12NiAKcdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 68
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func TBdyjcN2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 98
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func lfePGyMQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 171
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func NxajZsejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 174
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func PfO4lhdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 290
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func ehMVcTbbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 35
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

func 2vu0VAdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 36
        fmt.Printf("worker %d → %d\n", id, result)
    }
}

