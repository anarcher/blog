<!-- Title: Send-to-kindle 2014-03 -->
<!-- Tags: send-to-kindle,golang,devops,concurrency,datomic -->

> 읽어야 할 글들을 Kindle에 던지는데,뭘 킨들에 던졌는지 기록해 두면 좋을 것 같아서 적어본다. 

# [Why Atom Can’t Replace Vim : Learning the lesson of vi](https://medium.com/p/433852f4b4d1)

Atom이 Emacs의 확장가능성은 가질수 있지만, Vi의 조합가능성에 대해서는 배운 바가 없다는 글. 

조합가능성에 대한 글로는 
http://blog.dahlia.kr/post/78940164278 홍민희님의 글과 함께 보면 좋다.

내가 생각하는 조합 가능성은  명령어들의 조합이 명령어와 차별 받지 않고 같은 취급을 받는 것이 아닌가 싶다. 유닉스의 파이프로 연결하여 또 하나의 명령어를 만들듯이.  

또 하나의 생각은 Emacs은 어떤 기능이 필요할때 플러그인으로 확장해서 기능을 구현하지만, Vi은 외부 명령어와의 조합으로 기능을 구현한다라는 점. 

# [Sane Concurrency with Go](https://www.readability.com/articles/mmtzoqbu)

공유되는 상태의 병렬로 상태를 바꾸려고 할때, 뮤텍스가 아닌 고루틴과 채널을 이용하여 동기화하는 방법과 디자인에 대한 글 

상태를 가지고 있는 객체 내부에 하나의 고루틴  이벤트 루프를 두어서 동기화를 하고,버퍼가 없는 채널을 이용하는 메소드 인터페이스를 구성하여, 외부에서는 간단히  일반적인 호출을 하는 디자인을 설명.

Heka의 기본적인 디자인 패턴이기도 함.

TwistedPython의 저자 Glyph의 Unyielding에 있는 계좌이체 예제를 Go언어의 고루틴과 채널을 이용하여 구현하는 방법에 대한 글. 

# [Unyielding](https://glyph.twistedmatrix.com/2014/02/unyielding.html)

TwistedPython의 저자 Glyph의 동시성에 대한 글. 

#  [Why Threads Are A Bad Idea (for most purposes)](http://www.stanford.edu/~ouster/cgi-bin/papers/threads.pdf)

대부분의 경우 쓰레드가 왜 나쁜 아이디어인지에 대한 키노트.

많은 경우 쓰레드보다는 이벤트 방식이 좀더 유용하다. (GUI 이벤트,분산 시스템 등등..)
쓰레드가 이벤트보다 강력하지만 그 만큼 위험하다.
쓰레드는 CPU기반 동시성이 필요할때만 사용하자.

# [etcd + ansible = crazy delicious](http://www.unicornclouds.com/blog_posts/etcd_ansible_integration)

ansible은 기본적으로 파일을 기반으로 접속할 서버를 찾거나 기타 정보를 얻는다. 
ansible에서 이러한 정보를 인벤토리라고 부르는데, ansible tower라는 인벤토리 서버와 웹 관리도구을 상용으로 제공하기도 한다. 

ansible에는 동적 인벤토리라는 기능을 제공하는데. 이 기능을 이용하여 인벤토리를 etcd에 저장하고 ansible을 사용하는 방법에 대한 글.

# [Datomic Rationale Why did we build Datomic?](http://www.datomic.com/rationale.html)

기존의 데이터베이스는 Client/Server구조를 기본적으로 가진다. Scalability을 위해 Server을 Clustering을 하거나 Sharing,Replicating을 하게 된다. 
그리고 Client(Application)에서 일종의 Caching을 구성하게 된다. 

Datomic은 이러한 전통적인 데이터베이스의 구성요소를 해체(Deconstruct)해서 다음과 같은 구조로 각 요소에 대한 좀 더  효율성(Efficiency)과 확장성(Scalability)을 구성할수 있도록 한다.

- Peers : LiveIndex,Local Cache,Query (쿼리 할 때 Live Index,Local cache를 참고한다 (모든 데이터에 대한 캐쉬가 각 Peer에 존재하는 듯 하다(기본적으로)) 
- Transactor : 각 Peer의 Writes에 대한 모든 Transaction을 담당한다.
- Storage services : 영속적인 스토리지 서비스에 대한 인터페이스 
- Cache : Optional

![img](http://www.datomic.com/uploads/3/5/9/7/3597326/119999_orig.jpg)


# [BUILDING YOUR FIRST APP ON COREOS: START TO FINISH](http://www.centurylinklabs.com/building-your-first-app-on-coreos/)

mac osx에서 (다른 OS도 가능하겠지만) coreos와 fig등을 이용하여 wordpress을 사용하는 방법을 적은 글. 
개인적으로 fleet에 대한 관심이 생기게 된 글이다.
