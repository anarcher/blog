<!-- Title:Bidirectional RPC in go -->
<!-- Tags: golang,rpc -->
<!-- Languages: shell,go -->


예를 들어, Master Server와 Master의 지령을 받는 Worker Daemon가 있다고 하자. 

Master가 특정 Task를 Worker에 요청을 하려고 하는데, 여러가지 이유로 Master가 Worker에 접속을 하는 구조가 아닌, Worker가 실행할때 Master에 접속 할수만 있다고 하자. 

(특히 내부 IP만 가진 Worker라면 Master가 Worker에 접속을 할수가 없다.)

ThriftRPC나 MsgpackRPC의 구조를 보면, Client/Server 구조로 Client가 Server의 Procedure을 원격 실행하는 방법만 지원한다. 

(Thrift에서 양방향 RPC을 쓰고 싶다면 http://joelpm.com/2009/04/03/thrift-bidirectional-async-rpc.html 을 참고,써보지는 않았다. 자바만 지원.)

밑으로 내려가서 보면, TCP은 기본적으로 Full-duplex 데이터 전달이 가능하다. 
TCP, 또는 웹이라면 WebSocket 기반으로 Bidirectional RPC을 만들면 가능하다고 할수 있는데.

RPC을 만들기 보다는 기존의 RPC 라이브러리를 사용하고 싶어서, 다음과 같이 구성을 해보려고 한다.

- MasterServer는 두개의 정의된 Port를 Listen 한다. 
- 한 Port는 Client가 Server에 RPC하는 Port이고,
- 나머지 한 Port은 Server가 Client에 RPC하는 Port이다. 

Connection이 두개 생기기는 하지만, 기존의 RPC 라이브러리를 사용하여 양방향 RPC을 가능해진다. 
golang의 [net/rpc](http://golang.org/pkg/net/rpc/) 을 가지고 간단히 Server에서 Client에 Remote Procedure Call을 해 볼수 있다.

<pre><code data-language="go">package main

import (
    "net"
    "net/rpc"
    "log"
    "time"
    "fmt"
)

func requestToClient(conn net.Conn) {
    client := rpc.NewClient(conn)
    var out string

    for i := 0 ; i < 5 ; i ++  {
        log.Println("S: echo.echo")
        client.Call("Echo.Echo",fmt.Sprintf("%s:%d","hello client!",i),&out)
        log.Println("E: echo.echo",out)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {

    li,err := net.Listen("tcp",":1234")
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Starting server")

    for {
        conn,_err := li.Accept()
        if _err != nil {
            log.Fatal(err)
        }

        log.Printf("Connected client: %v",conn.RemoteAddr())
        go requestToClient(conn)
    }
}
</code></pre>

Client가 접속을 하면 Server은 접속한 Connection을 가지고  rpc client을 만들어서 요청을 한다.

<pre><code data-language="go">package main
import (
    "net"
    "net/rpc"
    "log"
    "fmt"
)


func main() {

    echo := new(Echo)
    rpc.Register(echo)

    conn,err := net.Dial("tcp",":1234")
    if err != nil {
        log.Fatal(err)
    }

    go rpc.ServeConn(conn)

    var input string
    fmt.Scanln(&input)

}
</code></pre>

Client은 Server에 접속한 후에 등록된 Procedure의 요청을 받기 위해 rpc.ServeConn을 사용한다.

