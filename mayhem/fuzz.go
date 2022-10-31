package fuzz

import "strconv"
import "github.com/dave/jennifer/jen"


func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) > 1 {
        num, _ = strconv.Atoi(string(bytes[0]))

        switch num {
    
        case 0:
            content := string(bytes)
            jen.NewFile(content)
            return 0

        case 1:
            content := string(bytes)
            jen.NewFilePath(content)
            return 0

        case 2:
            content := string(bytes)
            var test jen.File
            test.HeaderComment(content)
            return 0

        case 3:
            content := string(bytes)
            var test jen.File
            test.Save(content)
            return 0

        default:
            jen.LitByte(bytes[0])
            return 0

        }
    }
    return 0
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}