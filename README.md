## Image 2 ASCII converter

This is a small project to study and to experiment with image manipulation, CLIs and creating a cool tools using Golang.

This tool can generate ASCII images directly from a local image or URLs. There's also a command to desaturate images that can use URL as well 

### Installation
First of all, clone this repo:

```shell
git clone https://github.com/davidpalves/img2ascii
```
Then, inside the cloned directory, install all dependencies using:
```shell
go get
```
Compile the code using the following command:
```
go build
```

Now you're ready to use it!

### Usage

You can see the available commands of img2ascii using the following command:

```shell
./img2ascii
```
or 

```shell
./img2ascii --help
```

The following output will be given:
```
Convert images to ASCII art

Usage:
  img2ascii [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  convert     Converts a given image to ASCII Art
  desaturate  Creates a desaturated version of an image
  help        Help about any command

Flags:
  -h, --help            help for img2ascii
      --output string   Path to the desaturated image. E.g.: samples/image.png (default "grayscale.png")
      --path string     Path to image file to be converted
      --url string      Image URL to be converted
      --width int       Width of the output image (default 80)

Use "img2ascii [command] --help" for more information about a command.
```

You can check each commands usage as well by typing:

```shell
./img2ascii [command] --help
```

### Demo
Command:
```shell
./img2ascii convert --url=https://www.hardwinsoftware.com/blog/wp-content/uploads/2018/02/golang-gopher.png --width=130
```
Output:

```
                                                                                  .^^^..
                                                                           "l_]{)(((||()1{[_l^
                                                                    "". ,-)//\)}]]]]]??{(|\tt|1iI+__:
                                                                  -({}fr|j/{--])|[][]{/}+;:l?\)|zj1{t1
                                                                 }t)Xzr1\i      lt}]1\.      "t[?c&t+f?
                                                                 n~c 0_\?-X(     !f-fIJkf     if?_Lr+t]
                                                                 [t{X]?jl& b:    ^j-tiBBz.    if?[}u|1
                                                                  ltf?]))i/!     {{?(1;~     It}][?u!
                                                                   {(][](1:`..,~1rk@@Wr+~!!_1\}][[?/_
                                                                   t}[[[]{||)(||/|pakC(f(((1}]][[[])|
                                                                   f[[[[[]]]][?(|;!__l;{/-]]][[[[[[}t
                                                                   f[[[[[[[[[[[[\jI|__x({[[[[[[[[[[[f
     i?]_!^                                                        f[[[[[[[[[[[[]j`\!lj-[[[[[[[[[[[]j`
  ;CWodpkoW?                                                       /{][[[[[[[[[[]{/(||}[[[[[[[[[[[[]r^
 l&@}    .l"  .xqppL-  ,0n       -On    uQI  l0+  Izqpppm?         ()][[[[[[[[[[[]]]]]][[[[[[[[[[[[]f^
 C (         ,a@|i-m j ! k      .M  ]   a &: ] ( ]BM(ii-\~         }|?[[[[[[[[[[[[[[[[[[[[[[[[[[[[[]t^
 b ~   ij\|I X (    8$`; q      v@;o#   aaQ8.+ 1 h !               _/?[[[[[[[[[[[[[[[[[[[[[[[[[[[[[]j`
 Q \   ]C* ) O ] .  * ": p     ; o(m r  ao qa] 1 M , `Odb{       .:[t?[[[[[[[[[[[[[[[[[[[[[[[[[[[[[?r-;
 iB |`   v { { q.  ) Z l d     pBzJX*$; oM  bB ( v Q^ :z f      -{_|j?[[[[[[[[[[[[[[[[[[[[[[[[[[[[[?j1_(l
  i0%MbdkWW?  )oamdMu  lB&bbbqt$|   _BL kh  .dB{  xMMpZa8)     .u{~jt?[[[[[[[[[[[[[[[[[[[[[[[[[[[[[?ft~n_
    .!?]_I.     l-~.    ",;;;;,"     ": ^^    "     I_-!^       ^+;]/?[[[[[[[[[[[[[[[[[[[[[[[[[[[[[?/]I;
                                                                   }|?[[[[[[[[[[[[[[[[[[[[[[[[[[[[[?\]
                                                                   ()][[[[[[[[[[[[[[[[[[[[[[[[[[[[[]|{
                                                                   /{][[[[[[[[[[[[[[[[[[[[[[[[[[[[[]({
                                                                   t}[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[]|{
                                                                   f}[[[[[[[[[[[[[[[[[[[[[[[[[[[[[[?\]
                                                                   /}][[[[[[[[[[[[[[[[[[[[[[[[[[[[[?f!
                                                                   ()][[[[[[[[[[[[[[[[[[[[[[[[[[[[[]j.
                                                                   +t?[[[[[[[[[[[[[[[[[[[[[[[[[[[[?|1
                                                                    j}][[[[[[[[[[[[[[[[[[[[[[[[[[[]r^
                                                                    lr]][[[[[[[[[[[[[[[[[[[[[[[[]]ri
                                                                     l/1?][[[[[[[[[[[[[[[[[[[]]])xl
                                                                     `{r/1[]]]]]][[[[[[]]]]]]}|/\)i
                                                                    !/-:?rx||)1{}[[[[[[}{1(||{(/I!f?
                                                                    +c[}}; `I+?}{1)(()1{}?~I`  ;1{n~
                                                                     :i;                         ,.

```