package cryptoutil

import (
    "crypto/md5"
    "fmt"
)

func Md5Encode(str string) string {
    data := []byte(str)
    /*
       这里的Sum()是对hash.Hash对象内部存储的内容进行校验和计算,
       然后将其追加到data的后面形成一个新的byte切片.
    */
    encodingStr := fmt.Sprintf("%x\n", md5.Sum(data))
    //fmt.Println(encodingStr)
    return encodingStr

    /*
       也可以初始化一个MD5对象.
       这里实际返回的是hash.Hash对象, 函数原型为 func New() hash.Hash,
       该对象md5Ctx的Sum()实现了hash.Hash的Sum()接口
       因为md5Ctx已经写入了data的内容, 因此Sum()需要将参数置为nil
    */
    //md5Ctx := md5.New()
    //md5Ctx.Write(data)
    //cipherStr := md5Ctx.Sum(nil)
    //// 除了在Printf中用%x输出16进制字符串, 也可以用hex的方法输出
    //encodingStr2 := fmt.Sprint(hex.EncodeToString(cipherStr))
    //fmt.Println(encodingStr2)

}
