+++
title="Send-to-kindle 2014-04"
tags=["send-to-kindle","golang","concurrency","coding","http"]
date="2014-04-03"
draft=false
+++

> 읽어야 할 글들을 Kindle에 던지는데,뭘 킨들에 던졌는지 기록해 두면 좋을 것 같아서 적어본다. 
> 이번달은 읽은 글들이 많지가 않다. T_T

# Go and Package Versioning 

- http://zduck.com/2014/go-and-package-versioning/ 

 > "My dream is for go get to support a version number component in the import path. So github.com/jpoehls/gophermail would fetch the HEAD of the repository as it does today. github.com/jpoehls/gophermail#v1 would fetch the v1 branch or tag."

 http://gopkg.in 을 사용하면 github에 한해, branch,tag로 구분되어 있는 버젼을 패캐지버젼을 사용할수 있게 된다. 

 일반적으로 버젼이 존재함에 있어서 생기는 문제가 있어서 go tools에서 버젼을 지원하지 않는 것은 감정적으로는 조금은 무책임하다고 생각한다.  저자처럼 go get에서 branch나 tag을 지원해주는 정도는 무난하지 않나 싶기도 하다. 

 지금 go get은  branch나 tag을 go version에 대응해서 사용하고 있다.


#  Write code every day

- http://ejohn.org/blog/write-code-every-day/ 

"the feeling of making progress is just as important as making actual progress. " - Inspirational blog post


#  Throttled: Guardian Of The Web Server

- http://0value.com/throttled--guardian-of-the-web-server 

http  요청 접근에 대한 여러가지 방법(rate-limiting,메모리 사용에 대한 요청에 허가 등)을 구현한 라이브러리. 

go 표준 라이브러리인 http.Handler에 적용해서 사용할수 있다

# Go Concurrency Patterns: Pipelines and cancellation

- http://blog.golang.org/pipelines 

goroutine은 동시성을 얻기 위한 언어도구이다. 각 단계가 독립적으로 각 입력과 출력을 다른 동시성을 가지고 있는 파이프라인을 고루틴으로 가지고 어떻게 구성하는지와 동작의 취소에 대한 구현방법 이야기.

# Defining HTTP Status Codes

- http://refugeeks.com/defining-http-status-codes-list-http-status-codes/ : Defining HTTP Status Codes

고양이 사진으로 보는 http 상태 코드들.


# Some Thoughts on Go and Erlang

- http://blog.erlware.org/2014/04/27/some-thoughts-on-go-and-erlang/ 

제목과 달리 (Erlang에 비해) Go의 단점에 대한 이야기.
결론적으로 결함 감내(Fault tolerant )와 낮은 지연 시간(low latency)에 Go가 적합하지 않다는 내용. 


Fault tolerant이 필요한 복잡한 시스템에서   Erlang에 비해 상태를 공유하는 다른 언어처럼 쉽게 망가지고,   Go의 현재 GC가 Global mark-and-sweep이기 때문에  Low latency을 얻기 힘든다는 내용이다.
