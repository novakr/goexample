/*
引用：
http://www.oschina.net/code/snippet_173630_12006

编译--》运行--》http://127.0.0.1:8080/

这个验证码包含下面三层元素
•随机大小和颜色的10个点
•4位数字的验证码（随机偏转方向、每个点间距随机）
•一条类似删除线的干扰线

对应的带代码注释的源文件如下：
*/
 package main
 import (
     crand "crypto/rand"
     "fmt"
     "image"
     "image/color"
     "image/png"
     "io"
     "math/rand"
     "net/http"
     "strconv"
     "time"
 )
 const (
     stdWidth  = 100 // 固定图片宽度
     stdHeight = 40  // 固定图片高度
     maxSkew   = 2
 )
 // 字体常量信息
 const (
     fontWidth  = 5 // 字体的宽度
     fontHeight = 8 // 字体的高度
     blackChar  = 1
 )
 // 简化期间使用的字体库
 var font = [][]byte{
     { // 0
         0, 1, 1, 1, 0,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         0, 1, 1, 1, 0},
     { // 1
         0, 0, 1, 0, 0,
         0, 1, 1, 0, 0,
         1, 0, 1, 0, 0,
         0, 0, 1, 0, 0,
         0, 0, 1, 0, 0,
         0, 0, 1, 0, 0,
         0, 0, 1, 0, 0,
         1, 1, 1, 1, 1},
     { // 2
         0, 1, 1, 1, 0,
         1, 0, 0, 0, 1,
         0, 0, 0, 0, 1,
         0, 0, 0, 1, 1,
         0, 1, 1, 0, 0,
         1, 0, 0, 0, 0,
         1, 0, 0, 0, 0,
         1, 1, 1, 1, 1},
     { // 3
         1, 1, 1, 1, 0,
         0, 0, 0, 0, 1,
         0, 0, 0, 1, 0,
         0, 1, 1, 1, 0,
         0, 0, 0, 1, 0,
         0, 0, 0, 0, 1,
         0, 0, 0, 0, 1,
         1, 1, 1, 1, 0},
     { // 4
         1, 0, 0, 1, 0,
         1, 0, 0, 1, 0,
         1, 0, 0, 1, 0,
         1, 0, 0, 1, 0,
         1, 1, 1, 1, 1,
         0, 0, 0, 1, 0,
         0, 0, 0, 1, 0,
         0, 0, 0, 1, 0},
     { // 5
         1, 1, 1, 1, 1,
         1, 0, 0, 0, 0,
         1, 0, 0, 0, 0,
         1, 1, 1, 1, 0,
         0, 0, 0, 0, 1,
         0, 0, 0, 0, 1,
         0, 0, 0, 0, 1,
         1, 1, 1, 1, 0},
     { // 6
         0, 0, 1, 1, 1,
         0, 1, 0, 0, 0,
         1, 0, 0, 0, 0,
         1, 1, 1, 1, 0,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         0, 1, 1, 1, 0},
     { // 7
         1, 1, 1, 1, 1,
         0, 0, 0, 0, 1,
         0, 0, 0, 0, 1,
         0, 0, 0, 1, 0,
         0, 0, 1, 0, 0,
         0, 1, 0, 0, 0,
         0, 1, 0, 0, 0,
         0, 1, 0, 0, 0},
     { // 8
         0, 1, 1, 1, 0,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         0, 1, 1, 1, 0,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         0, 1, 1, 1, 0},
     { // 9
         0, 1, 1, 1, 0,
         1, 0, 0, 0, 1,
         1, 0, 0, 0, 1,
         1, 1, 0, 0, 1,
         0, 1, 1, 1, 1,
         0, 0, 0, 0, 1,
         0, 0, 0, 0, 1,
         1, 1, 1, 1, 0},
 }
 type Image struct {
     *image.NRGBA
     color   *color.NRGBA
     width   int //a digit width
     height  int //a digit height
     dotsize int
 }
 func init() {
     // 打乱随机种子
     rand.Seed(int64(time.Second))
 }
 // 产生指定长宽的图片
 func NewImage(digits []byte, width, height int) *Image {
     img := new(Image)
     r := image.Rect(img.width, img.height, stdWidth, stdHeight)
     img.NRGBA = image.NewNRGBA(r)
     img.color = &color.NRGBA{
         uint8(rand.Intn(129)),
         uint8(rand.Intn(129)),
         uint8(rand.Intn(129)),
         0xFF}
     img.calculateSizes(width, height, len(digits))
     // Draw background (10 random circles of random brightness)
     // 画背景， 10个随机亮度的园
     img.fillWithCircles(10, img.dotsize)
     maxx := width - (img.width+img.dotsize)*len(digits) - img.dotsize
     maxy := height - img.height - img.dotsize*2
     x := rnd(img.dotsize*2, maxx)
     y := rnd(img.dotsize*2, maxy)
     // Draw digits. 画验证码
     for _, n := range digits {
         img.drawDigit(font[n], x, y)
         x += img.width + img.dotsize // 下一个验证码字符的起始位置
     }
     // Draw strike-through line. 画类似删除线的干扰线
     img.strikeThrough()
     return img
 }
 func (img *Image) WriteTo(w io.Writer) (int64, error) {
     return 0, png.Encode(w, img)
 }
 // 计算几个要显示字符的尺寸，没有开始绘画。
 func (img *Image) calculateSizes(width, height, ncount int) {
     // Goal: fit all digits inside the image.
     var border int // 边距
     if width > height {
         border = height / 5
     } else {
         border = width / 5
     }
     // Convert everything to floats for calculations.
     w := float64(width - border*2)  //  100 - 8*2=84
     h := float64(height - border*2) //  40 - 8*2 = 24
     fmt.Println("ddd%v;%v", w, h)
     // fw takes into account 1-dot spacing between digits.
     fw := float64(fontWidth) + 1 // 6
     fh := float64(fontHeight)    // 8
     nc := float64(ncount)        // 4
     fmt.Println("eee%v;%v;%v", fw, fh, nc)
     // Calculate the width of a single digit taking into account only the
     // width of the image.
     nw := w / nc //  84/ 4 = 21
     // Calculate the height of a digit from this width.
     nh := nw * fh / fw //  21*8/6 = 28
     // Digit too high?
     if nh > h {
         // Fit digits based on height.
         nh = h            // nh = 24
         nw = fw / fh * nh // 6 / 8 * 24 = 18
     }
     // Calculate dot size.
     img.dotsize = int(nh / fh) // 24/8 = 3
     // Save everything, making the actual width smaller by 1 dot to account
     // for spacing between digits.
     img.width = int(nw)                // 18
     img.height = int(nh) - img.dotsize // 24 - 3 = 21
     fmt.Printf("format:%v;%v;%v/r/n", img.dotsize, img.width, img.height)
 }
 // 随机画指定个数个圆点
 func (img *Image) fillWithCircles(n, maxradius int) {
     color := img.color
     maxx := img.Bounds().Max.X
     maxy := img.Bounds().Max.Y
     for i := 0; i < n; i++ {
         setRandomBrightness(color, 255) // 随机颜色亮度
         r := rnd(1, maxradius)          // 随机大小
         img.drawCircle(color, rnd(r, maxx-r), rnd(r, maxy-r), r)
     }
 }
 // 画 水平线
 func (img *Image) drawHorizLine(color color.Color, fromX, toX, y int) {
     // 遍历画每个点
     for x := fromX; x <= toX; x++ {
         img.Set(x, y, color)
     }
 }
 // 画指定颜色的实心圆
 func (img *Image) drawCircle(color color.Color, x, y, radius int) {
     f := 1 - radius
     dfx := 1
     dfy := -2 * radius
     xx := 0
     yy := radius
     img.Set(x, y+radius, color)
     img.Set(x, y-radius, color)
     img.drawHorizLine(color, x-radius, x+radius, y)
     for xx < yy {
         if f >= 0 {
             yy--
             dfy += 2
             f += dfy
         }
         xx++
         dfx += 2
         f += dfx
         img.drawHorizLine(color, x-xx, x+xx, y+yy)
         img.drawHorizLine(color, x-xx, x+xx, y-yy)
         img.drawHorizLine(color, x-yy, x+yy, y+xx)
         img.drawHorizLine(color, x-yy, x+yy, y-xx)
     }
 }
 // 画一个随机干扰线
 func (img *Image) strikeThrough() {
     r := 0
     maxx := img.Bounds().Max.X
     maxy := img.Bounds().Max.Y
     y := rnd(maxy/3, maxy-maxy/3)
     for x := 0; x < maxx; x += r {
         r = rnd(1, img.dotsize/3)
         y += rnd(-img.dotsize/2, img.dotsize/2)
         if y <= 0 || y >= maxy {
             y = rnd(maxy/3, maxy-maxy/3)
         }
         img.drawCircle(img.color, x, y, r)
     }
 }
 // 画指定的验证码其中一个字符
 func (img *Image) drawDigit(digit []byte, x, y int) {
     // 随机偏转方向
     skf := rand.Float64() * float64(rnd(-maxSkew, maxSkew))
     xs := float64(x)
     minr := img.dotsize / 2               // minumum radius
     maxr := img.dotsize/2 + img.dotsize/4 // maximum radius
     y += rnd(-minr, minr)
     for yy := 0; yy < fontHeight; yy++ {
         for xx := 0; xx < fontWidth; xx++ {
             if digit[yy*fontWidth+xx] != blackChar {
                 continue
             }
             // Introduce random variations.
             // 引入一些随机变化，不过这里变化量非常小
             or := rnd(minr, maxr)
             ox := x + (xx * img.dotsize) + rnd(0, or/2)
             oy := y + (yy * img.dotsize) + rnd(0, or/2)
             img.drawCircle(img.color, ox, oy, or)
         }
         xs += skf
         x = int(xs)
     }
 }
 // 设置随机颜色亮度
 func setRandomBrightness(c *color.NRGBA, max uint8) {
     minc := min3(c.R, c.G, c.B)
     maxc := max3(c.R, c.G, c.B)
     if maxc > max {
         return
     }
     n := rand.Intn(int(max-maxc)) - int(minc)
     c.R = uint8(int(c.R) + n)
     c.G = uint8(int(c.G) + n)
     c.B = uint8(int(c.B) + n)
 }
 // 三个数中的最小数
 func min3(x, y, z uint8) (o uint8) {
     o = x
     if y < o {
         o = y
     }
     if z < o {
         o = z
     }
     return
 }
 // 三个数的最大数
 func max3(x, y, z uint8) (o uint8) {
     o = x
     if y > o {
         o = y
     }
     if z > o {
         o = z
     }
     return
 }
 // rnd returns a random number in range [from, to].
 // 然会指定范围的随机数
 func rnd(from, to int) int {
     //println(to+1-from)
     return rand.Intn(to+1-from) + from
 }
 const (
     // Standard length of uniuri string to achive ~95 bits of entropy.
     StdLen = 16
     // Length of uniurl string to achive ~119 bits of entropy, closest
     // to what can be losslessly converted to UUIDv4 (122 bits).
     UUIDLen = 20
 )
 // Standard characters allowed in uniuri string.
 // 验证码中标准的字符 大小写与数字
 var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
 // New returns a new random string of the standard length, consisting of
 // standard characters.
 func New() string {
     return NewLenChars(StdLen, StdChars)
 }
 // NewLen returns a new random string of the provided length, consisting of standard characters.
 // 返回指定长度的随机字符串
 func NewLen(length int) string {
     return NewLenChars(length, StdChars)
 }
 // NewLenChars returns a new random string of the provided length, consisting
 // of the provided byte slice of allowed characters (maximum 256).
 // 返回指定长度，指定候选字符的随机字符串（最大256）
 func NewLenChars(length int, chars []byte) string {
     b := make([]byte, length)
     r := make([]byte, length+(length/4)) // storage for random bytes. 随机字节的存储空间, 多读几个以免
     clen := byte(len(chars))
     maxrb := byte(256 - (256 % len(chars))) // 问题， 为什么要申请这么长的数组？ 看下面循环的 continue 条件
     i := 0
     for {
         // rand.Read() 和 io.ReadFull(rand.Reader) 的区别?
         // http://www.cnblogs.com/ghj1976/p/3435940.html
         if _, err := io.ReadFull(crand.Reader, r); err != nil {
             panic("error reading from random source: " + err.Error())
         }
         for _, c := range r {
             if c >= maxrb {
                 // Skip this number to avoid modulo bias.
                 // 跳过 maxrb， 以避免麻烦,这也是随机数要多读几个的原因。
                 continue
             }
             b[i] = chars[c%clen]
             i++
             if i == length {
                 return string(b)
             }
         }
     }
     panic("unreachable")
 }
 func pic(w http.ResponseWriter, req *http.Request) {
     // 产生验证码byte数组
     d := make([]byte, 4)
     s := NewLen(4)
     d = []byte(s)
     // 把验证码变成需要显示的字符串
     ss := ""
     for v := range d {
         d[v] %= 10
         ss += strconv.FormatInt(int64(d[v]), 32)
     }
     // 图片流方式输出
     w.Header().Set("Content-Type", "image/png")
     NewImage(d, 100, 40).WriteTo(w)
     // 打印出这次使用的验证码
     fmt.Println(ss)
 }
 func index(w http.ResponseWriter, req *http.Request) {
     str := "<meta charset=\"utf-8\"><h3>golang 图片验证码例子</h3><img border=\"1\" src=\"/pic\" alt=\"图片验证码\" onclick=\"this.src='/pic'\" />"
     w.Header().Set("Content-Type", "text/html")
     w.Write([]byte(str))
 }
 func main() {
     http.HandleFunc("/pic", pic)
     http.HandleFunc("/", index)
     s := &http.Server{
         Addr:           ":8080",
         ReadTimeout:    30 * time.Second,
         WriteTimeout:   30 * time.Second,
         MaxHeaderBytes: 1 << 20}
     s.ListenAndServe()
 }
