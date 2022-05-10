package main
import "fmt"
 type book struct{
 bid int 
 bname,bauther string
 bprice float32
 }
 func main(){
     var n int 
     fmt.Printf("\n Enter book details = ")
     fmt.Scanf("%d",&n)
     b1:= make([]book,n)
     fmt.Printf("\n accept book details")
      for i:=0;i<n;i++{
          fmt.Printf("\n enter book id:")
          fmt.Scanf("%d",&b1[i].bid)
          var p string
          fmt.Printf("\n enter book name:")
          fmt.Scanln(&p)
          b1[i].bname=p[0:len(p)]
           var q string
          fmt.Printf("\n enter book Author:")
          fmt.Scanln(&q)
            b1[i].bauther=q[0:len(q)]
             fmt.Printf("\n enter book price:")
          fmt.Scanf("%f",&b1[i].bprice)
          
      }
      for i:=0;i<n;i++{
          fmt.Println(b1[i].bid)
          fmt.Println(b1[i].bname)
          fmt.Println(b1[i].bauther)
          fmt.Println(b1[i].bprice)
      }
      
 }
